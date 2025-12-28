package solver

import (
	"testing"
)

func TestSolve(t *testing.T) {
	defaultGoal := []int{1, 2, 3, 4}
	tests := []struct {
		name      string
		start     []int
		goal      []int
		rows      int
		cols      int
		wantMoves int
		wantErr   error
	}{
		{
			name:    "validation error (empty)",
			start:   []int{},
			goal:    defaultGoal,
			rows:    2,
			cols:    2,
			wantErr: ErrEmptyBoard,
		},
		{
			name:    "unsolvable",
			start:   []int{2, 1, 3, 4},
			goal:    defaultGoal,
			rows:    2,
			cols:    2,
			wantErr: ErrUnsolvable,
		},
		{
			name:      "already solved",
			start:     []int{1, 2, 3, 4},
			goal:      defaultGoal,
			rows:      2,
			cols:      2,
			wantMoves: 0,
			wantErr:   nil,
		},
		{
			name:      "2x2 1 move",
			start:     []int{1, 2, 4, 3}, // 4 is blank. 1 2 / _ 3 -> 1 2 / 3 _
			goal:      defaultGoal,
			rows:      2,
			cols:      2,
			wantMoves: 1,
			wantErr:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path, err := Solve(tt.start, tt.goal, tt.rows, tt.cols)
			if err != tt.wantErr {
				t.Errorf("Solve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr == nil {
				if len(path)-1 != tt.wantMoves {
					t.Errorf("Solve() moves = %d, want %d", len(path)-1, tt.wantMoves)
				}
			}
		})
	}
}

func TestIsSolvable(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		rows  int
		cols  int
		want  bool
	}{
		{"2x2 solvable", []int{1, 2, 3, 4}, 2, 2, true},      // inv=0, dist=0 -> even
		{"2x2 unsolvable", []int{2, 1, 3, 4}, 2, 2, false},   // inv=1, dist=0 -> odd
		{"2x2 solvable move", []int{1, 2, 4, 3}, 2, 2, true}, // inv=1, dist=1 (4 at 1,0 -> 1,1 dist=1) -> 1+1=2 even
		{"3x3 solvable", []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3, 3, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			goal := make([]int, len(tt.input))
			for i := range goal {
				goal[i] = i + 1
			}
			if got := isSolvable(tt.input, goal, tt.rows, tt.cols); got != tt.want {
				t.Errorf("isSolvable() = %v, want %v", got, tt.want)
			}
		})
	}
}
