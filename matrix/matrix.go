package matrix

import (
	"errors"
)

type MatrixPayload struct {
	MatrixA [][]float64 `json:"matrixA"`
	MatrixB [][]float64 `json:"matrixB"`
}

type Matrix struct {
	Data [][]float64
	Rows int
	Cols int
}

func NewMatrix(data [][]float64) Matrix {
	rows := len(data)
	cols := 0
	if rows > 0 {
		cols = len(data[0])
	}

	return Matrix{
		Data: data,
		Rows: rows,
		Cols: cols,
	}
}

func (m Matrix) Get(row, col int) (float64, error) {
	if row < 0 || row >= m.Rows || col < 0 || col >= m.Cols {
		return 0, errors.New("indices out of range")
	}
	return m.Data[row][col], nil
}

func (m *Matrix) Set(row, col int, value float64) error {
	if row < 0 || row >= m.Rows || col < 0 || col >= m.Cols {
		return errors.New("indices out of range")
	}
	m.Data[row][col] = value
	return nil
}

type MatrixService interface {
	Multiply(MatrixPayload, *Matrix) error
	Ping(struct{}, *struct{}) error // Define a ping function for health checks.

}

type MatrixServiceImpl struct{}

func (msi *MatrixServiceImpl) Multiply(payload MatrixPayload, reply *Matrix) error {
	matrixA := NewMatrix(payload.MatrixA)
	matrixB := NewMatrix(payload.MatrixB)

	result, err := Multiply(matrixA, matrixB)
	if err != nil {
		return err
	}

	*reply = result
	return nil
}

func (msi *MatrixServiceImpl) Ping(struct{}, *struct{}) error {
	return nil
}
