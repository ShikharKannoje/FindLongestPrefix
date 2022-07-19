package pkg

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"

	"github.com/ShikharKannoje/prefixSearch/pkg/internalBST"
)

const FilePath = "../configs/"

func InitiateSearchFor(inputStringOfString []string) {
	data, err := ioutil.ReadFile(FilePath + "prefixList.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	slicedata := strings.Split(string(data), "\n")
	sort.Strings(slicedata)
	slicedata = slicedata[1:]
	bst := internalBST.CreateBST(slicedata)
	for _, inputString := range inputStringOfString {
		fmt.Println(searchLongestPrefix(inputString, bst, 0, ""))
	}

}

func searchLongestPrefix(inputStr string, bstNode *internalBST.Node, maxPreLen int, prefix string) string {
	_, _, longestPrefix := matchPrefix(inputStr, bstNode, maxPreLen, prefix)
	return longestPrefix
}

func matchPrefix(inputStr string, bstNode *internalBST.Node, maxPreLen int, prefix string) (bool, int, string) {

	if bstNode == nil {
		return false, maxPreLen, prefix
	}
	check := strings.HasPrefix(inputStr, bstNode.Data)
	if check {
		prefix = bstNode.Data
		maxPreLen = len(bstNode.Data)
	}
	op := strings.Compare(inputStr, bstNode.Data)
	if op > 0 {
		return matchPrefix(inputStr, bstNode.Right, maxPreLen, prefix)
	}
	if op < 0 {
		return matchPrefix(inputStr, bstNode.Left, maxPreLen, prefix)
	}
	if op == 0 {
		maxPreLen = len(inputStr)
		return true, maxPreLen, prefix
	}
	return true, maxPreLen, prefix
}
