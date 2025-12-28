package solver

import (
	"reflect"
	"testing"
)

func TestNewNode(t *testing.T) {
	input := []int{1, 2, 4, 3}
	n := newNode(input, 2, 2)

	if !reflect.DeepEqual(n.board, input) {
		t.Errorf("newNode().board = %v, want %v", n.board, input)
	}
	// 4 is the blank tile (len=4)
	// index of 4 in {1, 2, 4, 3} is 2
	if n.blankIdx != 2 {
		t.Errorf("newNode().blankIdx = %v, want 2", n.blankIdx)
	}
	if n.rows != 2 {
		t.Errorf("newNode().rows = %v, want 2", n.rows)
	}
	if n.cols != 2 {
		t.Errorf("newNode().cols = %v, want 2", n.cols)
	}
	if n.cost != 0 {
		t.Errorf("newNode().cost = %v, want 0", n.cost)
	}
	if n.parent != nil {
		t.Error("newNode().parent should be nil")
	}
}

func TestNode_Has(t *testing.T) {
	n := newNode([]int{1, 2, 3, 4}, 2, 2)
	if !n.has([]int{1, 2, 3, 4}) {
		t.Error("has() should return true for identical board")
	}
	if n.has([]int{1, 2, 4, 3}) {
		t.Error("has() should return false for different board")
	}
}

func TestNode_Copy(t *testing.T) {
	n := newNode([]int{1, 2, 3, 4}, 2, 2)
	c := n.copy()

	if !reflect.DeepEqual(n.board, c.board) {
		t.Error("copy() should create identical board")
	}
	if n.blankIdx != c.blankIdx {
		t.Error("copy() should copy blankIdx")
	}

	// Modify copy to ensure deep copy
	c.board[0] = 99
	if n.board[0] == 99 {
		t.Error("copy() should be deep copy")
	}
}

func TestNode_CanMove(t *testing.T) {
	// 2x2 board
	// 0 1
	// 2 3
	// Indices:
	// 0 1
	// 2 3

	// Blank at top-left (0)
	n := newNode([]int{4, 1, 2, 3}, 2, 2) // 4 is blank
	if n.canMoveUp() {
		t.Error("Should not be able to move up from top row")
	}
	if n.canMoveLeft() {
		t.Error("Should not be able to move left from first column")
	}
	if !n.canMoveDown() {
		t.Error("Should be able to move down")
	}
	if !n.canMoveRight() {
		t.Error("Should be able to move right")
	}

	// Blank at bottom-right (3)
	n = newNode([]int{1, 2, 3, 4}, 2, 2) // 4 is blank
	if !n.canMoveUp() {
		t.Error("Should be able to move up")
	}
	if !n.canMoveLeft() {
		t.Error("Should be able to move left")
	}
	if n.canMoveDown() {
		t.Error("Should not be able to move down from bottom row")
	}
	if n.canMoveRight() {
		t.Error("Should not be able to move right from last column")
	}
}

func TestNode_Move(t *testing.T) {
	// Start:
	// 1 2
	// 3 4(blank)
	n := newNode([]int{1, 2, 3, 4}, 2, 2)

	// Move Up
	// Expected:
	// 1 4(blank)
	// 3 2
	n.moveUp()
	expected := []int{1, 4, 3, 2}
	if !reflect.DeepEqual(n.board, expected) {
		t.Errorf("moveUp() failed. got %v, want %v", n.board, expected)
	}
	if n.blankIdx != 1 {
		t.Errorf("moveUp() blankIdx = %d, want 1", n.blankIdx)
	}
}

func TestNode_ChildNodes(t *testing.T) {
	n := newNode([]int{1, 2, 3, 4}, 2, 2)
	child := n.upNode()

	if child.parent != n {
		t.Error("upNode() should set parent")
	}
	if child.cost != n.cost+1 {
		t.Error("upNode() should increment cost")
	}
	// Check if original node is unmodified
	if !reflect.DeepEqual(n.board, []int{1, 2, 3, 4}) {
		t.Error("upNode() should not modify original node")
	}
}
