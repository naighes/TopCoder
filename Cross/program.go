package main

func exist(board []string) string {
	for i := 1; i < len(board)-1; {
		for j := 1; j < len(board[i])-1; {
			h := HMatch(board, i, j)
			if h == 0 {
				if match(board, i, j) {
					return "Exists"
				} else {
					j = j + 1
				}
			} else {
				j = j + h
			}
		}
		i++
	}
	return "Does not exist"
}

func match(board []string, i int, j int) bool {
	return board[i+1][j] == '#' &&
		board[i-1][j] == '#'
}

func HMatch(board []string, i int, j int) int {
	if board[i][j+1] != '#' {
		return 3
	}
	if board[i][j] != '#' {
		return 2
	}
	if board[i][j-1] != '#' {
		return 1
	}
	return 0
}
