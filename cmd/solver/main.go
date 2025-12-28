package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/awsedr2023/slide-puzzle-solver/solver"
)

func main() {
	rows := flag.Int("rows", 0, "number of rows")
	cols := flag.Int("cols", 0, "number of columns")
	flag.Parse()

	if *rows < 2 || *cols < 2 {
		fmt.Println("Usage: solver -rows <rows> -cols <cols> <numbers...>")
		fmt.Println("Example: solver -rows 3 -cols 3 1 2 3 4 5 6 7 8 9")
		fmt.Println("Example with custom goal: solver -rows 2 -cols 2 1 2 4 3 1 2 3 4")
		os.Exit(1)
	}

	args := flag.Args()
	totalCells := *rows * *cols
	if len(args) != totalCells && len(args) != 2*totalCells {
		fmt.Printf("Error: Expected %d (start) or %d (start followed by goal) numbers, got %d\n", totalCells, 2*totalCells, len(args))
		os.Exit(1)
	}

	input, err := parseBoard(args[:totalCells])
	if err != nil {
		fmt.Printf("Error parsing start board: %v\n", err)
		os.Exit(1)
	}

	var goal []int
	if len(args) == 2*totalCells {
		goal, err = parseBoard(args[totalCells:])
		if err != nil {
			fmt.Printf("Error parsing goal board: %v\n", err)
			os.Exit(1)
		}
	} else {
		goal = solver.StandardGoal(*rows, *cols)
	}

	path, err := solver.Solve(input, goal, *rows, *cols)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Solved in %d moves:\n", len(path)-1)

	board := make([]int, len(input))
	copy(board, input)

	blankIdx := path[0]
	for i, nextBlank := range path[1:] {
		tile := board[nextBlank]

		r1, c1 := blankIdx / *cols, blankIdx%*cols
		r2, c2 := nextBlank / *cols, nextBlank%*cols

		var dir string
		if r2 < r1 {
			dir = "Down"
		} else if r2 > r1 {
			dir = "Up"
		} else if c2 < c1 {
			dir = "Right"
		} else {
			dir = "Left"
		}

		fmt.Printf("%d: Move tile %d %s\n", i+1, tile, dir)
		board[blankIdx], board[nextBlank] = board[nextBlank], board[blankIdx]
		blankIdx = nextBlank
	}
}

func parseBoard(strs []string) ([]int, error) {
	board := make([]int, 0, len(strs))
	for _, s := range strs {
		val, err := strconv.Atoi(s)
		if err != nil {
			return nil, fmt.Errorf("invalid number '%s'", s)
		}
		board = append(board, val)
	}
	return board, nil
}
