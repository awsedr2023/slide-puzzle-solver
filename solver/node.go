package solver

// node represents a state in the search space.
type node struct {
	board    []int
	blankIdx int
	rows     int
	cols     int
	cost     int   // g(n): cost from start
	parent   *node // parent node to reconstruct the path
}

const (
	up    = 0
	down  = 1
	left  = 2
	right = 3
)

// newNode creates a new node with the given board configuration, rows and columns.
// It initializes the blank tile position based on the input board.
//
// Example:
//
//	n := newNode([]int{1, 2, 3, 4}, 2, 2)
func newNode(input []int, rows, cols int) *node {
	board := make([]int, len(input))
	copy(board, input)

	blankIdx := -1
	for i, v := range board {
		if v == len(board) {
			blankIdx = i
			break
		}
	}

	return &node{
		board:    board,
		blankIdx: blankIdx,
		rows:     rows,
		cols:     cols,
	}
}

// has checks if the node's board matches the given target board.
//
// Example:
//
//	n.has([]int{1, 2, 3, 4})
func (n *node) has(board []int) bool {
	for i, v := range board {
		if v != n.board[i] {
			return false
		}
	}
	return true
}

// upNode creates a new child node by moving the blank tile up.
// It returns a new node with updated board, parent, and cost.
func (n *node) upNode() *node {
	next := n.copy()
	next.moveUp()
	next.parent = n
	next.cost = n.cost + 1
	return &next
}

// downNode creates a new child node by moving the blank tile down.
// It returns a new node with updated board, parent, and cost.
func (n *node) downNode() *node {
	next := n.copy()
	next.moveDown()
	next.parent = n
	next.cost = n.cost + 1
	return &next
}

// rightNode creates a new child node by moving the blank tile right.
// It returns a new node with updated board, parent, and cost.
func (n *node) rightNode() *node {
	next := n.copy()
	next.moveRight()
	next.parent = n
	next.cost = n.cost + 1
	return &next
}

// leftNode creates a new child node by moving the blank tile left.
// It returns a new node with updated board, parent, and cost.
func (n *node) leftNode() *node {
	next := n.copy()
	next.moveLeft()
	next.parent = n
	next.cost = n.cost + 1
	return &next
}

// copy creates a deep copy of the current node.
func (n *node) copy() node {
	board := make([]int, len(n.board))
	copy(board, n.board)
	return node{
		board:    board,
		blankIdx: n.blankIdx,
		rows:     n.rows,
		cols:     n.cols,
		cost:     n.cost,
		parent:   n.parent,
	}
}

// moveUp moves the blank tile up in the current node.
// It updates the board and blank tile index.
func (n *node) moveUp() {
	n.moveBlank(up)
}

// moveDown moves the blank tile down in the current node.
// It updates the board and blank tile index.
func (n *node) moveDown() {
	n.moveBlank(down)
}

// moveLeft moves the blank tile left in the current node.
// It updates the board and blank tile index.
func (n *node) moveLeft() {
	n.moveBlank(left)
}

// moveRight moves the blank tile right in the current node.
// It updates the board and blank tile index.
func (n *node) moveRight() {
	n.moveBlank(right)
}

// canMoveUp checks if the blank tile can be moved up.
// It returns true if the blank tile is not in the top row.
func (n *node) canMoveUp() bool {
	return n.blankIdx >= n.cols
}

// canMoveDown checks if the blank tile can be moved down.
// It returns true if the blank tile is not in the bottom row.
func (n *node) canMoveDown() bool {
	return n.blankIdx < len(n.board)-n.cols
}

// canMoveLeft checks if the blank tile can be moved left.
// It returns true if the blank tile is not in the first column.
func (n *node) canMoveLeft() bool {
	return n.blankIdx%n.cols != 0
}

// canMoveRight checks if the blank tile can be moved right.
// It returns true if the blank tile is not in the last column.
func (n *node) canMoveRight() bool {
	return n.blankIdx%n.cols != n.cols-1
}

// moveBlank moves the blank tile in the specified direction.
// It swaps the blank tile with the adjacent tile and updates the blank index.
func (n *node) moveBlank(dir int) {
	switch dir {
	case up:
		n.swap(n.blankIdx, n.blankIdx-n.cols)
		n.blankIdx -= n.cols
	case down:
		n.swap(n.blankIdx, n.blankIdx+n.cols)
		n.blankIdx += n.cols
	case left:
		n.swap(n.blankIdx, n.blankIdx-1)
		n.blankIdx--
	case right:
		n.swap(n.blankIdx, n.blankIdx+1)
		n.blankIdx++
	}
}

// swap swaps the values at indices i and j in the board.
func (n *node) swap(i, j int) {
	n.board[i], n.board[j] = n.board[j], n.board[i]
}
