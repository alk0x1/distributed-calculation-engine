package tests

import (
	"matrix"
	"net/rpc"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatrixServiceImplementation(t *testing.T) {
	// Connect to the existing RPC service:
	client, err := rpc.Dial("tcp", "localhost:8080")
	assert.NoError(t, err, "Error connecting to RPC service")

	defer client.Close()

	// Create a test payload:
	payload := matrix.MatrixPayload{
		MatrixA: [][]float64{{1, 2, 3}, {4, 5, 6}},
		MatrixB: [][]float64{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
	}

	var result matrix.Matrix
	err = client.Call("MatrixServiceImpl.Multiply", payload, &result)
	assert.NoError(t, err, "Error in RPC call")

	// Expected result:
	expected := [][]float64{{1, 2, 3}, {4, 5, 6}}

	// Validate the result:
	assert.Equal(t, expected, result.Data, "Multiplication result mismatch")
}

func TestCommunication(t *testing.T) {
	// Ensure an RPC service is running separately.
	client, err := rpc.Dial("tcp", "localhost:8080")
	assert.NoError(t, err, "Error connecting to RPC service")

	defer client.Close()

	// Make a test request or call:
	var result matrix.Matrix
	payload := matrix.MatrixPayload{
		MatrixA: [][]float64{{1, 2, 3}},
		MatrixB: [][]float64{{1}, {2}, {3}},
	}

	err = client.Call("MatrixServiceImpl.Multiply", payload, &result)
	assert.NoError(t, err, "Error in RPC call")

	expected := [][]float64{{14}}
	assert.Equal(t, expected, result.Data, "Multiplication mismatch")
}

func TestCommunicationInterProcess(t *testing.T) {

}

func generate8x8Matrix() [][]float64 {
	matrix := make([][]float64, 8)
	for i := range matrix {
		matrix[i] = make([]float64, 8)
		for j := range matrix[i] {
			matrix[i][j] = float64((i + 1) * (j + 1))
		}
	}
	return matrix
}

func TestMultiplyServer(t *testing.T) {
	matrixA := generate8x8Matrix()
	matrixB := generate8x8Matrix()

	payload := matrix.MatrixPayload{
		MatrixA: matrixA,
		MatrixB: matrixB,
	}

	client, err := rpc.Dial("tcp", "localhost:8080")
	assert.NoError(t, err, "Error connecting to server")

	defer client.Close()

	var result matrix.Matrix
	t.Logf("payload: %v\n", payload)
	err = client.Call("MatrixServiceImpl.Multiply", payload, &result)
	assert.NoError(t, err, "Error calling Multiply")

	expected := matrix.Matrix{
		Data: [][]float64{
			{204, 408, 612, 816, 1020, 1224, 1428, 1632},
			{408, 816, 1224, 1632, 2040, 2448, 2856, 3264},
			{612, 1224, 1836, 2448, 3060, 3672, 4284, 4896},
			{816, 1632, 2448, 3264, 4080, 4896, 5712, 6528},
			{1020, 2040, 3060, 4080, 5100, 6120, 7140, 8160},
			{1224, 2448, 3672, 4896, 6120, 7344, 8568, 9792},
			{1428, 2856, 4284, 5712, 7140, 8568, 9996, 11424},
			{1632, 3264, 4896, 6528, 8160, 9792, 11424, 13056},
		},
		Rows: 8,
		Cols: 8,
	}

	assert.Equal(t, expected, result, "Server Multiply result mismatch")
}
