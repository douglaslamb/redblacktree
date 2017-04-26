package main

import "testing"

func TestIsLeftChild(t *testing.T) {
	// is left child
	parent := &StringNode{}
	child := &StringNode{}
	parent.left = child
	child.parent = parent
	isLeft, ok := child.isLeftChild()
	if !ok || !isLeft {
		t.Errorf("child.isLeftChild() = (%v, %v); want true, true)", isLeft, ok)
	}
	// is right child
	parent = &StringNode{}
	child = &StringNode{}
	parent.right = child
	child.parent = parent
	isLeft, ok = child.isLeftChild()
	if !ok || isLeft {
		t.Errorf("child.isLeftChild() = (%v, %v); want false, true)", isLeft, ok)
	}
	// has no parent
	child = &StringNode{}
	isLeft, ok = child.isLeftChild()
	if ok {
		t.Errorf("child.isLeftChild() = (%v, %v); want false, false)", isLeft, ok)
	}
}

func TestSetLeft(t *testing.T) {
	parent := &StringNode{}
	child := &StringNode{}
	parent.setLeft(child)
	if parent.left != child {
		t.Errorf("parent.setLeft failed; parent.left = %v; want %v", parent.left, child)
	}
	if child.parent != parent {
		t.Errorf("parent.setLeft failed; child.parent = %v; want %v", child.parent, parent)
	}
}

func TestSetRight(t *testing.T) {
	parent := &StringNode{}
	child := &StringNode{}
	parent.setRight(child)
	if parent.right != child {
		t.Errorf("parent.setRight failed; parent.right = %v; want %v", parent.right, child)
	}
	if child.parent != parent {
		t.Errorf("parent.setRight failed; child.parent = %v; want %v", child.parent, parent)
	}
}