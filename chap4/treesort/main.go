package main

import "fmt"

/**
 * @Description: 使用二叉树实现插入排序
 * @author lichang
 * @date 20-1-5
 * @time 下午2:58
*/
type tree struct{
	// 递归定义一个二叉树
	value	int
	left, right *tree
}

func Sort(values []int){
	/**
	   @Description: 将给出数组中的值插入到二叉树中
	   @Param:
	        values：int类型数组，待插入到二叉树
	   @Return:

	 */
	var root *tree
	for _, v:= range values{  //先构造排序二叉树
		root = add(root, v)
	}
	appendValue(values[:0], root)	// 从排序二叉树中取出数值
}

func appendValue(values []int, t *tree) []int {
	/**
	   @Description: 中序遍历排序二叉树得到排序后的序列
	   @Param:
	        values[] ：实际上是将原来的values的地址传过来，因此这里对values的修改，实际上是在原数组上的修改
	   @Return：
	        values：插入操作之后的数组
	 */
	if t != nil {
		values = appendValue(values, t.left)	// 先左子树
		values = append(values, t.value)
		values = appendValue(values, t.right)	// 再右子树
	}
	return values
}

func add (t *tree, value int)*tree{
	/**
	   @Description: 插入新的节点，按二叉排序树的规则，左小右大
	   @Param:
	        t: 树的当前结点
			value: 待插入值
	   @Return:
			tree: 被插入的结点指针

	 */
	if t==nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	}else{
		t.right = add(t.right, value)
	}
	return t
}

func main(){
	var values = [4]int{10,8,6,9}
	Sort(values[:])
	fmt.Println(values)
}