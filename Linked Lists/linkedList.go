package main

import "fmt"

type link struct {
	value int
	next  *link
}

type linkedList struct {
	head link
	tail link
	elem int
}

func newLink(value int) link {
	return link{value: value, next: nil}
}

// append to the last of the ll
func (ll *linkedList) append(value int) {
	if ll.head.value == 0 {

		ll.head = newLink(value)

	} else if ll.tail.value == 0 {

		ll.tail = newLink(value)

	} else {

		node := &ll.head

		for {

			if node.next == nil {

				li := newLink(value)
				node.next = &li
				break

			} else {
				node = node.next
			}

		}

	}

	ll.elem += 1

}

func (ll *linkedList) getNode(position int) (*link, error) {
	node := &ll.head

	if position == -1 {
		return &ll.tail, nil

	} else if ll.elem < position {

		return nil, fmt.Errorf("Element not found at position %v", position)

	} else if position == 2 && ll.elem == 2 {

		return &ll.tail, nil

	} else {
		for i := 0; i <= position; i++ {
			if i != position {
				node = node.next
			} else {
				break
			}
		}
		return node, nil
	}
}

// Will have to go to that position and shift all others to the next?
func (ll *linkedList) update(position int, value int) error {

	newL := newLink(value)

	// if position is greater than num elements
	if position >= ll.elem {

		tail := ll.tail
		ll.tail = newL
		ll.update(ll.elem-1, tail.value)
		return

		// If no elements and update is called then append
	} else if ll.elem == 0 {
		ll.append(value)
		return
	} else {

		node, err := ll.getNode(position)

		if err != nil {
			// handle new node
		}
	}
}

func main() {
	ll := linkedList{}
	ll.append(100)
	ll.append(99)
	ll.append(198)
	fmt.Println(ll)
	node, err := ll.getNode(10)

	if err != nil {
		panic(err)
	}

	fmt.Println(*node)
}
