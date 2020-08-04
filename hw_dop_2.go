package main

import (
	"fmt"
	"sync"
)

var m sync.RWMutex

type calling interface {
	callName() string
	callProfession() string
	callPosition() string
}

type head struct {
	pers *human
	dept string
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

type cache struct {
	m map[int]calling
}

type cacheWorker map[int]worker

type cacheHead map[int]head

func (q worker) toHuman() human {
	res := (*q.pers)
	return res
}

func (c *cacheWorker) call(wg *sync.WaitGroup) {
	fmt.Println("worker_map")
	for _, v := range *c {
		fmt.Println(v)
	}

	wg.Done()
	m.Unlock()
}

func (c *cacheHead) call(wg *sync.WaitGroup) {

	fmt.Println("head_map")

	for _, v := range *c {
		fmt.Println(v)
	}

	wg.Done()
	m.Unlock()
}

func main() {
	q := worker{&human{"John", 24}, 1, "junior"}
	fmt.Println(q.toHuman())

	cWorker := make(cacheWorker)
	cHead := make(cacheHead)

	q1 := worker{&human{"John", 24}, 1, "junior"}
	q2 := worker{&human{"Jack", 26}, 3, "middle"}
	q3 := worker{&human{"Mike", 23}, 2, "junior"}

	cWorker[1] = q1
	cWorker[2] = q2
	cWorker[3] = q3

	p1 := head{&human{"Mary", 35}, "sales"}
	p2 := head{&human{"George", 40}, "development"}
	p3 := head{&human{"Jack", 36}, "hr"}

	cHead[1] = p1
	cHead[2] = p2
	cHead[3] = p3

	var wg sync.WaitGroup

	wg.Add(2)
	m.Lock()
	go cHead.call(&wg)

	m.Lock()
	go cWorker.call(&wg)

	wg.Wait()
}
