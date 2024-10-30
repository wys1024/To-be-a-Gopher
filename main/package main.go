package main

type ListNode struct {
	Val  int
	Right *ListNode
	Left *ListNode
}

func main() {
	//测试用例
	L := &ListNode{1, nil, nil}

	//输出
	checkTree(L, L)
}

func checkTree(h1, h2 *ListNode) bool {
	//都为nil,返回true
	if h1 == nil || h2 == nil {
		return h1 == h2
	}

	//都不为nil,比较值
	if h1.Val == h2.Val && checkTree(h1.Right, h2.Left) && checkTree(h1.Left, h2.Right) {
		return true
	} else {
		return false
	}
}
