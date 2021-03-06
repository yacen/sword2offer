package main

import (
	"fmt"
	. "github.com/yacen/sword2offer/tree"
)

// 输入两棵二叉树A，B，判断B是不是A的子结构。（ps：我们约定空树不是任意一个树的子结构）

/*
		1
	2		3
4	5		6
7           8

		8
	7  		6


*/

func main() {
	n1 := &Node{Data: 1}
	n2 := &Node{Data: 2}
	n3 := &Node{Data: 3}
	n4 := &Node{Data: 4}
	n5 := &Node{Data: 5}
	n6 := &Node{Data: 6}
	n7 := &Node{Data: 7}
	n8 := &Node{Data: 8}

	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5
	n3.Left = n6
	n4.Left = n7
	n6.Left = n8

	fmt.Println(HasSubtree(n1, n8))

	n61 := &Node{Data: 6}
	n71 := &Node{Data: 7}
	n81 := &Node{Data: 8}
	n81.Left = n71
	n81.Right = n61
	fmt.Println(HasSubtree(n1, n81))

}

func HasSubtree(pRoot1 *Node, pRoot2 *Node) bool {
	if pRoot1 == nil || pRoot2 == nil {
		return false
	}
	if pRoot1 == pRoot2 {
		return true
	}
	if CheckTree(pRoot1, pRoot2) {
		return true
	} else {
		if HasSubtree(pRoot1.Left, pRoot2) || HasSubtree(pRoot1.Right, pRoot2) {
			return true
		}
		return false
	}
}

func CheckTree(pRoot1 *Node, pRoot2 *Node) bool {
	if pRoot2 == nil {
		return true
	}
	if pRoot1 == nil {
		return false
	}
	if pRoot1.Data != pRoot2.Data {
		return false
	} else {
		return CheckTree(pRoot1.Left, pRoot2.Left) && CheckTree(pRoot1.Right, pRoot2.Right)
	}
}
