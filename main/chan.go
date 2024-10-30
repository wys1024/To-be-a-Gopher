package main

func main() {
	c := make(chan string)
	count := 0
	go func() {
		//time.Sleep(1 * time.Second)
		c <- "hello"
		close(c)
	}()

	for{
		select{
		case k, ok := <-c:
			if ok{
				println(k)
			}else{
				c = nil
			}
		default:
			count++
			if count > 20{
				return
			}
			println("default")
		}	
	}
}
