package solver

import (
	"errors"
	"math"
	"slices"
)

var (
	ErrEmptyBoard     = errors.New("board cannot be empty")
	ErrInvalidSize    = errors.New("invalid puzzle size")
	ErrSizeMismatch   = errors.New("board length does not match rows*cols")
	ErrInvalidElement = errors.New("board must contain numbers from 1 to len(board)")
)

// validate checks if the board slice represents a valid puzzle configuration.
// It verifies that the board is not empty, the grid size is within limits,
// the board length matches the specified rows and columns, and the board contains
// a permutation of numbers from 1 to len(board).
//
// Example:
//
//	validate([]int{1, 2, 3, 4}, 2, 2) // returns nil
//	validate([]int{1, 2, 3, 0}, 2, 2) // returns ErrInvalidElement
func validate(board []int, rows, cols int) error {
	if len(board) == 0 {
		return ErrEmptyBoard
	}

	if rows < 2 || cols < 2 || rows > math.MaxInt/cols {
		return ErrInvalidSize
	}

	if len(board) != rows*cols {
		return ErrSizeMismatch
	}

	for i := 1; i <= len(board); i++ {
		if !slices.Contains(board, i) {
			return ErrInvalidElement
		}
	}

	return nil
}
