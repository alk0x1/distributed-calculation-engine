package tests

import (
	"matrix"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
	func TestCheckServerHealth(t *testing.T) {
		originalServers := matrix.Servers

		scenarios := []struct {
			initialServers []string
			expectedCount  func() int
			description    string
		}{
			{[]string{"localhost:8080", "localhost:8081"}, func() int { return 2 }, "Both servers running"},
			{[]string{"localhost:8080", "localhost:8081"}, func() int { return 1 }, "One server running"},
			{[]string{"localhost:8080", "localhost:8081"}, func() int { return 0 }, "No servers running"},
		}

		for _, scenario := range scenarios {
			matrix.Servers = scenario.initialServers

			matrix.CheckServerHealth()

			actualCount := len(matrix.Servers)
			assert.Equal(t, scenario.expectedCount(), actualCount, scenario.description)

			matrix.Servers = originalServers // Reset for the next iteration.
		}
	}
*/
func TestSplitMatrix(t *testing.T) {
	// Test matrix:
	test_matrix := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	chunks := matrix.SplitMatrix(test_matrix, 2)

	// Check the chunks' structures:
	assert.Equal(t, 2, len(chunks), "Chunk count mismatch")
	assert.Equal(t, [][]float64{{1, 2, 3}, {4, 5, 6}}, chunks[0], "First chunk mismatch")
	assert.Equal(t, [][]float64{{7, 8, 9}}, chunks[1], "Second chunk mismatch")

	test_matrix2 := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{10, 11, 12},
	}

	numChunks := 2

	expectedChunks := [][][]float64{
		{{1, 2, 3}, {4, 5, 6}},
		{{7, 8, 9}, {10, 11, 12}},
	}

	chunks = matrix.SplitMatrix(test_matrix2, numChunks)

	if !reflect.DeepEqual(chunks, expectedChunks) {
		t.Errorf("SplitMatrix did not produce the expected result. Got: %v, Expected: %v", chunks, expectedChunks)
	}
}

func TestSplitMatrices(t *testing.T) {
	matrixA := [][]float64{
		{1, 2, 3, 4, 5, 6, 7, 8},
		{1, 2, 3, 4, 5, 6, 7, 8},
		{1, 2, 3, 4, 5, 6, 7, 8},
		{1, 2, 3, 4, 5, 6, 7, 8},
		{1, 2, 3, 4, 5, 6, 7, 8},
		{1, 2, 3, 4, 5, 6, 7, 8},
		{1, 2, 3, 4, 5, 6, 7, 8},
		{1, 2, 3, 4, 5, 6, 7, 8},
	}

	matrixB := [][]float64{
		{1, 0, 0, 1},
		{1, 0, 0, 1},
		{1, 0, 0, 1},
		{1, 0, 0, 1},
	}

	payload := matrix.MatrixPayload{
		MatrixA: matrixA,
		MatrixB: matrixB,
	}

	chunksA, chunksB := matrix.SplitMatrices(payload, 2)

	expectedChunksA := [][][]float64{
		{
			{1, 2, 3, 4, 5, 6, 7, 8},
			{1, 2, 3, 4, 5, 6, 7, 8},
			{1, 2, 3, 4, 5, 6, 7, 8},
			{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			{1, 2, 3, 4, 5, 6, 7, 8},
			{1, 2, 3, 4, 5, 6, 7, 8},
			{1, 2, 3, 4, 5, 6, 7, 8},
			{1, 2, 3, 4, 5, 6, 7, 8},
		},
	}

	expectedChunksB := [][][]float64{
		{
			{1, 0, 0, 1},
			{1, 0, 0, 1},
		},
		{
			{1, 0, 0, 1},
			{1, 0, 0, 1},
		},
	}

	// Check the chunks' structures:
	assert.Equal(t, expectedChunksA, chunksA, "MatrixA split mismatch")
	assert.Equal(t, expectedChunksB, chunksB, "MatrixB split mismatch")
}

func TestMergeMatrices(t *testing.T) {
	m1 := matrix.Matrix{
		Data: [][]float64{
			{1, 2, 3},
			{4, 5, 6},
		},
		Cols: 3,
	}

	m2 := matrix.Matrix{
		Data: [][]float64{
			{7, 8, 9},
			{10, 11, 12},
		},
		Cols: 3,
	}

	result := matrix.MergeMatrices(m1, m2)

	expected := matrix.Matrix{
		Data: [][]float64{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
			{10, 11, 12},
		},
		Cols: 3,
	}

	assert.Equal(t, expected, result, "Merged matrices mismatch")
}
func generate8x8Matrixx() [][]float64 {
	matrix := make([][]float64, 8)
	for i := range matrix {
		matrix[i] = make([]float64, 8)
		for j := range matrix[i] {
			matrix[i][j] = float64((i + 1) * (j + 1))
		}
	}
	return matrix
}
func TestDistributeTask(t *testing.T) {
	matrixA := generate8x8Matrixx()
	matrixB := generate8x8Matrixx()

	// Define the payload:
	payload := matrix.MatrixPayload{
		MatrixA: matrixA,
		MatrixB: matrixB,
	}

	// Call the DistributeTask function:
	result, err := matrix.DistributeTask(payload, "multiply")
	assert.NoError(t, err, "Error distributing task")

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

	assert.Equal(t, expected, result, "DistributeTask result mismatch")
}
