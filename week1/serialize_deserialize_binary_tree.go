package main

import (
	"fmt"
	"strconv"
	"strings"
)

// The code is still in progress!

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// indexing way of constructing a binary tree

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {

	if root == nil {
		return ""
	}
	arr := make(map[int]string)

	this.traverse(root, arr, 0)

	data := ""
	for i := 0; i < len(arr)-1; i++ {
		data += arr[i] + ","
	}
	data += arr[len(arr)-1]

	return data
}

func (this *Codec) traverse(root *TreeNode, arr map[int]string, index int) {
	if root != nil {
		arr[index] = fmt.Sprintf("%d", root.Val)
		// left
		this.traverse(root.Left, arr, 2*index+1)
		// right
		this.traverse(root.Right, arr, 2*index+2)
	} else {
		arr[index] = "null"
	}

	return
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {

	if len(data) == 0 {
		return nil
	}

	arr := strings.Split(data, ",")

	root := new(TreeNode)
	this.populateTree(root, arr, 0)

	return root
}

func (this *Codec) populateTree(root *TreeNode, arr []string, index int) {
	root.Val, _ = strconv.Atoi(arr[index])

	if 2*index+1 < len(arr) && len(arr[2*index+1]) >= 0 && arr[2*index+1] != "null" {
		root.Left = new(TreeNode)
		this.populateTree(root.Left, arr, 2*index+1)
	}

	if 2*index+2 < len(arr) && len(arr[2*index+2]) >= 0 && arr[2*index+2] != "null" {
		root.Right = new(TreeNode)
		this.populateTree(root.Right, arr, 2*index+2)
	}

	return
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
