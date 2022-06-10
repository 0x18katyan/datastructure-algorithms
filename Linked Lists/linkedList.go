package main

import (
	"fmt"
)

type link struct {
	value int
	next  *link
}

type linkedList struct {
	head link
	tail link
	elem int
}

// Iterates over the linkedList to get the value at link; starts at 1
func (ll *linkedList) getNode(position int) (*link, error) {
	var link *link

	if position == 0 { //link indexing starts at 1

		return link, fmt.Errorf("link starts at 1")

	} else if position == 1 { // return head if 1

		return &ll.head, nil

	} else if position == ll.elem || position == -1 { // if position eq length of linkedList or position is -1
		return &ll.tail, nil

	} else if position > ll.elem {

		return link, fmt.Errorf("linked list out of range")

	} else {

		link = ll.head.next
		for i := 2; i <= position; i++ { // iterate over the linkedList
			if i == position {
				break
			} else {
				link = link.next // set next link as link for the upcoming iteration
			}
		}
	}
	return link, nil
}

// append to the last of the ll
func (ll *linkedList) append(value int) {

	if ll.head.value == 0 { //Write to head if head is empty

		ll.head.value = value

	} else if ll.tail.value == 0 { //Write to tail if tail is empty

		ll.tail.value = value

	} else { //For everything else just add the value to tail and shift tail to the last of head chain

		next := &ll.head.next

		for i := 1; i <= ll.elem; i++ {

			if *next == nil { // if the dereferenced value of next is nil that means we have reached the end of the link

				*next = &link{value: ll.tail.value}

				ll.tail = link{value: value}

				break

			} else { // keep on iterating over the linkedList

				nextLink := *next

				next = &nextLink.next

			}

		}

	}

	ll.elem += 1

}

// Fetches link at position and updates value for that link
func (ll *linkedList) update(position int, value int) error {

	node, err := ll.getNode(position)

	if err != nil {
		return fmt.Errorf("error getting value at %v : %v", ll.elem, err)
	}

	*node = link{value: value, next: node.next}

	return err
}

// Remove Node
func (ll *linkedList) remove(position int) error {

	if position > 1 && position < ll.elem { // normal case
		prevNode, err := ll.getNode(position - 1)
		if err != nil {
			return err
		}
		currentNode := prevNode.next
		prevNode.next = currentNode.next

	} else if position == 1 { // if removing the head
		ll.head = *ll.head.next
	} else if position == -1 || position == ll.elem { // if removing the last element
		prevNode, err := ll.getNode(ll.elem - 1)
		if err != nil {
			return err
		}
		ll.tail = link{value: prevNode.value}
	}

	ll.elem -= 1
	return nil
}

// Insert to the position, pushes the current value to next
func (ll *linkedList) insert(position int, value int) error {
	//Get previous node then add new node to previous nodes' next then append node to new node
	prevNode, err := ll.getNode(position - 1)
	prevNext := *prevNode.next
	if err != nil {
		return fmt.Errorf("error getting previous node: %v", err)
	}

	*prevNode.next = link{value: value, next: &prevNext}

	ll.elem += 1
	return nil
}

// Prepend to the top of the list
func (ll *linkedList) prepend(value int) {
	prev_head := ll.head
	node := link{value: value, next: &prev_head}
	ll.head = node
	ll.elem += 1
}

// Pop value
func (ll *linkedList) pop() {
	//Remove last value
	ll.remove(ll.elem)
}

//Print linked list
func (ll *linkedList) printLinkedList() {
	for node := 1; node <= ll.elem; node++ {
		link, err := ll.getNode(node)

		if err != nil {
			fmt.Println("Error while fetching link: ", err)
		} else {
			fmt.Printf("link at node %v is %v. \n", node, link)
		}
	}
}
