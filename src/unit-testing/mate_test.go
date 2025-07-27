package main

import (
	"testing"
)

func TestSuma(t *testing.T) {
	result := Suma(1, 2)
	if result != 3 {
		t.Errorf("Suma(1, 2) = %d; want 3", result)
	}
}

func TestSumaNegative(t *testing.T) {
	result := Suma(-1, -2)
	if result != -3 {
		t.Errorf("Suma(-1, -2) = %d; want -3", result)
	}
}

func TestSuma2(t *testing.T) {
	casos := []struct {
		a, b, expected int
	}{
		{1, 2, 3},
		{-1, -2, -3},
		{2, 3, 5},
		{0, 0, 0},
		{-1, 1, 0},
	}
	for _, c := range casos {
		result := Suma(c.a, c.b)
		if result != c.expected {
			t.Errorf("Suma(%d, %d) = %d; want %d", c.a, c.b, result, c.expected)
		}
	}
}

func TestMayor(t *testing.T) {
	casos := []struct {
		a, b, expected int
	}{
		{2, 1, 2},
		{-1, -2, -1},
		{3, 2, 3},
		{0, 0, 0},
		{1, -1, 1},
	}

	for _, c := range casos {
		result := Mayor(c.a, c.b)
		if result != c.expected {
			t.Errorf("Mayor(%d, %d) = %d; want %d", c.a, c.b, result, c.expected)
		}
	}
}

func TestFibonacci(t *testing.T) {
	casos := []struct {
		n, expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{40, 102334155},
		{50, 12586269025},
	}

	for _, c := range casos {
		result := Fibonacci(c.n)
		if result != c.expected {
			t.Errorf("Fibonacci(%d) = %d; want %d", c.n, result, c.expected)
		}
	}
}
