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

func (ht *hashTable) get(p person) (person, error) {
	key := hashFunc(p.name)
	per := ht.table[key]

	if per.name == "" {
		return per, fmt.Errorf("Couldn't find the person. \n")
	} else {
		return per, nil
	}

}

func (ht *hashTable) keys() []string {
	var keys []string
	for _, val := range ht.table {
		if val.name != "" {
			keys = append(keys, val.name)
		}
	}
	return keys
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

	p8 := person{name: "some person 8", age: 32}
	fmt.Println(table.get(p3))

	fmt.Println("getting p8")

	p, err := table.get(p8)

	if err != nil {
		fmt.Println("Person Not Found")
	} else {
		fmt.Println(p)
	}

	fmt.Println(table.table)
	fmt.Println(table.keys())
}
