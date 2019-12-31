package dict

import (
	"errors"
	"fmt"
)

// TrieNode is the the node structure of trie tree.
type TrieNode struct {
	Char byte
	Word string
	Sep  bool
	Next map[byte]*TrieNode
}

// Dict is a dictionary implemented by trie tree.
type Dict struct {
	Root *TrieNode
}

// NewDictionary creates a new Dict.
func NewDictionary() *Dict {
	dict := Dict{}
	dict.Root = &TrieNode{
		Char: byte('#'),
		Next: map[byte]*TrieNode{},
	}
	return &dict
}

// Find finds a string. Return nil if not found.
func (d *Dict) Find(str string) *TrieNode {
	node, found := find(d.Root, "#"+str)
	if !found {
		return nil
	}
	if !node.Sep {
		return nil
	}
	return node
}

func find(currNode *TrieNode, subStr string) (*TrieNode, bool) {
	nextC, nextS, err := next(subStr)
	if err != nil {
		// No next sub string exists. Match the end of string
		// There has 2 conditions:
		// 1. Match totally
		// 2. Match the end of string, but this node is not the end of word.
		//    Querying string is the sub string of the current node (to the leaf).
		//    Ex: curr=app[le], query=app
		return currNode, true
	}

	nextN, ok := currNode.Next[nextC]
	if !ok {
		// Find no next node. String isn't matched.
		// current node is the sub string of the querying one.
		// Ex: curr=app, quer=apple
		return currNode, false
	}

	return find(nextN, nextS)
}

// Predict s
func (d *Dict) Predict(str string) (o []string) {
	node, found := find(d.Root, "#"+str)
	if !found {
		return
	}
	// Expand all the sub-tree by depth-first.
	o = append(o, depthFirstTraverse(node)...)
	return
}

func depthFirstTraverse(curr *TrieNode) (o []string) {
	if len(curr.Next) == 0 {
		o = append(o, curr.Word)
		return
	}

	if curr.Sep {
		o = append(o, curr.Word)
	}

	for _, next := range curr.Next {
		o = append(o, depthFirstTraverse(next)...)
	}
	return
}

// Add adds a string into the trie tree.
func (d *Dict) Add(str string) {
	closedN, found := find(d.Root, "#"+str)
	if found {
		// Found. Mark the `Sep` flag.
		closedN.Sep = true
		return
	}

	// Not found. Create the new node.
	var (
		subStr   = str[len(closedN.Word):len(str)]
		word     = closedN.Word
		currNode = closedN
	)
	for _, c := range []byte(subStr) {
		word = word + string(c)
		currNode.Next[c] = &TrieNode{
			Char: c,
			Word: word,
			Next: map[byte]*TrieNode{},
		}
		currNode = currNode.Next[c]
	}
	currNode.Sep = true
	return
}

// next returns 1) the next char 2) the substring start from next char.
// ex: INPUT:  hello
//     OUTPUT: e, ello, error
func next(str string) (byte, string, error) {
	if len(str) < 2 {
		return 0, "", errors.New("no next")
	}
	return str[1], str[1:len(str)], nil
}

// Dump dumps the whole tree.
func (d *Dict) Dump() []string {
	o := depthFirstTraverse(d.Root)
	for i, v := range o {
		fmt.Printf("%8v %v\n", i+1, v)
	}
	return o
}
