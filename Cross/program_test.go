package main

import "testing"

func AssertExists(result string, t *testing.T) {
	if result != "Exists" {
		t.Error("expected 'Exists'")
	}
}

func AssertDoesNotExist(result string, t *testing.T) {
	if result != "Does not exist" {
		t.Error("expected 'Does not exist'")
	}
}

func TestEmptyBoard(t *testing.T) {
	result := exist([]string{})
	AssertDoesNotExist(result, t)
}

func Test1(t *testing.T) {
	result := exist([]string{".##",
		"###",
		"##."})
	AssertExists(result, t)
}

func Test2(t *testing.T) {
	result := exist([]string{".##",
		"###",
		"#.."})
	AssertDoesNotExist(result, t)
}

func Test3(t *testing.T) {
	result := exist([]string{"######",
		"######",
		"######",
		"######",
		"######",
		"######",
		"######"})
	AssertExists(result, t)
}

func Test4(t *testing.T) {
	result := exist([]string{".#.#",
		"#.#.",
		".#.#",
		"#.#."})
	AssertDoesNotExist(result, t)
}

func Test5(t *testing.T) {
	result := exist([]string{".#.#",
		"###.",
		".###",
		"#.#."})
	AssertExists(result, t)
}

/*
{"######",
 "##.###",
 "######",
 "######",
 "######",
 "######",
 "######"}
*/
