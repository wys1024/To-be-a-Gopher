// 2.怎么防止写一个已关闭的channel（请写出代码

package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		c <- 1
		close(c)
	}()

	// 防止写一个已关闭的channel
	_, ok := <-c
	if!ok {
		fmt.Println("channel已关闭")
	}
}
