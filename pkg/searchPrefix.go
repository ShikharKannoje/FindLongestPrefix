package pkg

import (
	"strings"
)

func SearchLongPrefix(num string, node *internalBST.Node, maxPreLen int, prefix string) (bool, int, string) {

	if node == nil {
		return false, maxPreLen, prefix
	}
	check := strings.HasPrefix(num, node.Data)
	if check == true {
		prefix = node.Data
		maxPreLen = len(node.Data)
	}
	op := strings.Compare(num, node.Data)
	if op > 0 {
		return SearchLongPrefix(num, node.Right, maxPreLen, prefix)
	}
	if op < 0 {
		return SearchLongPrefix(num, node.Left, maxPreLen, prefix)
	}
	if op == 0 {
		maxPreLen = len(num)
		return true, maxPreLen, prefix
	}
	return true, maxPreLen, prefix
}
