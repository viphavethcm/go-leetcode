package linked

type ListNode[T any] struct {
	Val  T
	Next *ListNode[T]
}

func NewListNode[T any](vals ...T) *ListNode[T] {
	var head, tail *ListNode[T]
	for _, v := range vals {
		node := &ListNode[T]{Val: v, Next: nil}
		if head == nil {
			head = node
			tail = node
		} else {
			tail.Next = node
			tail = node
		}
	}
	return head
}
