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

// PREORDER-TREE-WALK(x)
func (t *Tree) PREORDER_WALK(walker func(*Node)) {
	if t != nil && t.root != nil {
		t.root.PREORDER_WALK(walker)
	}
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

// POSTORDER-TREE-WALK(x)
func (t *Tree) POSTORDER_WALK(walker func(*Node)) {
	if t != nil && t.root != nil {
		t.root.POSTORDER_WALK(walker)
	}
}

// LEVELORDER-TREE-WALK(x)
func (t *Tree) LEVELORDER_WALK(walker func(*Node)) {
	if t != nil && t.root != nil {
		t.root.LEVELORDER_WALK(walker)
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

// TREE-INSERT-RECURSION(T,z)
//   if t.root == nil
//     t.root = z
//   if z.key < t.root.key
//     INSER_LEFT(t.root, z)
//   if z.key > t.root.key
//     INSER_RIGHT(t.root, z)
// - - -
// INSERT_LEFT(p,z)
//   if p.left == nil
//     p.left = z
//   else
//     x = p.left
//     if z.key < x.key
//       INSERT_LEFT(x, z)
//     else if z.key > x.key
//       INSERT_RIGHT(x, z)
// INSERT_RIGHT(p,z)
//   if p.right == nil
//     p.right = z
//   else
//     x = p.right
//     if z.key < x.key
//       INSERT_LEFT(x, z)
//     else if z.key > x.key
//       INSERT_RIGHT(x, z)
func (t *Tree) INSERT_RECURSION(z *Node) {
	if t.root == nil {
		t.root = z
	}
	x := t.root
	if z.key < x.key {
		insert_left(x, z)
	} else if z.key > x.key {
		insert_right(x, z)
	}
}

func insert_left(p, z *Node) {
	if p.left == nil {
		p.left = z
		return
	}

	x := p.left
	if z.key < x.key {
		insert_left(x, z)
	} else if z.key > x.key {
		insert_right(x, z)
	}
}

func insert_right(p, z *Node) {
	if p.right == nil {
		p.right = z
		return
	}

	x := p.right
	if z.key < x.key {
		insert_left(x, z)
	} else if z.key > x.key {
		insert_right(x, z)
	}
}
