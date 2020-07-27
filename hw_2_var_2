package main

import "fmt"

type profession(type T) interface {
	type hr, coder, sales, athlete, lawyer
}

type res struct {
	a        int
	hr1      worker(hr)
	coder1   worker(coder)
	sales1   worker(sales)
	athlete1 worker(athlete)
	lawyer1  worker(lawyer)
}

type worker(type T profession) struct {
	pers           person
	specialization T
	exp int

}

type person struct {
	name string
	age  int
}

type hr struct {
	country string
	field   string
}

type coder struct {
	primaryTech   string
	secondaryTech string
}

type sales struct {
	field string
	wins   string
}

type athlete struct {
	sport  string
	wins   string

}

type lawyer struct {
	specialization string
	wins   string
}

func main() {
	/*
	a := worker(hr){
	person{"Maksym Sydorov", 27},
	 hr{"Ukraine", "IT"}, 4,
	}
	fmt.Print(a)
	*/
	var b res
	x := []string{"coder", "Ukraine", "IT"}
	b = workerBuilder(person{"Maksym Sydorov", 28}, 5, x)
	Print(b)
	
	

}

func workerBuilder(pers person, exp int, a []string) (res1 res) {
	switch a[0] {
	case "hr":
		res1.a = 0
		res1.hr1 = worker(hr){pers, hr{a[1], a[2]}, exp}
	case "coder":
		res1.a = 1
		res1.coder1 = worker(coder){pers, coder{a[1], a[2]}, exp}
	case "sales":
		res1.a = 2
		res1.sales1 = worker(sales){pers, sales{a[1], a[2]}, exp}
	case "athlete":
		res1.a = 3
		res1.athlete1 = worker(athlete){pers, athlete{a[1], a[2]}, exp}
	case "lawyer":
		res1.a = 4
		res1.lawyer1 = worker(lawyer){pers, lawyer{a[1], a[2]}, exp}
	}
	return res1
}


func Print(b res) {
	switch b.a {
	case 0:
		fmt.Println(b.hr1)
	case 1:
		fmt.Println(b.coder1)
	case 2:
		fmt.Println(b.sales1)
	case 3:
		fmt.Println(b.athlete1)
	case 4:
		fmt.Println(b.lawyer1)
	}
}
