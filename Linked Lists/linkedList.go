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
				*next = &link{value: ll.tail.value}
				ll.tail = link{value: value}
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
func (ll *linkedList) update(position int, value int) error {

	node, err := ll.getNode(position)

	if err != nil {
		return fmt.Errorf("Error getting value at %v : %v .\n", ll.elem, err)
	}

	*node = link{value: value, next: node.next}

	return err
}

// Iterates over the linkedList to get the value at link; starts at 1
func (ll *linkedList) getNode(position int) (*link, error) {
	var link *link

	if position == 0 {

		return link, fmt.Errorf("Link starts at 1")

	} else if position == 1 {

		return &ll.head, nil

	} else if position == ll.elem || position == -1 {
		return &ll.tail, nil

	} else if position > ll.elem {

		return link, fmt.Errorf("Linked List out of range")

	} else {

		link = ll.head.next
		for i := 2; i <= position; i++ {
			if i == position {
				break
			} else {
				link = link.next
			}

		}
	}
	return link, nil
}

func main() {
	ll := linkedList{}
	ll.append(1)
	ll.append(2)
	ll.append(3)
	ll.append(4)
	ll.append(5)
	ll.append(6)
	ll.append(7)
	ll.append(8)

	node := 4
	ll.update(4, 12)

	link, err := ll.getNode(node)

	fmt.Println("Linked List is", ll)
	fmt.Println("ll.head", ll.head)
	fmt.Println("ll.head.next", *ll.head.next)

	link, err = ll.getNode(node)

	if err != nil {
		fmt.Println("Error while fetching link: ", err)
	} else {
		fmt.Printf("link at node %v is %v. \n", node, link)
	}
}
