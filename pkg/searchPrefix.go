package pkg

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"

	"github.com/ShikharKannoje/prefixSearch/pkg/internalBST"
)

const FilePath = "../configs/"

//InitiateSearchFor initiates the operation by first creating a BST of the worklist and then searching prefix
func InitiateSearchFor(inputStringOfString []string) {
	//reading the prefixList.txt kept in the config folder
	data, err := ioutil.ReadFile(FilePath + "prefixList.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	slicedata := strings.Split(string(data), "\n")
	sort.Strings(slicedata)
	slicedata = slicedata[1:] // this formating required based on the sampleFile provided
	//creating the Binary Search Tree of the prefix list provided above
	bst := internalBST.CreateBST(slicedata)
	for _, inputString := range inputStringOfString {

		//searching the longest prefix
		fmt.Println(searchLongestPrefix(inputString, bst, 0, ""))
	}

}

//searchLongestPrefix calls internal matchPrefix method which does the actual work
func searchLongestPrefix(inputStr string, bstNode *internalBST.Node, maxPreLen int, prefix string) string {
	_, _, longestPrefix := matchPrefix(inputStr, bstNode, maxPreLen, prefix)
	return longestPrefix
}

//matchPrefix returns
//boolean true if the prefix == inputstring
//number of character matched in prefix
//longest prefix matched
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
