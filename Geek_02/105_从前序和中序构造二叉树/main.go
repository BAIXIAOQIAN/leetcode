package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{
		Val:   preorder[0],
		Left:  nil,
		Right: nil,
	}

	var s int
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			s = i
			break
		}
	}
	root.Left = buildTree(preorder[1:s+1], inorder[:s+1])
	root.Right = buildTree(preorder[s+1:], inorder[s+1:])
	return root
}
