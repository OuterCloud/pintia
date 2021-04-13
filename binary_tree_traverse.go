package pintia

// BinaryTree 二叉树
type BinaryTree struct {
	RootContent string
	LeftNode    *BinaryTree
	RightNode   *BinaryTree
}

// CreateBinaryTree 创建二叉树
func CreateBinaryTree() BinaryTree {
	return BinaryTree{
		RootContent: "1-1",
		LeftNode: &BinaryTree{
			RootContent: "2-1",
			LeftNode: &BinaryTree{
				RootContent: "3-1",
				LeftNode: &BinaryTree{
					RootContent: "4-1",
					LeftNode:    &BinaryTree{},
					RightNode:   &BinaryTree{},
				},
				RightNode: &BinaryTree{
					RootContent: "4-2",
					LeftNode:    &BinaryTree{},
					RightNode:   &BinaryTree{},
				},
			},
			RightNode: &BinaryTree{
				RootContent: "3-2",
				LeftNode: &BinaryTree{
					RootContent: "4-3",
					LeftNode:    &BinaryTree{},
					RightNode:   &BinaryTree{},
				},
				RightNode: &BinaryTree{
					RootContent: "4-4",
					LeftNode:    &BinaryTree{},
					RightNode:   &BinaryTree{},
				},
			},
		},
		RightNode: &BinaryTree{
			RootContent: "2-2",
			LeftNode: &BinaryTree{
				RootContent: "3-3",
				LeftNode: &BinaryTree{
					RootContent: "4-5",
					LeftNode:    &BinaryTree{},
					RightNode:   &BinaryTree{},
				},
				RightNode: &BinaryTree{
					RootContent: "4-6",
					LeftNode:    &BinaryTree{},
					RightNode:   &BinaryTree{},
				},
			},
			RightNode: &BinaryTree{
				RootContent: "3-4",
				LeftNode: &BinaryTree{
					RootContent: "4-7",
					LeftNode:    &BinaryTree{},
					RightNode:   &BinaryTree{},
				},
				RightNode: &BinaryTree{
					RootContent: "4-8",
					LeftNode:    &BinaryTree{},
					RightNode:   &BinaryTree{},
				},
			},
		},
	}
}
