package main

import "fmt"

// This calls the "generic" cons implementation for int
// generated by running cpp on pair.go.H
// No type conversions and type safe
func generated() {
	lst := Cons_int(3)
	lst = lst.Cons(Cons_int(4))
	first := lst.Car()
	second := lst.Cdr().Car()
	fmt.Println(first)
	fmt.Println(second)
	fmt.Printf("Add: %d\n", first+second)
}

type Int int

// One can just use the + operator here, but I'm making it
// as comparable to the version using reflection
func (i Int) Add(j Int) Int {
	return i + j
}

func generated_sum() {
	lst := Cons_Int(Int(1))
	lst = lst.Cons(Cons_Int(Int(2)))
	fmt.Println(lst.Sum() + 1)
}
