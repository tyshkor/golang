package main

import "fmt"

type worker struct {
	pers *human
	exp  int
}

type human struct {
	name string
	age  int
}

type hr struct {
	wor     *worker
	country string
	field   string
}

type coder struct {
	wor           worker
	primaryTech   string
	secondaryTech string
}

func (receiver *human) HappyBirthday() {
	fmt.Printf("Happy Birthday yo you %v!\n", (*receiver).name)
}

func (receiver *human) promoted() string {

	res := "You are promoted!"
	return res
}

func (receiver *human) yourAge5YearsAgo() int {

	res := (*receiver).age - 5
	return res
}

func main() {

	q := human{"James", 29}

	q.HappyBirthday()

	fmt.Println(q.promoted())

	fmt.Printf("Your age 5 yers ago was %d", q.yourAge5YearsAgo())
  
	//panic("some problem")
	var r *human
	fmt.Printf(r.age)
}
