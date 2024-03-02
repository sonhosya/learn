package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil}
}

// 双向链表节点结构体
type DoublyListNode struct {
	Val  int             // 节点值
	Next *DoublyListNode // 指向后继节点的指针
	Prev *DoublyListNode // 指向前驱节点的指针
}

// NewDoublyListNode 初始化
func NewDoublyListNode(val int) *DoublyListNode {
	return &DoublyListNode{
		Val:  val,
		Next: nil,
		Prev: nil,
	}
}

// 将node节点插入current节点的后面
func InsertNode(current *ListNode, node *ListNode) {
	cNext := current.Next
	node.Next = cNext
	current.Next = node
}

// 移除current节点之后的节点链表
func DropTail(current *ListNode) {
	current.Next = nil
}

// 移除current节点之后的第一个节点
func RemoveNode(current *ListNode) {
	if current.Next == nil {
		return
	}
	cNext := current.Next
	current.Next = cNext.Next
}

// 访问链表head索引为index的节点
func Access(head *ListNode, index int) *ListNode {
	for i := 0; i < index; i++ {
		if head == nil {
			return nil
		}
		head = head.Next
	}
	return head
}

// 查找链表中值为target的首个节点
func FindNode(head *ListNode, target int) *ListNode {
	for head != nil {
		if head.Val == target {
			return head
		}
		head = head.Next
	}
	return nil
}
