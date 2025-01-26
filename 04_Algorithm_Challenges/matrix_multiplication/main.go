package main

import "fmt"

type Matrix struct {
	Rows, Cols int
	Elements   [][]int
}

// Function to create a new matrix
func newMatrix(rows, cols int) *Matrix {
	elements := make([][]int, rows)
	for i := range elements {
		elements[i] = make([]int, cols)
	}
	return &Matrix{
		Rows:     rows,
		Cols:     cols,
		Elements: elements,
	}
}

// Set sets the value at a specific row and column
func (m *Matrix) Set(row, col, val int) {
	if row >= 0 && row < m.Rows && col >= 0 && col < m.Cols {
		m.Elements[row][col] = val
	} else {
		fmt.Println("Error: Index out of bounds")
	}
}

// Multiply performs matrix multiplication
func (m *Matrix) Multiply(n *Matrix) (*Matrix, error) {
	if m.Cols != n.Rows {
		return nil, fmt.Errorf("incompatible matrix dimensions: %dx%d and %dx%d", m.Rows, m.Cols, n.Rows, n.Cols)
	}
	result := newMatrix(m.Rows, n.Cols)
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < n.Cols; j++ {
			for k := 0; k < m.Cols; k++ {
				result.Elements[i][j] += m.Elements[i][k] * n.Elements[k][j]
			}
		}
	}
	return result, nil
}
