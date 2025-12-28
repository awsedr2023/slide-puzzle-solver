package solver

import (
	"testing"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name    string
		board   []int
		rows    int
		cols    int
		wantErr error
	}{
		{
			name:    "valid 2x2",
			board:   []int{1, 2, 3, 4},
			rows:    2,
			cols:    2,
			wantErr: nil,
		},
		{
			name:    "empty board",
			board:   []int{},
			rows:    2,
			cols:    2,
			wantErr: ErrEmptyBoard,
		},
		{
			name:    "invalid size (rows=0)",
			board:   []int{1, 2, 3, 4},
			rows:    0,
			cols:    2,
			wantErr: ErrInvalidSize,
		},
		{
			name:    "invalid size (rows=1)",
			board:   []int{1, 2},
			rows:    1,
			cols:    2,
			wantErr: ErrInvalidSize,
		},
		{
			name:    "invalid size (cols=1)",
			board:   []int{1, 2},
			rows:    2,
			cols:    1,
			wantErr: ErrInvalidSize,
		},
		{
			name:    "size mismatch",
			board:   []int{1, 2, 3},
			rows:    2,
			cols:    2,
			wantErr: ErrSizeMismatch,
		},
		{
			name:    "invalid element (0 included)",
			board:   []int{0, 1, 2, 3},
			rows:    2,
			cols:    2,
			wantErr: ErrInvalidElement,
		},
		{
			name:    "invalid element (duplicate)",
			board:   []int{1, 1, 3, 4},
			rows:    2,
			cols:    2,
			wantErr: ErrInvalidElement,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validate(tt.board, tt.rows, tt.cols)
			if err != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
