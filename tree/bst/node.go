// 二叉搜索树性质
//   设 x 是二叉搜索树的一个结点。如果 y 是 x 左子树的一个结点，那么 y.key <= x.key。
//   如果 y 是 x 的右子树的一个结点，那么 y.key >= x.key。
//          4
//        .   .
//      2       5
//    .   .
//  1       3
//
// 树的深度优先遍历(Depth First): 根据输出的子树根的关键字位于其左子树的关键字和右子树的关键字之间, 遍历
//     算法有
//		- "先序遍历"(Preorder) : 4, 2, 1, 3, 5
//		- "中序遍历"(Inorder)  : 1, 2, 3, 4, 5
//		- "后续遍历"(Postorder): 1, 3, 2, 5, 4
// 树的广度优先遍历(Breadth Fist)或"层次遍历" (Level Order): 4, 2, 5, 1, 3
package bst

import (
	"fmt"
)

type Node struct {
	p     *Node
	left  *Node
	right *Node
	key   int
}

func NewNode(key int) *Node {
	return &Node{key: key}
}

func (x *Node) Key() int {
	return x.key
}

// PREORDER-TREE-WALK(x)
//   if x != NIL
//      print x
//		PREORDER-TREE-WALK(x.left)
//		PREORDER-TREE-WALK(x.right)
func (x *Node) PREORDER_WALK(walker func(*Node)) {
	if x != nil {
		walker(x)
		x.left.PREORDER_WALK(walker)
		x.right.PREORDER_WALK(walker)
	}
}

// INORDER-TREE-WALK(x)
//   if x != NIL
//     INORDER-TREE-WALK(x.left)
//     print x
//     INORDER-TREE-WALK(x.right)
func (x *Node) INORDER_WALK(walker func(*Node)) {
	if x != nil {
		if x.left != nil {
			x.left.INORDER_WALK(walker)
		}

		if walker != nil {
			walker(x)
		} else {
			fmt.Printf(" %d", x.key)
		}

		if x.right != nil {
			x.right.INORDER_WALK(walker)
		}
	}
}

// POSTORDER-TREE-WALK(x)
//   if x != NIL
//      print x
//		POSTORDER-TREE-WALK(x.left)
//		POSTORDER-TREE-WALK(x.right)
func (x *Node) POSTORDER_WALK(walker func(*Node)) {
	if x != nil {
		walker(x)
		x.left.POSTORDER_WALK(walker)
		x.right.POSTORDER_WALK(walker)
	}
}

// LEVELORDER-TREE-WALK(x)
func (x *Node) LEVELORDER_WALK(walker func(*Node)) {
	q := &nodeq{}
	for z := x; z != nil; z = q.Deq() {
		walker(z)
		q.Enq(z.left)
		q.Enq(z.right)
	}
}

// TREE-MINIMUM(x)
//   while x.left != NIL
//     x = x.left
//   return x
func (x *Node) MINIMUM() *Node {
	for x.left != nil {
		x = x.left
	}
	return x
}

// TREE-MAXIMUM(x)
// while x.right != NIL
//   x = x.right
// return x
func (x *Node) MAXIMUM() *Node {
	for x.right != nil {
		x = x.right
	}
	return x
}
