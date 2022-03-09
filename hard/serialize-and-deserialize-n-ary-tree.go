// https://leetcode.com/problems/serialize-and-deserialize-n-ary-tree
package hard

import (
	"fmt"
	"strconv"
)

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type Node struct {
	Val      int
	Children []*Node
}

type Codec struct {
}

func Constructor() *Codec {
	var pa *Codec
	pa = new(Codec)
	return pa
}

func (this *Codec) serializeHelper(root *Node, curIndex int, totalSiblings int) string {
	if root == nil {
		return ""
	}
	v := root.Val
	var s string
	if curIndex == 0 {
		s += "["
	}
	s += strconv.Itoa(v)
	noOfChildren := len(root.Children)
	if noOfChildren != 0 {
		s += " "
	}
	for i, c := range root.Children {
		s += this.serializeHelper(c, i, noOfChildren)
	}
	if curIndex == totalSiblings-1 {
		s += "]"
	} else {
		s += " "
	}
	return s
}

func (this *Codec) serialize(root *Node) string {
	if root == nil {
		return ""
	}
	s := this.serializeHelper(root, 0, 1)
	fmt.Printf("%v\n", s)
	return s
}

func (this *Codec) deserializeHelper(data string, curIndex int) ([]*Node, int) {
	ci := curIndex
	var sib [](*Node)
	cur_sid := 0
	if data[ci] == '[' {
		// got its value
		ci += 1
		v := 0
		fmt.Printf("%v\n", ci)
		for data[ci] != ' ' && data[ci] != ']' {
			fmt.Printf("%v === %v\n", ci, data[ci:ci+1])
			v = (v * 10) + int(data[ci]-'0')
			ci += 1
			fmt.Println(data[ci] != ' ')
		}
		fmt.Printf("exit %v\n", ci)
		var pa *Node
		pa = new(Node)
		(*pa).Val = v
		sib = append(sib, pa)
	}
	for data[ci] != ']' {
		if data[ci] == ' ' {
			ci += 1
		}
		if data[ci] == '[' {
			(*sib[cur_sid]).Children, ci = this.deserializeHelper(data, ci)
		} else {
			// is a sibling
			v := 0
			for data[ci] != ' ' && data[ci] != ']' {
				v = (v * 10) + int(data[ci]-'0')
				ci += 1
			}
			var pa *Node
			pa = new(Node)
			pa.Val = v
			sib = append(sib, pa)
			cur_sid += 1
		}
	}
	return sib, ci + 1
}

func (this *Codec) deserialize(data string) *Node {
	if len(data) == 0 {
		return nil
	}
	a, _ := this.deserializeHelper(data, 0)
	return a[0]
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
