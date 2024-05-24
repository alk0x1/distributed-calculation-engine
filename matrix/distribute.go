package matrix

import (
	"errors"
	"fmt"
	"net/rpc"
	"sync"
)

var Servers = []string{
	"localhost:8080",
	"localhost:8081",
}

func CheckServerHealth() {
	for i := 0; i < len(Servers); i++ {
		client, err := rpc.Dial("tcp", Servers[i])
		if err != nil {
			fmt.Println("Server", Servers[i], "is unhealthy.")
			Servers = append(Servers[:i], Servers[i+1:]...)
			i--
			continue
		}

		defer client.Close()

		err = client.Call("MatrixServiceImpl.Ping", struct{}{}, new(struct{}))
		if err != nil {
			fmt.Println("Server", Servers[i], "is unhealthy.")
			Servers = append(Servers[:i], Servers[i+1:]...)
			i--
		} else {
			fmt.Println("Server", Servers[i], "is healthy.")
		}
	}
}
func DistributeTask(payload MatrixPayload, operation string) (Matrix, error) {
	CheckServerHealth()
	chunksA, chunksB := SplitMatrices(payload, len(Servers))

	var wg sync.WaitGroup
	resultsChan := make(chan Matrix, len(chunksA))

	for i, server := range Servers {
		wg.Add(1)

		go func(i int, server string) {
			defer wg.Done()

			fmt.Printf("Connecting to server: %s\n", server)

			client, err := rpc.Dial("tcp", server)
			if err != nil {
				fmt.Printf("Connection error with server %s: %v\n", server, err)
				return
			}

			payloadChunk := MatrixPayload{
				MatrixA: chunksA[i],
				MatrixB: chunksB[i],
			}
			fmt.Printf("PayloadChunk[%v]: %v\n", i, payloadChunk)

			var result Matrix
			err = client.Call("MatrixServiceImpl.Multiply", payloadChunk, &result)
			if err != nil {
				fmt.Printf("Error calling Multiply on server %s: %v\n", server, err)
				client.Close()
				return
			}

			resultsChan <- result
			client.Close()
		}(i, server)
	}

	wg.Wait()
	close(resultsChan)

	var combined Matrix
	for res := range resultsChan {
		combined = MergeMatrices(combined, res)
	}

	if len(combined.Data) > 0 {
		return combined, nil
	}

	return Matrix{}, errors.New("all servers failed")
}

func SplitMatrices(payload MatrixPayload, numChunks int) ([][][]float64, [][][]float64) {
	chunksA := SplitMatrix(payload.MatrixA, numChunks)
	chunksB := SplitMatrix(payload.MatrixB, numChunks)
	return chunksA, chunksB
}

func SplitMatrix(matrix [][]float64, numChunks int) [][][]float64 {
	chunkSize := (len(matrix) + numChunks - 1) / numChunks

	chunks := make([][][]float64, numChunks)
	for i := range chunks {
		start := i * chunkSize
		end := (i + 1) * chunkSize
		if end > len(matrix) {
			end = len(matrix)
		}
		chunks[i] = matrix[start:end]
	}

	return chunks
}

func MergeMatrices(m1, m2 Matrix) Matrix {
	if m1.Cols != m2.Cols {
		return Matrix{}
	}

	m1.Data = append(m1.Data, m2.Data...)

	return Matrix{
		Data: m1.Data,
		Cols: m1.Cols,
	}
}
