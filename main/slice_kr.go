package main

import "fmt"

func main() {
    s := make([]int, 0, 2) // 创建一个初始容量为 2 的 slice
    fmt.Println("Initial slice:", s, "len:", len(s), "cap:", cap(s))

    // 添加元素
    for i := 0; i < 5; i++ {
        s = append(s, i)
        fmt.Println("After appending", i, ":", s, "len:", len(s), "cap:", cap(s))
    }
}
