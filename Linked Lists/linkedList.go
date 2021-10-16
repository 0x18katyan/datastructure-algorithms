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
	fmt.Println("ll at the top", *ll)

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

// Will have to go to that position and shift all others to the next?
func (ll *linkedList) update(position int) {

}

func main() {
	ll := linkedList{}
	ll.append(100)
	ll.append(99)
	ll.append(198)
}
