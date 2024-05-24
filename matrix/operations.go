package matrix

import (
	"errors"
	"fmt"
	"strings"
)

func MultiplyRow(m1, m2 Matrix, rowIndex int, totalRows int) ([]float64, error) {
	if m1.Cols != m2.Rows {
		return nil, errors.New("dimensions are incompatible for multiplication")
	}

	row := make([]float64, m2.Cols)

	for j := 0; j < m2.Cols; j++ {
		for k := 0; k < m1.Cols; k++ {
			row[j] += m1.Data[rowIndex][k] * m2.Data[k][j]
		}
	}

	printProgress(rowIndex+1, totalRows)

	return row, nil
}
func printProgress(current, total int) {
	percentage := (float64(current) / float64(total)) * 100
	progressBar := int(percentage / 2)

	// Create a visual representation.
	fmt.Printf("\r[%-50s] %d%%", strings.Repeat("#", progressBar), int(percentage))
}

func Multiply(m1, m2 Matrix) (Matrix, error) {
	if m1.Cols != m2.Rows {
		return Matrix{}, errors.New("dimensions are incompatible for multiplication")
	}

	// Initialize the resulting matrix with appropriate dimensions.
	result := NewMatrix(make([][]float64, m1.Rows))
	for i := range result.Data {
		result.Data[i] = make([]float64, m2.Cols)
	}

	totalRows := m1.Rows

	// Dispatch row computations to different nodes.
	for i := 0; i < totalRows; i++ {
		row, err := MultiplyRow(m1, m2, i, totalRows)
		if err != nil {
			return Matrix{}, err
		}

		result.Data[i] = row
	}

	// Ensure the resulting matrix's dimensions are set.
	result.Cols = m2.Cols

	fmt.Printf("result: %v\n", result)

	return result, nil
}

/*
// func Transpose(m Matrix) Matrix {
// 	result := NewMatrix(make([][]float64, m.Cols))
// 	for i := range result.Data {
// 		result.Data[i] = make([]float64, m.Rows)
// 	}

// 	// Transposing elements.
// 	for i := 0; i < m.Rows; i++ {
// 		for j := 0; j < m.Cols; j++ {
// 			result.Data[j][i] = m.Data[i][j]
// 		}
// 	}

// 	result.Cols = m.Rows
// 	return result
// }

// func Add(m1, m2 Matrix) (Matrix, error) {
// 	if m1.Rows != m2.Rows || m1.Cols != m2.Cols {
// 		return Matrix{}, errors.New("dimensions are incompatible for addition")
// 	}

// 	result := NewMatrix(make([][]float64, m1.Rows))
// 	for i := range result.Data {
// 		result.Data[i] = make([]float64, m1.Cols)
// 	}

// 	for i := 0; i < m1.Rows; i++ {
// 		for j := 0; j < m1.Cols; j++ {
// 			result.Data[i][j] = m1.Data[i][j] + m2.Data[i][j]
// 		}
// 	}

// 	result.Cols = m1.Cols
// 	return result, nil
// }

// func Subtract(m1, m2 Matrix) (Matrix, error) {
// 	// Ensure the matrices have compatible dimensions.
// 	if m1.Rows != m2.Rows || m1.Cols != m2.Cols {
// 		return Matrix{}, errors.New("dimensions are incompatible for subtraction")
// 	}

// 	// Initialize the resulting matrix with the same dimensions.
// 	result := NewMatrix(make([][]float64, m1.Rows))
// 	for i := range result.Data {
// 		result.Data[i] = make([]float64, m1.Cols)
// 	}

// 	// Perform the subtraction.
// 	for i := 0; i < m1.Rows; i++ {
// 		for j := 0; j < m1.Cols; j++ {
// 			result.Data[i][j] = m1.Data[i][j] - m2.Data[i][j]
// 		}
// 	}

// 	// Ensure the resulting matrix's dimensions are set.
// 	result.Cols = m1.Cols

// 	return result, nil
// }
*/
