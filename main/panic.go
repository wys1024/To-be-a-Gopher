package main

import "time"

func main() {
	go panicFunc()
	time.Sleep(1 * time.Second)
}

func panicFunc() {
	defer func() {
		if err := recover(); err!= nil {
			println("recover")
		}
	}()

	panic("panic")
}
