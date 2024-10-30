package main

type animal interface {
	move()
}

type Dog struct {
}

func (d Dog) move() {
	println("move ...")
}

func main() {
	a := Dog{}
	a.move()
}