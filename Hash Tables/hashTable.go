package main

import "fmt"

type hashTable struct {
	table [100]person
}

type person struct {
	name string
	age  int
}

func hashFunc(val string) int {
	hash := 0
	for i, char := range val {
		hash = (hash + int(char)*i) % len(val)
	}
	return hash
}

func (ht *hashTable) set(p person) {
	key := hashFunc(p.name)
	ht.table[key] = p
}

func (ht *hashTable) get(p person) person {
	key := hashFunc(p.name)
	return ht.table[key]
}

func main() {
	var table hashTable
	var p1 = person{name: "some person 1", age: 25}
	var p2 = person{name: "some person 2", age: 26}
	var p3 = person{name: "some person 3", age: 27}
	var p4 = person{name: "some person 4", age: 28}
	var p5 = person{name: "some person 5", age: 29}
	var p6 = person{name: "some person 6", age: 30}
	var p7 = person{name: "some person 7", age: 31}

	table.set(p1)
	table.set(p2)
	table.set(p3)
	table.set(p4)
	table.set(p5)
	table.set(p6)
	table.set(p7)

	fmt.Println(table.get(p3))
	fmt.Println(table.table)
}
