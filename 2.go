package main

import "fmt"

type worker struct {
	pers person
	exp  int "expirience"
}

type person struct {
	name string "fullname"
	age  int
}

type hr struct {
	wor     worker
	country string
	field   string "typeOfCompany"
}

type coder struct {
	wor           worker
	primaryTech   string
	secondaryTech string
}

type sales struct {
	wor   worker
	field string "areaOfinterest"
}

type athlete struct {
	wor    worker
	sport  string
	wins   int
	losses int
}

type lawyer struct {
	wor            worker
	specialization string
}

func main() {

	q := hr{worker{person{"Maksym Sydorov", 27}, 5}, "Ukraine", "IT"}
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
