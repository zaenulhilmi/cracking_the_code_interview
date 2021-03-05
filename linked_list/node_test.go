package linked_list

import (
	"testing"
)

func TestAddNode(t *testing.T) {
	node := Node{1, nil}
	node.Add(2)
	if node.Next.Data != 2 {
		t.Logf("Data should be 2 instead of %v", node.Data)
		t.Fail()
	}
}

func TestRemoveNode(t *testing.T) {
	node := Node{1, nil}
	node.Add(2)
	node.Add(3)
	node.Remove(2)

	if node.Next.Data != 3 {
		t.Logf("Data should be 3 but found %v", node.Next.Data)
		t.Fail()
	}

	node2 := Node{1, nil}
	node2.Add(2)
	node2.Remove(1)

	if node2.Data != 2 {
		t.Logf("Data should be 2 but found %v", node2.Data)
		t.Fail()
	}
}

