package main

import (
	"reflect"
	"testing"
)

func TestMatrixMultiply(t *testing.T) {
	// Matrix A
	A := newMatrix(2, 3)
	A.Set(0, 0, 1)
	A.Set(0, 1, 2)
	A.Set(0, 2, 3)
	A.Set(1, 0, 4)
	A.Set(1, 1, 5)
	A.Set(1, 2, 6)

	// Matrix B
	B := newMatrix(3, 2)
	B.Set(0, 0, 7)
	B.Set(0, 1, 8)
	B.Set(1, 0, 9)
	B.Set(1, 1, 10)
	B.Set(2, 0, 11)
	B.Set(2, 1, 12)

	// Expected result
	expected := newMatrix(2, 2)
	expected.Set(0, 0, 58)
	expected.Set(0, 1, 64)
	expected.Set(1, 0, 139)
	expected.Set(1, 1, 154)

	// Perform multiplication
	result, err := A.Multiply(B)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Check if result matches the expected matrix
	if !reflect.DeepEqual(result.Elements, expected.Elements) {
		t.Errorf("expected %v, got %v", expected.Elements, result.Elements)
	}

	// Test incompatible dimensions
	C := newMatrix(2, 2)
	if res, err := A.Multiply(C); res != nil || err == nil {
		t.Error("expected error for incompatible dimensions, got result")
	}

	// Test empty matrix
	D := newMatrix(0, 0)
	E := newMatrix(0, 0)
	result, err = D.Multiply(E)
	if err != nil || result.Rows != 0 || result.Cols != 0 {
		t.Error("unexpected behavior for empty matrices")
	}

	// Test single-element matrix
	F := newMatrix(1, 1)
	F.Set(0, 0, 2)
	G := newMatrix(1, 1)
	G.Set(0, 0, 3)
	result, err = F.Multiply(G)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Elements[0][0] != 6 {
		t.Errorf("expected 6, got %d", result.Elements[0][0])
	}
}
