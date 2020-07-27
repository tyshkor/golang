package main

import "fmt"

type worker struct {
	name           string
	age            int
	specialization string
	expirience     int
}

func main() {
	q := worker{
		"Maksym Sydorov",
		27,
		"HR",
		4}
	fmt.Println(q)
	fmt.Println("Hello World")
	a := make([]int, 4)

	a = a[:2]
	fmt.Println(a)

	a = a[:cap(a)]
	fmt.Println(a)

	a = a[2:]
	fmt.Println(a)

	a = a[:cap(a)]
	fmt.Println(a)

}
