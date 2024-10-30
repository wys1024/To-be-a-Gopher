package main

import "fmt"


func sjs(n int, mapk map[int]string) string{
	var s string
	for _, v := range mapk {
		s = s + v
		n--
		if n == 0 {
			return s
		}
	}
	return s
}

func main() {
	//mapk := make(map[int]string)
	mapk := map[int]string{
		0: "0",
		1: "1",
		2: "2",
		3: "3",
		4: "4",
		5: "5",
		6: "6",
		7: "7",
		8: "8",
		9: "9",
	}
	var n int
	fmt.Scan(&n)
	fmt.Println(sjs(n, mapk))
}
