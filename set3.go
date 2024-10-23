package main

import "strings"

func main() {
	// Placeholder
}

// Relationship status
func relationshipStatus(fromMember string, toMember string, socialGraph map[string]map[string]string) string {
	if _, exists := socialGraph[fromMember]; !exists {
		return "no relationship"
	}
	if _, exists := socialGraph[toMember]; !exists {
		return "no relationship"
	}

	fromFollowingStr := socialGraph[fromMember]["following"]
	toFollowingStr := socialGraph[toMember]["following"]

	var fromFollowing []string
	var toFollowing []string

	if fromFollowingStr != "" {
		fromFollowing = strings.Split(fromFollowingStr, ",")
	}
	if toFollowingStr != "" {
		toFollowing = strings.Split(toFollowingStr, ",")
	}

	fromFollowsTo := contains(fromFollowing, toMember)
	toFollowsFrom := contains(toFollowing, fromMember)

	if fromFollowsTo && toFollowsFrom {
		return "friends"
	} else if fromFollowsTo {
		return "follower"
	} else if toFollowsFrom {
		return "followed by"
	}
	return "no relationship"
}

func contains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}

// Tic tac toe
func ticTacToe(board [][]string) string {
	size := len(board)
	lines := make([][]string, 2*size+2)

	// Collect all lines to check (rows, columns, diagonals)
	for i := 0; i < size; i++ {
		lines[i] = board[i]                 // rows
		lines[size+i] = getColumn(board, i) // columns
	}
	lines[2*size] = getDiagonal(board, true)    // main diagonal
	lines[2*size+1] = getDiagonal(board, false) // anti-diagonal

	// Check all lines
	for _, line := range lines {
		if winner := checkLine(line); winner != "" {
			return winner
		}
	}

	// Check if the board is full
	for _, row := range board {
		for _, cell := range row {
			if cell == "" {
				return "NO WINNER" // Game not finished
			}
		}
	}

	return "NO WINNER" // Draw
}

func checkLine(line []string) string {
	if line[0] == "" {
		return ""
	}
	for _, cell := range line[1:] {
		if cell != line[0] {
			return ""
		}
	}
	return line[0]
}

func getColumn(board [][]string, col int) []string {
	column := make([]string, len(board))
	for i := range board {
		column[i] = board[i][col]
	}
	return column
}

func getDiagonal(board [][]string, isMainDiagonal bool) []string {
	size := len(board)
	diagonal := make([]string, size)
	for i := 0; i < size; i++ {
		if isMainDiagonal {
			diagonal[i] = board[i][i]
		} else {
			diagonal[i] = board[i][size-1-i]
		}
	}
	return diagonal
}

// ETA
func eta(firstStop string, secondStop string, routeMap map[string]map[string]int) int {
	if firstStop == secondStop {
		return 0
	}
	visited := make(map[string]bool)
	times := make(map[string]int)

	for stop := range routeMap {
		times[stop] = int(^uint(0) >> 1) // Maximum int value
	}
	times[firstStop] = 0

	for {
		minTime := int(^uint(0) >> 1)
		var currentStop string
		for stop := range routeMap {
			if !visited[stop] && times[stop] < minTime {
				minTime = times[stop]
				currentStop = stop
			}
		}

		if currentStop == "" || minTime == int(^uint(0)>>1) {
			break
		}

		visited[currentStop] = true

		for nextStop, duration := range routeMap[currentStop] {
			if !visited[nextStop] {
				newTime := times[currentStop] + duration
				if newTime < times[nextStop] {
					times[nextStop] = newTime
				}
			}
		}
	}

	return times[secondStop]
}
