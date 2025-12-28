package solver

import "testing"

func TestInversionNumber(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{"sorted", []int{1, 2, 3}, 0},
		{"reverse", []int{3, 2, 1}, 3},       // (3,2), (3,1), (2,1)
		{"mixed", []int{1, 3, 2}, 1},         // (3,2)
		{"4 elements", []int{2, 4, 1, 3}, 3}, // (2,1), (4,1), (4,3)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inversionNumber(tt.input); got != tt.want {
				t.Errorf("inversionNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestManhattanDistance(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		rows  int
		cols  int
		want  int
	}{
		{"2x2 solved", []int{1, 2, 3, 4}, 2, 2, 0},               // 4 is at (1,1), goal (1,1) -> 0
		{"2x2 start", []int{4, 1, 2, 3}, 2, 2, 4},                // 1:1, 2:2, 3:1 -> 4
		{"3x3 mixed", []int{1, 2, 3, 4, 5, 6, 9, 7, 8}, 3, 3, 2}, // 7:1, 8:1 -> 2
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			goal := StandardGoal(tt.rows, tt.cols)
			if got := boardManhattanDistance(tt.input, goal, tt.rows, tt.cols); got != tt.want {
				t.Errorf("manhattanDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateDistance(t *testing.T) {
	tests := []struct {
		name string
		idx1 int
		idx2 int
		rows int
		cols int
		want int
	}{
		{"same position", 0, 0, 3, 3, 0},
		{"adjacent horizontal", 0, 1, 3, 3, 1},
		{"adjacent vertical", 0, 3, 3, 3, 1},
		{"diagonal", 0, 4, 3, 3, 2},        // (0,0) to (1,1) -> 1+1=2
		{"far corners 3x3", 0, 8, 3, 3, 4}, // (0,0) to (2,2) -> 2+2=4
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := manhattanDistance(tt.idx1, tt.idx2, tt.rows, tt.cols); got != tt.want {
				t.Errorf("calculateDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinearConflict(t *testing.T) {
	tests := []struct {
		name  string
		board []int
		rows  int
		cols  int
		want  int
	}{
		{"no conflict", []int{1, 2, 3, 4}, 2, 2, 0},
		{"row conflict 2x2", []int{2, 1, 3, 4}, 2, 2, 2},               // 2 and 1 swapped in row 0 -> 1 conflict * 2 = 2
		{"col conflict 2x2", []int{3, 2, 1, 4}, 2, 2, 2},               // 3 and 1 swapped in col 0 -> 1 conflict * 2 = 2
		{"3x3 row reverse", []int{3, 2, 1, 4, 5, 6, 7, 8, 9}, 3, 3, 4}, // 3,2,1 in row 0. 3 conflicts with 2 and 1. Remove 3 -> 2,1 conflict. Remove 2 -> no conflict. Total 2 conflicts -> 4
		{"3x3 mixed", []int{
			2, 1, 3, // Row 0: 2-1 conflict (2)
			4, 6, 5, // Row 1: 6-5 conflict (2)
			7, 8, 9, // Row 2: ok
		}, 3, 3, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			goal := StandardGoal(tt.rows, tt.cols)
			if got := linearConflict(tt.board, goal, tt.rows, tt.cols); got != tt.want {
				t.Errorf("%s: linearConflict() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestCalculateHeuristic(t *testing.T) {
	tests := []struct {
		name  string
		board []int
		rows  int
		cols  int
		want  int
	}{
		{"solved", []int{1, 2, 3, 4}, 2, 2, 0},
		{"manhattan only", []int{1, 3, 2, 4}, 2, 2, 4},       // 3:(0,1)->(1,0)=2, 2:(1,0)->(0,1)=2. Total 4. No conflict.
		{"conflict + manhattan", []int{2, 1, 3, 4}, 2, 2, 4}, // Manhattan: 2:(0,0)->(0,1)=1, 1:(0,1)->(0,0)=1. Total 2. Conflict: 2-1 in row 0 -> +2. Total 4.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			goal := StandardGoal(tt.rows, tt.cols)
			if got := calculateHeuristic(tt.board, goal, tt.rows, tt.cols); got != tt.want {
				t.Errorf("%s: calculateHeuristic() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
