package service

import (
	"testing"
)

func TestSum(t *testing.T) {
	result := Sum(10, 5)
	if result != 15 {
		t.Errorf("Sum(10,5) = %f; want 15", result)
	}
}

func TestSub(t *testing.T) {
	result := Sub(10, 5)
	if result != 5 {
		t.Errorf("Sub(10,5) = %f; want 5", result)
	}
}

func TestMul(t *testing.T) {
	result := Mul(10, 5)
	if result != 50 {
		t.Errorf("Mul(10,5) = %f; want 50", result)
	}
}

func TestDiv(t *testing.T) {
	result, err := Div(10, 5)
	if err != nil {
		t.Errorf("Div(10,5) returned error: %v", err)
	}
	if result != 2 {
		t.Errorf("Div(10,5) = %f; want 2", result)
	}

	_, err = Div(10, 0)
	if err == nil {
		t.Error("Div(10,0) expected error, got nil")
	}
}
