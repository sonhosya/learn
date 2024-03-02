package linkedlist

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewListNode(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{"", args{val: 0}, &ListNode{Val: 0}},
		{"", args{val: 1}, &ListNode{Val: 1}},
		{"", args{val: 2}, &ListNode{Val: 2}},
		{"", args{val: 3}, &ListNode{Val: 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewListNode(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewListNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDoublyListNode(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		args args
		want *DoublyListNode
	}{
		// TODO: Add test cases.
		{"", args{val: 0}, &DoublyListNode{Val: 0}},
		{"", args{val: 1}, &DoublyListNode{Val: 1}},
		{"", args{val: 2}, &DoublyListNode{Val: 2}},
		{"", args{val: 3}, &DoublyListNode{Val: 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDoublyListNode(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDoublyListNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertNode(t *testing.T) {
	node0 := NewListNode(0)
	node1 := NewListNode(1)
	node2 := NewListNode(2)
	node3 := NewListNode(3)
	node0.Next = node1
	node1.Next = node2
	node2.Next = node3
	node5 := &ListNode{Val: 5}
	InsertNode(node1, node5)

	for node := node0; node != nil; node = node.Next {
		fmt.Printf("node val:%d\n", node.Val)
	}
	if !reflect.DeepEqual(node1.Next, node5) {
		t.Fatalf("InsertNode does not meet expectations")
	}

}

func TestDropTail(t *testing.T) {
	node0 := NewListNode(0)
	node1 := NewListNode(1)
	node2 := NewListNode(2)
	node3 := NewListNode(3)
	node0.Next = node1
	node1.Next = node2
	node2.Next = node3

	DropTail(node1)
	for node := node0; node != nil; node = node.Next {
		fmt.Printf("node val:%d\n", node.Val)
	}

	if node1.Next != nil {
		t.Fatalf("DropTail does not meet expectations")
	}
}

func TestRemoveNode(t *testing.T) {
	node0 := NewListNode(0)
	node1 := NewListNode(1)
	node2 := NewListNode(2)
	node3 := NewListNode(3)
	node0.Next = node1
	node1.Next = node2
	node2.Next = node3

	RemoveNode(node1)

	for node := node0; node != nil; node = node.Next {
		fmt.Printf("node val:%d\n", node.Val)
	}

	if !reflect.DeepEqual(node1.Next, node3) {
		t.Fatalf("RemoveNode does not meet expectations")
	}
}

func TestAccess(t *testing.T) {
	node0 := NewListNode(0)
	node1 := NewListNode(1)
	node2 := NewListNode(2)
	node3 := NewListNode(3)
	node0.Next = node1
	node1.Next = node2
	node2.Next = node3

	nodeAce := Access(node0, 2)

	for node := node0; node != nil; node = node.Next {
		fmt.Printf("node val:%d\n", node.Val)
	}

	if !reflect.DeepEqual(nodeAce, node2) {
		t.Fatalf("RemoveNode does not meet expectations")
	}
}

func TestFindNode(t *testing.T) {
	node0 := NewListNode(0)
	node1 := NewListNode(1)
	node2 := NewListNode(2)
	node3 := NewListNode(3)
	node4 := NewListNode(2) // 用于测试重复数据的查找顺序是否符合预期
	node0.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4

	for node := node0; node != nil; node = node.Next {
		fmt.Printf("node val:%d\n", node.Val)
	}

	nodeF := FindNode(node0, 2)
	if !reflect.DeepEqual(nodeF, node2) {
		t.Fatalf("FindNode does not meet expectations")
	}

	nodeF = FindNode(node3, 2)
	if reflect.DeepEqual(nodeF, node2) {
		t.Fatalf("FindNode does not meet expectations")
	}
	if !reflect.DeepEqual(nodeF, node4) {
		t.Fatalf("FindNode does not meet expectations")
	}

}
