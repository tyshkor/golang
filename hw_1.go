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
		"Ivan Ivanov",
		30,
		"coder",
		5}
	fmt.Print(q)
	fmt.Printf("Hello World")

}
