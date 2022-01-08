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

// append to the last of the ll
func (ll *linkedList) append(value int) {

	if ll.head.value == 0 { //Write to head if head is empty

		ll.head.value = value

	} else if ll.tail.value == 0 { //Write to tail if tail is empty

		ll.tail.value = value

	} else { //For everything else just add the value to tail and shift tail to the last of head chain

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

// Fetches link at position and updates value for that link
func (ll *linkedList) update(position int, value int) error {

	node, err := ll.getNode(position)

	if err != nil {
		return fmt.Errorf("Error getting value at %v : %v .\n", ll.elem, err)
	}

	*node = link{value: value, next: node.next}

	return err
}

// Remove Node
func (ll *linkedList) remove(position int) error {

	if position > 1 && position < ll.elem {
		prevNode, err := ll.getNode(position - 1)
		if err != nil {
			return err
		}
		currentNode := prevNode.next
		prevNode.next = currentNode.next

	} else if position == 1 {
		ll.head = *ll.head.next
	} else if position == -1 || position == ll.elem {
		prevNode, err := ll.getNode(ll.elem - 1)
		if err != nil {
			return err
		}
		ll.tail = link{value: prevNode.value}
	}

	ll.elem -= 1
	return nil
}

// Prepend to the position
func (ll *linkedList) prepend(position int, value int) error {

	prevNode, err := ll.getNode(position - 1)
	prevNext := *prevNode.next
	if err != nil {
		return fmt.Errorf("Error getting previous node: %v", err)
	}

	*prevNode.next = link{value: value, next: &prevNext}

	ll.elem += 1
	return nil
}

func (ll *linkedList) pop() {
	ll.remove(ll.elem)
}

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

	fmt.Println("ORIGINAL LINKED LIST")
	ll.printLinkedList()

	update_position, update_value := 4, 12
	ll.update(update_position, update_value)
	fmt.Printf("AFTER UPDATING %v index with %v value. \n", update_position, update_value)
	ll.printLinkedList()

	remove := 8
	ll.remove(remove)
	fmt.Printf("AFTER REMOVING %v position.\n", remove)
	ll.printLinkedList()

	prepend_position := 4
	prepend_value := 69
	ll.prepend(prepend_position, prepend_value)
	fmt.Printf("AFTER PREPENDING %v to position %v. \n", prepend_value, prepend_position)
	ll.printLinkedList()

	ll.pop()
	fmt.Println("AFTER POP()")
	ll.printLinkedList()

	fmt.Println("head is", ll.head)
	fmt.Println("tail is", ll.tail)

}
