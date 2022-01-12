package main

import (
	"errors"
	"fmt"
)

var (
	ErrDuplicatedNode error = errors.New("bst: found duplicated value on tree")
	ErrNodeNotFound   error = errors.New("bst: node not found")
)

type node struct {
	height, value int
	left, right   *node
}

func (n *node) balanceFactor() int {
	if n == nil {
		return 0
	}

	return n.left.Height() - n.right.Height()
}

func (n *node) updateHeight() {
	max := func(a, b int) int {
		if a > b {
			return a
		}

		return b
	}
	n.height = max(n.left.Height(), n.right.Height()) + 1
}

func (n *node) Height() int {
	if n == nil {
		return 0
	}

	return n.height
}

func (n *node) Value() int {
	return n.value
}

func (n *node) Left() *node {
	return n.left
}

func (n *node) Right() *node {
	return n.right
}

func newNode(val int) *node {
	return &node{
		height: 1,
		value:  val,
		left:   nil,
		right:  nil,
	}
}

func insertNode(node *node, val int) (*node, error) {
	// if there's no node, create one
	if node == nil {
		return newNode(val), nil
	}

	// if there's duplicated node returns error
	if node.value == val {
		return nil, ErrDuplicatedNode
	}

	// if value is greater than current node's value, insert to the right
	if val > node.value {
		right, err := insertNode(node.right, val)

		if err != nil {
			return nil, err
		}

		node.right = right
	}

	// if value is less than current node's value, insert to the left
	if val < node.value {
		left, err := insertNode(node.left, val)

		if err != nil {
			return nil, err
		}

		node.left = left
	}

	return rotateInsert(node, val), nil
}

func removeNode(node *node, val int) (*node, error) {
	if node == nil {
		return nil, ErrNodeNotFound
	}

	if val > node.value {
		right, err := removeNode(node.right, val)
		if err != nil {
			return nil, err
		}

		node.right = right
	} else if val < node.value {
		left, err := removeNode(node.left, val)
		if err != nil {
			return nil, err
		}

		node.left = left
	} else {
		if node.left != nil && node.right != nil {
			// has 2 children

			// find the successor
			successor := greatest(node.left)
			value := successor.value

			// remove the successor
			left, err := removeNode(node.left, value)
			if err != nil {
				return nil, err
			}
			node.left = left

			// copy the successor value to the current node
			node.value = value
		} else if node.left != nil || node.right != nil {
			// has 1 child
			// move the child position to the current node
			if node.left != nil {
				node = node.left
			} else {
				node = node.right
			}
		} else if node.left == nil && node.right == nil {
			// has no child
			// simply remove the node
			node = nil
		}
	}

	if node == nil {
		return nil, nil
	}

	return rotateDelete(node), nil
}

func findNode(node *node, val int) *node {
	if node == nil {
		return nil
	}

	// if the node is found, return the node
	if node.value == val {
		return node
	}

	// if value is greater than current node's value, search recursively to the right
	if val > node.value {
		return findNode(node.right, val)
	}

	// if value is less than current node's value, search recursively to the left
	if val < node.value {
		return findNode(node.left, val)
	}

	return nil
}

func rotateInsert(node *node, val int) *node {
	// update the height on every insertion
	node.updateHeight()

	// bFactor will tell you which side the weight is on
	bFactor := node.balanceFactor()

	// linearly to the left
	if bFactor > 1 && val < node.left.value {
		return rightRotate(node)
	}

	// linearly to the right
	if bFactor < -1 && val > node.right.value {
		return leftRotate(node)
	}

	// less than symbol
	if bFactor > 1 && val > node.left.value {
		node.left = leftRotate(node.left)
		return rightRotate(node)
	}

	// greater than symbol
	if bFactor < -1 && val < node.right.value {
		node.right = rightRotate(node.right)
		return leftRotate(node)
	}

	return node
}

func rotateDelete(node *node) *node {
	node.updateHeight()
	bFactor := node.balanceFactor()

	// linearly to the left
	if bFactor > 1 && node.left.balanceFactor() >= 0 {
		return rightRotate(node)
	}

	// less than symbol
	if bFactor > 1 && node.left.balanceFactor() < 0 {
		node.left = leftRotate(node.left)
		return rightRotate(node)
	}

	// linearly to the right
	if bFactor < -1 && node.right.balanceFactor() <= 0 {
		return leftRotate(node)
	}

	// greater than symbol
	if bFactor < -1 && node.right.balanceFactor() > 0 {
		node.right = rightRotate(node.right)
		return leftRotate(node)
	}

	return node
}

func rightRotate(x *node) *node {
	y := x.left
	t := y.right

	y.right = x
	x.left = t

	x.updateHeight()
	y.updateHeight()

	return y
}

func leftRotate(x *node) *node {
	y := x.right
	t := y.left

	y.left = x
	x.right = t

	x.updateHeight()
	y.updateHeight()

	return y
}

func greatest(node *node) *node {
	if node == nil {
		return nil
	}

	if node.right == nil {
		return node
	}

	return greatest(node.right)
}

func traverse(node *node) {
	// exit condition
	if node == nil {
		return
	}

	fmt.Println(node.value)
	traverse(node.left)
	traverse(node.right)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
