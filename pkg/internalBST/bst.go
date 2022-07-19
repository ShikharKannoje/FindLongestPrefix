package internal

import (
	"fmt"
	"strings"
)

type Node struct {
	Data  string
	Left  *Node
	Right *Node
}

func createBST(array []string) *Node {

	root := constructBST(array, 0, len(array)-1)
	//	display(root)
	return root
}

func display(node *Node) {
	if node == nil {
		return
	}

	display(node.Left)

	display(node.Right)
	fmt.Println(node.Data)
}

func constructBST(arr []string, low, high int) *Node {
	if low > high {
		return nil
	}

	mid := (low + high) / 2
	data := arr[mid]
	lc := constructBST(arr, low, mid-1)
	rc := constructBST(arr, mid+1, high)
	data = strings.Trim(data, string(byte(13)))
	node := Node{Data: data, Left: lc, Right: rc}
	return &node
}
