package heap

// 父节点
func PARENT(i int) int {
	return (i - 1) / 2
}

// 左孩子
func LEFT(i int) int {
	return (i * 2) + 1
}

// 右孩子
func RIGHT(i int) int {
	return (i * 2) + 2
}
