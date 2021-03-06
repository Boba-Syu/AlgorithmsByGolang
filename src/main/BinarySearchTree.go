package main

/**
 * 二叉查找树
 */
type TreeNode struct {
	// 二叉树结点结构体
	val   int       // 二叉树结点的值
	key   int       // 二叉树结点的键
	n     int       //以该结点为根的字数中的结点总数
	left  *TreeNode // 左结点
	right *TreeNode // 右结点
}

func size(root *TreeNode) int { // 返回树的高度
	return root.n
}

func newTreeNode(key, value, n int) *TreeNode { // 新建结点
	t := new(TreeNode)
	t.key = key
	t.val = value
	t.n = n
	return t
}

func get(root *TreeNode, key int) int { // 查找键为key的结点的值
	if root == nil {
		return -1
	}
	if key < root.key {
		return get(root.left, key)
	}
	if key > root.key {
		return get(root.right, key)
	}
	return root.val
}

func floor(root *TreeNode, key int) *TreeNode { // 查找键为key的结点
	if root == nil {
		return nil
	}
	if key < root.key {
		return floor(root.left, key)
	}
	if key > root.key {
		return floor(root.right, key)
	}
	return root
}

func find(root *TreeNode, k int) *TreeNode { // 查找排名为k的结点
	if root == nil {
		return nil
	}
	t := root.left.n
	if k < t {
		return find(root.left, k)
	}
	if k > t {
		return find(root.right, k-t-1)
	}
	return root
}

func rank(root *TreeNode, key int) int { // 返回键为key的结点的排名
	if root == nil {
		return 0
	}
	if key < root.key {
		return rank(root.left, key)
	}
	if key > root.key {
		return rank(root.right, key) + root.n + 1
	}
	return root.left.n
}

func put(root *TreeNode, key, value int) *TreeNode { // 更新键为key的结点的值为value
	if root == nil {
		return newTreeNode(key, value, 1)
	}
	if key < root.key {
		root.left = put(root.left, key, value)
	} else if key > root.key {
		root.right = put(root.right, key, value)
	} else {
		root.val = value
	}
	root.n = root.left.n + root.right.n
	return root
}

func min(root *TreeNode) *TreeNode { // 返回键最小的结点
	if root.left == nil {
		return root
	}
	return min(root.left)
}

func max(root *TreeNode) *TreeNode { // 返回键最小的结点
	if root.right == nil {
		return root
	}
	return max(root)
}

func delMin(root *TreeNode) *TreeNode { // 删除键最小的结点
	if root.left == nil {
		return root.right
	}
	root.left = delMin(root.left)
	root.n = root.left.n + root.right.n + 1
	return root
}

func delMax(root *TreeNode) *TreeNode {
	if root.right == nil {
		return root.left
	}
	root.right = delMin(root.right)
	root.n = root.left.n + root.right.n + 1
	return root
}

func delete(root *TreeNode, key int) *TreeNode { // 删除键为key的结点
	if root == nil {
		return nil
	}
	if key < root.key {
		root.left = delete(root.left, key)
	} else if key > root.key {
		root.right = delete(root.right, key)
	} else {
		if root.right == nil {
			return root.left
		}
		if root.left == nil {
			return root.right
		}
		t := min(root.right)
		t.right = delMin(t.right)
		root.left = t.left
	}
	root.n = root.left.n + root.right.n + 1
	return root
}
