package solver

// StandardGoal creates a standard goal state for a puzzle of the given rows and columns.
// It returns a slice containing numbers from 1 to rows*cols in ascending order.
//
// Example:
//
//	StandardGoal(2, 2) // returns []int{1, 2, 3, 4}
func StandardGoal(rows, cols int) []int {
	goal := make([]int, rows*cols)
	for i := 0; i < rows*cols; i++ {
		goal[i] = i + 1
	}
	return goal
}
