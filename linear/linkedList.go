package linear

import "fmt"

//ListNode represents a linked list data structure
type ListNode struct {
	Val  int
	Next *ListNode
}

//InsertFront adds new node to the beginning of the linked list
//Remeber to call this from head
func (l *ListNode) InsertFront(val int) *ListNode {

	head := l

	newNode := &ListNode{Val: val}

	if head == nil {

		head = newNode
		return head
	}

	newNode.Next = head

	return newNode
}

//DeleteNode deletes node with the given value in the linked list
func (l *ListNode) DeleteNode(val int) {

	current := l

	if current == nil {

		return
	}

	if current.Val == val {

		l = l.Next
		return
	}

	for current.Next != nil {

		if current.Next.Val == val {

			current.Next = current.Next.Next
			break
		}

		current = current.Next
	}
}

//AddBack adds a new node to the end of the linked list
func (l *ListNode) AddBack(val int) {

	current := l

	newNode := &ListNode{Val: val}

	if current == nil {

		current = newNode
		return
	}

	for current.Next != nil {

		current = current.Next
	}

	current.Next = newNode
}

//PrintLinkedList prints the linked list
func (l *ListNode) PrintLinkedList() {

	current := l
	fmt.Println()

	for current != nil {

		fmt.Printf("%d -> ", current.Val)

		current = current.Next
	}

	fmt.Println()
}

//FindMiddleNode finds the middle element of a linked list using 2 pointer approach
func (l *ListNode) FindMiddleNode() *ListNode {

	slow := l
	fast := l

	for slow != nil && fast != nil && fast.Next != nil {

		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
