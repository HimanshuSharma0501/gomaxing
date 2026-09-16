package ll

import (
	"fmt"
)

type LL[T any] struct {
	Val  T
	Next *LL[T]
}

func NewLinkedList[T any](val T) *LL[T] {
	return &LL[T]{
		Val:  val,
		Next: nil,
	}
}

func InsertAtLast[T any](val T, ll *LL[T]) *LL[T] {
	temp := ll
	for temp.Next != nil {
		temp = temp.Next
	}
	temp.Next = NewLinkedList(val)

	return ll
}

func InsertAtFirst[T any](val T, ll *LL[T]) *LL[T] {
	newLL := NewLinkedList(val)
	newLL.Next = ll

	return newLL
}

func DeleteFromFirst[T any](ll *LL[T]) (*LL[T], error) {
	if ll == nil {
		return nil, fmt.Errorf("cannot delete from an empty list")
	}
	head := ll
	if head.Next == nil {
		return head, fmt.Errorf("Can not delete the only element via this func")
	}
	head = head.Next
	return head, nil
}

func DeleteFromLast[T any](ll *LL[T]) (*LL[T], error) {
	if ll == nil {
		return nil, fmt.Errorf("cannot delete from an empty list")
	}
	if ll.Next == nil {
		return nil, fmt.Errorf("cannot delete the only element via this func")
	}
	head := ll
	temp := head
	for temp.Next.Next != nil {
		temp = temp.Next
	}
	temp.Next = nil
	return head, nil
}

func DeleteFromPos[T any](ll *LL[T], pos int) (*LL[T], error) {
	if ll == nil {
		return nil, fmt.Errorf("cannot delete from an empty list")
	}

	if pos < 0 {
		return nil, fmt.Errorf("position cannot be negative")
	}

	size := Size(ll)

	if pos >= size {
		return nil, fmt.Errorf(
			"position %d is out of bounds, list size is %d",
			pos,
			size,
		)
	}

	if pos == 0 {
		return ll.Next, nil
	}

	head := ll
	temp := head

	for i := 0; i < pos-1; i++ {
		temp = temp.Next
	}

	temp.Next = temp.Next.Next

	return head, nil
}

func Size[T any](head *LL[T]) int {
	size := 0
	temp := head
	for temp.Next != nil {
		size++
		temp = temp.Next
	}
	return size
}

func PrintLL[T any](head *LL[T]) {
	temp := head
	for temp.Next != nil {
		fmt.Print(temp.Val, "->")
		temp = temp.Next
	}
	fmt.Println(temp.Val)
}

func ArrayToLL[T any](arr []T) *LL[T] {
	var ll *LL[T]
	for i, v := range arr {
		if i == 0 {
			ll = NewLinkedList(v)
		}
		ll = InsertAtLast(v, ll)
	}
	return ll
}

func ArrayToCycledLL[T any](arr []T, cyclePos int) (*LL[T], error) {
	if len(arr) == 0 {
		return nil, fmt.Errorf("cannot create cycle in an empty list")
	}

	if cyclePos < 0 || cyclePos >= len(arr) {
		return nil, fmt.Errorf("cycle position out of bounds")
	}

	var ll *LL[T]

	for i, v := range arr {
		if i == 0 {
			ll = NewLinkedList(v)
			continue
		}

		ll = InsertAtLast(v, ll)
	}
	cycleNode := ll
	for i := 0; i < cyclePos; i++ {
		cycleNode = cycleNode.Next
	}

	last := ll
	for last.Next != nil {
		last = last.Next
	}

	// Create the cycle.
	last.Next = cycleNode

	return ll, nil
}

func CheckIfCyclicLL[T any](list *LL[T]) (bool, error) {
	if list == nil {
		return false, fmt.Errorf("cannot check an empty list")
	}

	slow := list
	fast := list

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true, nil
		}
	}

	return false, nil
}

func FindCyclicPos[T any](list *LL[T]) (int, error) {
	if list == nil {
		return -1, fmt.Errorf("cannot check an empty list")
	}

	slow := list
	fast := list

	// Detect cycle
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			slow = list
			pos := 0
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
				pos++
			}

			return pos, nil
		}
	}

	return -1, nil
}

func RemoveCycleFromLL[T any](head *LL[T]) (*LL[T], error) {
	pos, err := FindCyclicPos(head)

	if err != nil {
		return nil, err
	}
	if pos == -1 {
		return head, nil
	}
	cycleStart := head
	for i := 0; i < pos; i++ {
		cycleStart = cycleStart.Next
	}
	cycleEnd := cycleStart
	for cycleEnd.Next != cycleStart {
		cycleEnd = cycleEnd.Next
	}
	cycleEnd.Next = nil
	return head, nil
}

func EvenFirstOddNext(head *LL[int]) (*LL[int], error) {
	if head == nil {
		return nil, fmt.Errorf("cannot process an empty list")
	}
	var evenHead, evenTail *LL[int]
	var oddHead, oddTail *LL[int]
	temp := head
	for temp != nil {
		next := temp.Next
		temp.Next = nil
		if temp.Val%2 == 0 {
			if evenHead == nil {
				evenHead = temp
				evenTail = temp
			} else {
				evenTail.Next = temp
				evenTail = temp
			}
		} else {
			if oddHead == nil {
				oddHead = temp
				oddTail = temp
			} else {
				oddTail.Next = temp
				oddTail = temp
			}
		}
		temp = next
	}
	if evenHead == nil {
		return oddHead, nil
	}
	evenTail.Next = oddHead

	return evenHead, nil
}

// func main() {
// 	// implementing ll in go

// 	// ll := NewLinkedList(10)
// 	// ll = InsertAtLast(20, ll)
// 	// ll = InsertAtLast(30, ll)
// 	// ll = InsertAtLast(40, ll)
// 	// ll = InsertAtLast(50, ll)

// 	ll := NewLinkedList("H")
// 	ll = InsertAtLast("I", ll)
// 	ll = InsertAtLast("M", ll)
// 	ll = InsertAtLast("A", ll)
// 	ll = InsertAtLast("N", ll)
// 	ll = InsertAtLast("S", ll)
// 	ll = InsertAtLast("H", ll)
// 	ll = InsertAtLast("U", ll)
// 	ll = InsertAtFirst("H", ll)
// 	ll = InsertAtLast("U", ll)
// 	PrintLL(ll)
// 	ll, err := DeleteFromFirst(ll)
// 	if err != nil {
// 		log.Fatal("couldnt delete due to ", err)
// 	}
// 	PrintLL(ll)
// 	ll, err = DeleteFromLast(ll)
// 	if err != nil {
// 		log.Fatal("couldnt delete due to ", err)
// 	}
// 	PrintLL(ll)
// 	size := Size(ll)
// 	fmt.Println(size)
// 	// ll, err = DeleteFromPos(ll, 90)
// 	ll, err = DeleteFromPos(ll, 4)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	PrintLL(ll)

// }
