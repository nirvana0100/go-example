package main

import (
	"fmt"
	p "go-example/polymorphic"
)

type Student struct {
	id   int
	name string
}

func (s *Student) Sayhi() {
	fmt.Println(s.id, s.name)
}

type Teacher struct {
	addr  string
	group string
}

func (t *Teacher) Sayhi() {
	fmt.Println(t.addr, t.group)
}
func sayhi(h p.Humaner) {
	h.Sayhi()
}
func main() {
	s := Student{666, "mike"}
	sayhi(&s)
	t := Teacher{"bj", "go"}
	sayhi(&t)
}
