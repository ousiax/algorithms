// 二叉搜索树性质
//   设 x 是二叉搜索树的一个结点。如果 y 是 x 左子树的一个结点，那么 y.key <= x.key。
//   如果 y 是 x 的右子树的一个结点，那么 y.key >= x.key。
//
// 树的深度优先遍历: 根据输出的子树根的关键字位于其左子树的关键字和右子树的关键字之间, 遍历
//     算法有"中序遍历", "先序遍历", "后续遍历"
// 树的广度优先遍历：层次遍历, 根结点->左子树->右子树
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

type Tree struct {
	root *Node
}

func (t *Tree) String() string {
	keys := make([]int, 0)
	t.INORDER_WALK(func(n *Node) {
		keys = append(keys, n.key)
	})
	return fmt.Sprintf("%v", keys)
}

// INORDER-TREE-WALK(x)
//   if x != NIL
//     INORDER-TREE-WALK(x.left)
//     print x
//     INORDER-TREE-WALK(x.right)
func (t *Tree) INORDER_WALK(walker func(*Node)) {
	if t != nil && t.root != nil {
		t.root.INORDER_WALK(walker)
	}
}

// TREE-SEARCH(x, k)
//   if == NIL or k == x.key
//     return x
//   if x < x.key
//     return TREE-SEARCH(x.left, k)
//   else
//     return TREE-SEARCH(x.right, k)
// - - -
// ITERATIVE-TREE-SEARCH(x, k)
//   while x != NIL and k != x.key
//     if x < x.key
//       x = x.left
//     else
//       x = x.right
//   return x
func (t *Tree) SEARCH(k int) *Node {
	x := t.root
	for x != nil && k != x.key {
		if k < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}
	return x
}

// TREE-MINIMUM(x)
//   while x.left != NIL
//     x = x.left
//   return x
func (t *Tree) MINIMUM() *Node {
	x := t.root
	return x.MINIMUM()
}

// TREE-MAXIMUM(x)
// while x.right != NIL
//   x = x.right
// return x
func (t *Tree) MAXIMUM() *Node {
	x := t.root
	return x.MAXIMUM()
}

// TREE-SUCCESSOR(x)
//   if x.right != NIL
//     return TREE-MINIMUM(x.right)
//   y = x.p
//   while y != NIL and x == y.right
//     x = y
//     y = y.p
//   return y
// 给定一棵二叉搜索树中的一个结点，有时候需要按中序遍历的次序查找它的后继。
// 如果所有的关键字互不相同，则一个结点的后继是大于 x.key 的最小关键字的结点。
func (t *Tree) SUCCESSOR(x *Node) *Node {
	if x.right != nil {
		return x.right.MINIMUM()
	}
	y := x.p

	for y != nil && x == y.right {
		x = y
		y = y.p
	}
	return y
}

// TREE-PREDECESSOR(x)
//   if x.left != NIL
//     return x.left
//   y = x.p
//   while y != NIL and y.left = x

// TREE-INSERT(T, z)
//   y = NIL
//   x = T.root
//   while x != NIL
//     y = x
//     if z.key < x.key
//       x = x.left
//     else
//       x = x.right
//   z.p = y
//   if y == NIL
//     T.root = x	// tree T was empty
//   elseif z.key < y.key
//     y.left = z
//   else
//     y.right =  z
func (t *Tree) INSERT(z *Node) {
	var y *Node = nil
	x := t.root
	for x != nil {
		y = x
		if z.key < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}
	z.p = y
	if y == nil {
		t.root = z
	} else if z.key < y.key {
		y.left = z
	} else {
		y.right = z
	}
}
