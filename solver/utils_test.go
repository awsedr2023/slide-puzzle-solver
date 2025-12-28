package solver

import (
	"reflect"
	"testing"
)

func TestStandardGoal(t *testing.T) {
	tests := []struct {
		name string
		rows int
		cols int
		want []int
	}{
		{
			name: "2x2",
			rows: 2,
			cols: 2,
			want: []int{1, 2, 3, 4},
		},
		{
			name: "3x3",
			rows: 3,
			cols: 3,
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "2x3",
			rows: 2,
			cols: 3,
			want: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StandardGoal(tt.rows, tt.cols); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StandardGoal() = %v, want %v", got, tt.want)
			}
		})
	}
}
