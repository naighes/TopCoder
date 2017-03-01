package main

func exist(board []string) string {
	for i := 1; i < len(board)-1; {
		for j := 1; j < len(board[i])-1; {
			if match(board, i, j) {
				return "Exists"
			}
			j++
		}
		i++
	}
	return "Does not exist"
}

func match(board []string, i int, j int) bool {
	return board[i][j] == '#' &&
		board[i][j+1] == '#' &&
		board[i][j-1] == '#' &&
		board[i+1][j] == '#' &&
		board[i-1][j] == '#'
}
