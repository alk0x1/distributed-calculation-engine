package tests

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"matrix"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func generateTestMatrix(size int) [][]float64 {
	matrix := make([][]float64, size)
	for i := range matrix {
		matrix[i] = make([]float64, size)
		for j := range matrix[i] {
			matrix[i][j] = float64((i + 1) * (j + 1))
		}
	}
	return matrix
}

func TestServerMultiplication(t *testing.T) {
	matrixA := generateTestMatrix(3)
	matrixB := generateTestMatrix(3)

	payload := matrix.MatrixPayload{
		MatrixA: matrixA,
		MatrixB: matrixB,
	}

	payloadBytes, err := json.Marshal(payload)
	assert.NoError(t, err, "Error marshaling payload")

	resp, err := http.Post("http://localhost:8080/multiply", "application/json", bytes.NewBuffer(payloadBytes))
	assert.NoError(t, err, "Error sending request to server")

	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode, "Server response error")

	respBody, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err, "Error reading response body")

	var result [][]float64
	err = json.Unmarshal(respBody, &result)
	assert.NoError(t, err, "Error unmarshalling response")

	// Validate the response's structure:
	assert.Equal(t, len(matrixA), len(result), "Row count mismatch")
	assert.Equal(t, len(matrixB[0]), len(result[0]), "Column count mismatch")

	// Validate specific multiplication outcomes:
	expectedMatrix, err := matrix.Multiply(matrix.NewMatrix(matrixA), matrix.NewMatrix(matrixB))
	assert.NoError(t, err, "Error in matrix multiplication")

	expectedResult := expectedMatrix.Data
	assert.Equal(t, expectedResult, result, "Multiplication result mismatch")
}
