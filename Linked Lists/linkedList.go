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

// append to the last of the ll
func (ll *linkedList) append(value int) {
	if ll.head.value == 0 {
		ll.head.value = value
	} else if ll.tail.value == 0 {
		ll.tail.value = value
	} else {
		next := &ll.head.next
		for i := 1; i <= ll.elem; i++ {
			if *next == nil {
				li := link{value: value, next: nil}
				// tail := ll.tail.value
				// *next = &ll.tail
				*next = &li
				break
			} else {
				nextLink := *next
				next = &nextLink.next
			}
		}
	}
	ll.elem += 1
}

// Will have to go to that position and shift all others to the next?
func (ll *linkedList) update(position int) {

}

// Iterates over the linkedList to get the value at link; starts at 1
func (ll *linkedList) get(position int) (link, error) {
	var link link
	if position == 1 {
		return ll.head, nil
	} else if position == ll.elem {
		return ll.tail, nil
	} else if position > ll.elem {
		return link, fmt.Errorf("Linked List out of range")
	} else {
		link = *ll.head.next
		for i := 2; i < ll.elem; i++ {
			if i == position {
				break
			} else {
				link = *link.next
			}
		}
	}
	return link, nil
}

func main() {
	ll := linkedList{}
	ll.append(100)
	ll.append(99)
	ll.append(198)
	ll.append(199)
	ll.append(200)
	ll.append(201)

	fmt.Println("printing ll at 83", ll)

	link, err := ll.get(5)

	if err != nil {
		panic(err)
	} else {
		fmt.Println(link.value)
	}
}
