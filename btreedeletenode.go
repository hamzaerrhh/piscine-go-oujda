// Empty file
package piscine


type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                 string
}


func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
if root.Left==nil || node==nil{
  return root
}
//if no child
if node.Left==nil && node.Right==nil{
 if node.Parent==nil{
  return nil
 }
 if node.Parent.Left==node {
   node.Parent.Left=nil
 } else{
 node.Parent.Right=nil }

  return root

} 
  

if node.Left==nil || node.Right==nil{
  child:=node.Left
  if child==nil{
    child=node.Right
  }
  if node.Parent.Left==node{
    node.Parent.Right=child
  }else{
    node.Parent.Right=child
  }
  child.Parent=node.Parent
  
return root
  
}

newNode:=BTreeMin(node.Right)
  node.Data=newNode.Data
  root=BTreeDeleteNode(root,newNode)
  return root

  



}

func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	current := root
	for current.Left != nil {
		current = current.Left
	}

	return current
}



func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	BTreeApplyInorder(root.Left, f)
	f(root.Data)
	BTreeApplyInorder(root.Right, f)
}
func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}

	if data < root.Data {
		root.Left = BTreeInsertData(root.Left, data)
		root.Left.Parent = root
	} else {
		root.Right = BTreeInsertData(root.Right, data)
		root.Right.Parent = root
	}

	return root
}

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	}

	if elem > root.Data {
		return BTreeSearchItem(root.Right, elem)
	} else if elem < root.Data {
		return BTreeSearchItem(root.Left, elem)
	} else if elem == root.Data {
		return root
	} else {
		return root
	}
}
func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Left != nil && root.Left.Data >= root.Data {
		return false
	}

	if root.Right != nil && root.Right.Data <= root.Data {
		return false
	}

	if !BTreeIsBinary(root.Left) || !BTreeIsBinary(root.Right) {
		return false
	}

	return true
}


