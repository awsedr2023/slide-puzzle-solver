package solver

// inversionNumber calculates the number of inversions in the puzzle configuration.
// It counts pairs of tiles (i, j) such that i < j and input[i] > input[j].
// Note: This implementation treats the blank tile as the largest number.
//
// Example:
//
//	inversionNumber([]int{1, 3, 2, 4}) // returns 1
//	inversionNumber([]int{3, 2, 1, 4}) // returns 3
func inversionNumber(input []int) int {
	count := 0
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			if input[i] > input[j] {
				count++
			}
		}
	}
	return count
}

// boardManhattanDistance calculates the sum of Manhattan distances of all tiles
// (excluding the blank tile) from their current positions to their goal positions.
//
// Example:
//
//	// In a 2x2 puzzle, 4 is the blank tile.
//	boardManhattanDistance([]int{4, 1, 2, 3}, []int{1, 2, 3, 4}, 2, 2) // returns 4
//	boardManhattanDistance([]int{1, 2, 3, 4}, []int{1, 2, 3, 4}, 2, 2) // returns 0
func boardManhattanDistance(start, goal []int, rows, cols int) int {
	distance := 0
	blank := len(start)

	goalPositions := make(map[int]int, len(goal))
	for i, v := range goal {
		goalPositions[v] = i
	}

	for i, v := range start {
		if v == blank {
			continue
		}

		goalIdx, ok := goalPositions[v]
		if !ok {
			continue
		}

		distance += manhattanDistance(i, goalIdx, rows, cols)
	}
	return distance
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// manhattanDistance calculates the Manhattan distance between two indices on the board.
//
// Example:
//
//	manhattanDistance(0, 3, 2, 2) // returns 2 (0 is at (0,0), 3 is at (1,1))
func manhattanDistance(idx1, idx2, rows, cols int) int {
	r1, c1 := idx1/cols, idx1%cols
	r2, c2 := idx2/cols, idx2%cols
	return abs(r1-r2) + abs(c1-c2)
}

// linearConflict calculates the linear conflict heuristic value.
// Two tiles t_j and t_k are in a linear conflict if they are in the same row (or column),
// their goal positions are also in that row (or column), and t_j is to the right of t_k
// but its goal position is to the left of t_k's goal position.
// It adds 2 to the cost for each conflict that must be resolved.
func linearConflict(board, goal []int, rows, cols int) int {
	conflicts := 0
	blank := len(board)

	goalPositions := make(map[int]int, len(goal))
	for i, v := range goal {
		goalPositions[v] = i
	}

	// Row conflicts
	for r := 0; r < rows; r++ {
		var line []int
		for c := 0; c < cols; c++ {
			idx := r*cols + c
			tile := board[idx]
			if tile != blank {
				// Check if tile belongs to this row in goal
				if gIdx, ok := goalPositions[tile]; ok && gIdx/cols == r {
					line = append(line, tile)
				}
			}
		}
		conflicts += countConflicts(line, goalPositions)
	}

	// Column conflicts
	for c := 0; c < cols; c++ {
		var line []int
		for r := 0; r < rows; r++ {
			idx := r*cols + c
			tile := board[idx]
			if tile != blank {
				// Check if tile belongs to this column in goal
				if gIdx, ok := goalPositions[tile]; ok && gIdx%cols == c {
					line = append(line, tile)
				}
			}
		}
		conflicts += countConflicts(line, goalPositions)
	}

	return conflicts * 2
}

func countConflicts(line []int, goalPositions map[int]int) int {
	conflicts := 0
	// Working copy to mark removed tiles
	tiles := make([]int, len(line))
	copy(tiles, line)

	for {
		maxConflicts := 0
		candidateIdx := -1

		for i, t1 := range tiles {
			if t1 == -1 {
				continue
			}
			currentConflicts := 0
			for j, t2 := range tiles {
				if i == j || t2 == -1 {
					continue
				}
				// Check if t1 and t2 are in conflict.
				// Since they are in 'line' in current order, i < j means t1 is currently before t2.
				// Conflict if goal(t1) > goal(t2).
				if (i < j && goalPositions[t1] > goalPositions[t2]) ||
					(i > j && goalPositions[t1] < goalPositions[t2]) {
					currentConflicts++
				}
			}

			if currentConflicts > maxConflicts {
				maxConflicts = currentConflicts
				candidateIdx = i
			}
		}

		if maxConflicts == 0 {
			break
		}

		// Remove the tile with the most conflicts
		tiles[candidateIdx] = -1
		conflicts++
	}

	return conflicts
}

// calculateHeuristic calculates the total heuristic cost (Manhattan distance + Linear Conflict).
func calculateHeuristic(board, goal []int, rows, cols int) int {
	return boardManhattanDistance(board, goal, rows, cols) +
		linearConflict(board, goal, rows, cols)
}
