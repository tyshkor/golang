package main

import (
	"fmt"
	"reflect"
)

type calling interface {
	callName() string
	callProfession() string
	callPosition() string
}

type worker struct {
	pers     *human
	exp      int
	position string
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

func (q hr) callName() string {
	a := (*(*(q.wor)).pers).name
	return a
}

func (q hr) callProfession() string {
	a := "hr"
	return a
}

func (q hr) callPosition() string {
	a := (*(q.wor)).position
	return a
}

type coder struct {
	wor           worker
	primaryTech   string
	secondaryTech string
}

func (q coder) callName() string {
	a := (*(q.wor).pers).name
	return a
}

func (q coder) callProfession() string {
	a := "coder"
	return a
}

func (q coder) callPosition() string {
	a := (q.wor).position
	return a
}

type cache struct {
	m map[int]calling
}


func getTypeName(t interface{}) string {
	t1 := reflect.TypeOf(t).Elem()
	return t1.Name()
}

func getTypeCache(c cache) []string {
	res := make([]string, 0)
	for _, v := range c.m {
		res = append(res, getTypeName(v))
	}
	return res
}

func main() {
	c := cache{m: make(map[int]calling)}

	q1 := &hr{&worker{&human{"John", 24}, 1, "junior"}, "Ukraine", "IT"}

	q2 := &hr{&worker{&human{"Jack", 26}, 3, "middle"}, "Ukraine", "Auto"}
	q3 := &coder{worker{&human{"Mike", 23}, 2, "middle"}, "golang", "javascript"}
	c.m[1] = q1
	c.m[2] = q2
	c.m[3] = q3
	//fmt.Print(q.callPosition())
	for _, v := range c.m {
		fmt.Println(v.callName())
		fmt.Println(v.callPosition())
		fmt.Println(v.callProfession())
	}
	fmt.Println("\n\n")
	fmt.Println(getTypeName(q1))
	fmt.Println("\n\n")
	fmt.Println(getTypeCache(c))

}
