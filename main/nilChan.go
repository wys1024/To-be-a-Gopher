package main

// 写未初始化的chan
func main() {
	var c chan int
	c <- 1
}
