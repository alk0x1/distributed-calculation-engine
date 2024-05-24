package tests

// import (
// 	"matrix"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestMultiplication(t *testing.T) {
// 	matrixA := matrix.NewMatrix([][]float64{
// 		{1, 2},
// 		{3, 4},
// 	})
// 	matrixB := matrix.NewMatrix([][]float64{
// 		{2, 0},
// 		{1, 3},
// 	})

// 	expected := matrix.NewMatrix([][]float64{
// 		{4, 6},
// 		{10, 12},
// 	})

// 	result, err := matrix.Multiply(matrixA, matrixB)
// 	if err != nil {
// 		t.Fatalf("Error in matrix multiplication: %v", err)
// 	}

// 	assert.Equal(t, expected, result)
// }

// func TestTranspose(t *testing.T) {
// 	matrixA := matrix.NewMatrix([][]float64{
// 		{1, 2, 3},
// 		{4, 5, 6},
// 	})

// 	expected := matrix.NewMatrix([][]float64{
// 		{1, 4},
// 		{2, 5},
// 		{3, 6},
// 	})

// 	result := matrix.Transpose(matrixA)

// 	assert.Equal(t, expected, result)
// }

// func TestAddition(t *testing.T) {
// 	matrixA := matrix.NewMatrix([][]float64{
// 		{1, 2},
// 		{3, 4},
// 	})

// 	matrixB := matrix.NewMatrix([][]float64{
// 		{5, 6},
// 		{7, 8},
// 	})

// 	expected := matrix.NewMatrix([][]float64{
// 		{6, 8},
// 		{10, 12},
// 	})

// 	result, err := matrix.Add(matrixA, matrixB)

// 	if err != nil {
// 		t.Fatalf("Error in matrix multiplication: %v", err)
// 	}
// 	assert.Equal(t, expected, result)
// }

// func TestSubtraction(t *testing.T) {
// 	matrixA := matrix.NewMatrix([][]float64{
// 		{7, 5},
// 		{8, 4},
// 	})

// 	matrixB := matrix.NewMatrix([][]float64{
// 		{2, 1},
// 		{3, 2},
// 	})

// 	expected := matrix.NewMatrix([][]float64{
// 		{5, 4},
// 		{5, 2},
// 	})

// 	result, err := matrix.Subtract(matrixA, matrixB)
// 	if err != nil {
// 		t.Fatalf("Error in matrix multiplication: %v", err)
// 	}
// 	assert.Equal(t, expected, result)
// }
