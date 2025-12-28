package solver

import (
	"errors"
	"math"
	"slices"
)

var ErrUnsolvable = errors.New("puzzle is unsolvable")

// Solve solves the sliding puzzle and finds the shortest path from the start configuration to the goal configuration.
// The blank tile is represented by the value rows*cols.
// It returns a sequence of blank tile indices representing the path from start to goal, including the initial position.
// Returns an error if the puzzle is unsolvable or inputs are invalid.
//
// Example:
//
//	start := []int{4, 1, 3, 2}
//	goal := []int{1, 2, 3, 4}
//	path, err := Solve(start, goal, 2, 2)
func Solve(start, goal []int, rows, cols int) ([]int, error) {
	if err := validate(start, rows, cols); err != nil {
		return nil, err
	}
	if err := validate(goal, rows, cols); err != nil {
		return nil, err
	}

	if !isSolvable(start, goal, rows, cols) {
		return nil, ErrUnsolvable
	}

	root := newNode(start, rows, cols)
	threshold := calculateHeuristic(root.board, goal, rows, cols)

	for {
		nextThreshold, found := search(root, threshold, goal, rows, cols)
		if found != nil {
			var path []int
			for n := found; n != nil; n = n.parent {
				path = append(path, n.blankIdx)
			}
			slices.Reverse(path)
			return path, nil
		}

		if nextThreshold == math.MaxInt {
			return nil, ErrUnsolvable
		}
		threshold = nextThreshold
	}
}

// isSolvable checks if the puzzle configuration can be solved to reach the target state.
//
// Example:
//
//	start := []int{4, 1, 3, 2}
//	goal := []int{1, 2, 3, 4}
//	isSolvable(start, goal, 2, 2) // returns true
func isSolvable(start, goal []int, rows, cols int) bool {
	startInversionNumber := inversionNumber(start)
	blank := rows * cols
	var startBlankIndex, goalBlankIndex int
	for i, v := range start {
		if v == blank {
			startBlankIndex = i
		}
	}
	for i, v := range goal {
		if v == blank {
			goalBlankIndex = i
		}
	}
	manhattanDistance := manhattanDistance(startBlankIndex, goalBlankIndex, rows, cols)
	goalInversionNumber := inversionNumber(goal)
	return (startInversionNumber+manhattanDistance)%2 == goalInversionNumber%2
}

// search performs the Depth-First Search for IDA*.
// It returns the next threshold (min f-value exceeding current threshold) or the goal node.
func search(currentNode *node, threshold int, goal []int, rows, cols int) (int, *node) {
	heuristic := calculateHeuristic(currentNode.board, goal, rows, cols)
	estimatedTotalCost := currentNode.cost + heuristic

	if estimatedTotalCost > threshold {
		return estimatedTotalCost, nil
	}

	if currentNode.has(goal) {
		return estimatedTotalCost, currentNode
	}

	minNextThreshold := math.MaxInt

	processNeighbor := func(neighbor *node) (int, *node) {
		if isCycle(neighbor) {
			return math.MaxInt, nil
		}
		return search(neighbor, threshold, goal, rows, cols)
	}

	var moves []func() *node
	if currentNode.canMoveUp() {
		moves = append(moves, currentNode.upNode)
	}
	if currentNode.canMoveDown() {
		moves = append(moves, currentNode.downNode)
	}
	if currentNode.canMoveLeft() {
		moves = append(moves, currentNode.leftNode)
	}
	if currentNode.canMoveRight() {
		moves = append(moves, currentNode.rightNode)
	}

	for _, move := range moves {
		res, found := processNeighbor(move())
		if found != nil {
			return res, found
		}
		if res < minNextThreshold {
			minNextThreshold = res
		}
	}

	return minNextThreshold, nil
}

func isCycle(node *node) bool {
	for ancestor := node.parent; ancestor != nil; ancestor = ancestor.parent {
		if node.has(ancestor.board) {
			return true
		}
	}
	return false
}
