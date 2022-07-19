package internalBST

import (
	"fmt"
	"strings"
)

//Node is the structure of BST Node
type Node struct {
	Data  string
	Left  *Node
	Right *Node
}

//CreateBST takes in the sorted array of prefix values and returns consturcted BST root-Node
func CreateBST(array []string) *Node {

	root := ConstructBST(array, 0, len(array)-1)
	return root
}

//Display is not being used here but I wrote this just to be view the BST
func Display(node *Node) {
	if node == nil {
		return
	}

	Display(node.Left)

	Display(node.Right)
	fmt.Println(node.Data)
}

//ConstructBST recursivly creates the BST
func ConstructBST(arr []string, low, high int) *Node {
	if low > high {
		return nil
	}

	mid := (low + high) / 2
	data := arr[mid]
	lc := ConstructBST(arr, low, mid-1)
	rc := ConstructBST(arr, mid+1, high)
	data = strings.Trim(data, string(byte(13)))
	node := Node{Data: data, Left: lc, Right: rc}
	return &node
}
