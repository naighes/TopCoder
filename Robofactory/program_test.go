package main

import "testing"

func TestOddFour(t *testing.T) {
	result := OddOrEven(4)
	if result == "Odd" {
		t.Error("expected Even for 4")
	}
}

func TestOddSeven(t *testing.T) {
	result := OddOrEven(7)
	if result == "Even" {
		t.Error("expected Odd for 7")
	}
}

func Test1(t *testing.T) {
	result := Reveal([]int{3, 2, 2}, []string{"Odd", "Odd", "Even"})
	if result != 1 {
		t.Error("expected 1 got ", result)
	}
}

func Test2(t *testing.T) {
	result := Reveal([]int{1, 3, 5, 10}, []string{"Odd", "Odd", "Odd", "Even"})
	if result != 0 {
		t.Error("expected 0 got ", result)
	}
}

func Test3(t *testing.T) {
	result := Reveal([]int{2, 3, 5, 10}, []string{"Even", "Odd", "Odd", "Even"})
	if result != -1 {
		t.Error("expected -1 got ", result)
	}
}

func Test4(t *testing.T) {
	result := Reveal([]int{2, 4, 6, 10}, []string{"Even", "Even", "Even", "Even"})
	if result != -1 {
		t.Error("expected -1 got ", result)
	}
}

func Test5(t *testing.T) {
	result := Reveal([]int{107}, []string{"Odd"})
	if result != 0 {
		t.Error("expected 0 got ", result)
	}
}

func Test6(t *testing.T) {
	result := Reveal([]int{1, 1, 1}, []string{"Odd", "Odd", "Even"})
	if result != 2 {
		t.Error("expected 2 got ", result)
	}
}
