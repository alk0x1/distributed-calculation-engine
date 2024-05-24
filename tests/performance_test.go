package tests

// import (
// 	"matrix"
// 	"testing"
// 	"time"

// 	"github.com/stretchr/testify/assert"
// )

// func generateLargeMatrix(size int) [][]float64 {
// 	matrix := make([][]float64, size)
// 	for i := range matrix {
// 		matrix[i] = make([]float64, size)
// 		for j := range matrix[i] {
// 			matrix[i][j] = float64((i + 1) * (j + 1))
// 		}
// 	}
// 	return matrix
// }

// func TestLargeMultiplication(t *testing.T) {
// 	matrixA := generateLargeMatrix(2000)
// 	matrixB := generateLargeMatrix(2000)

// 	payload := matrix.MatrixPayload{
// 		MatrixA: matrixA,
// 		MatrixB: matrixB,
// 	}

// 	start := time.Now()

// 	result, err := matrix.DistributeTask(payload, "multiply")
// 	assert.NoError(t, err, "Error distributing task")

// 	duration := time.Since(start)

// 	assert.Equal(t, len(matrixA), len(result.Data), "Row count mismatch")
// 	assert.Equal(t, len(matrixB[0]), len(result.Data[0]), "Column count mismatch")

// 	t.Logf("Time to compute multiplication: %v", duration)
// }
