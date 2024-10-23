package main

import (
	"fmt"
)

type BTS struct {
	Value int
	Left  *BTS
	Right *BTS
}

func Add(node *BTS, element int) *BTS {
	if node == nil {
		return &BTS{Value: element}
	}

	if element < node.Value {
		node.Left = Add(node.Left, element)
	} else {
		node.Right = Add(node.Right, element)
	}
	return node
}

func Print(node *BTS) {
	if node != nil {
		Print(node.Left)
		fmt.Print(node.Value, " ")
		Print(node.Right)
	}
}

func IsExist(node *BTS, element int) bool {
	if node == nil {
		return false
	}
	if node.Value == element {
		return true
	} else if element < node.Value {
		return IsExist(node.Left, element)
	} else {
		return IsExist(node.Right, element)
	}
}

func Delete(node *BTS, element int) *BTS {
	if node == nil {
		return nil
	}
	if element < node.Value {
		node.Left = Delete(node.Left, element)
	} else if element > node.Value {
		node.Right = Delete(node.Right, element)
	} else {
		if node.Left == nil && node.Right == nil {
			return nil
		} else if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		}
		minRight := find(node.Right)
		node.Value = minRight.Value
		node.Right = Delete(node.Right, minRight.Value)
	}
	return node
}

func find(node *BTS) *BTS {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}

func main() {
	var tree *BTS

	for _, element := range []int{1, 2, 3, 5, 6, 7, 9} {
		tree = Add(tree, element)
	}
	fmt.Println(IsExist(tree, 10))
	tree = Delete(tree, 10)
	Print(tree)
}
