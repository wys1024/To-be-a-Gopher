package main

import "fmt"

func main() {
    result := testDefer()
    fmt.Println("Result:", result)
}

func testDefer() int {
    defer fmt.Println("Defer 1")
    defer fmt.Println("Defer 2")

    fmt.Println("Before return")
    return 42  // 返回值在此时已经被计算但未实际返回
}
//return计算 -> defer -> 返回计算值