package main

import "time"

//控制协程并发数，同时最多10个。 协程创建后1秒自动退出。
func main() {
	ch := make(chan struct{}, 10)

	for i := 0; i < 10; i++ {
		ch <- struct{}{}
		go func() {
			time.Sleep(time.Second)
			<-ch
		}()

		
	}



}
