package main

import (
	"strings"
	"testing"
)

func TestDivide(t *testing.T) {
	result, err := divide(10, 2)
	if err != nil {
		t.Fatalf("divide(10, 2) returned an error: %v", err)
	}
	if result != 5 {
		t.Errorf("divide(10, 2) = %v, want 5", result)
	}

	_, err = divide(10, 0)
	if err == nil {
		t.Fatal("divide(10, 0) should return an error")
	}
	if !strings.Contains(err.Error(), "zero") {
		t.Errorf("error %q should explain division by zero", err)
	}
}

func TestSquareRoot(t *testing.T) {
	result, err := squareRoot(25)
	if err != nil {
		t.Fatalf("squareRoot(25) returned an error: %v", err)
	}
	if result != 5 {
		t.Errorf("squareRoot(25) = %v, want 5", result)
	}

	_, err = squareRoot(-4)
	if err == nil {
		t.Fatal("squareRoot(-4) should return an error")
	}
	if !strings.Contains(err.Error(), "negative") {
		t.Errorf("error %q should explain the negative number", err)
	}
}

func TestGreet(t *testing.T) {
	message, err := greet("Sanjeev")
	if err != nil {
		t.Fatalf("greet with a name returned an error: %v", err)
	}
	if message != "Hello, Sanjeev!" {
		t.Errorf("greet returned %q, want %q", message, "Hello, Sanjeev!")
	}

	_, err = greet(" ")
	if err == nil {
		t.Fatal("greet with an empty name should return an error")
	}
	if !strings.Contains(err.Error(), "empty") {
		t.Errorf("error %q should explain that the name is empty", err)
	}
}
