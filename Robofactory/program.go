package main

func OddOrEven(num int) string {
	if num%2 == 1 {
		return "Odd"
	}
	return "Even"
}

func answerIsWrong(num int, answer string) bool {
	return answer != OddOrEven(num)
}

func Reveal(query []int, answer []string) int {
	elected := []int{0}
	for i, j := 0, 1; i < len(query); i, j = i+1, j+1 {
		result := OddOrEven(query[i])
		if j == len(query) {
			break
		}
		if result == "Odd" && answerIsWrong(query[j], answer[j]) {
			// it's corrupted for sure
			return j
		}
		// it may be corrupted
		if result == "Even" {
			elected = append(elected, j)
		}
	}
	// just one investigated -> that's it
	if len(elected) == 1 {
		return elected[0]
	}
	// cannot be determined
	return -1
}
