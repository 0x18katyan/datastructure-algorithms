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
	fmt.Println("ll at the top", *ll)
	if ll.head.value == 0 {
		ll.head.value = value
	} else if ll.tail.value == 0 {
		ll.tail.value = value
	} else {
		next := &ll.head.next
		for {
			if *next == nil {
				li := link{value: value}
				fmt.Println("Printing *next", *next)
				*next = &li
				fmt.Println("Priting at 27", *ll)
				fmt.Println("print next", *next)
				break
			} else {
				// next = next.next
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
	fmt.Println(ll)
}
