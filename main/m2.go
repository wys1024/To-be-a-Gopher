package main

import "fmt"

func main() {
    var N, X int
    fmt.Scanf("%d %d", &N, &X)

    a := make([]int, N)
    for i := 0; i < N; i++ {
        fmt.Scanf("%d", &a[i])
    }

    m := make(map[int]int)
    for _, v := range a {
        m[v]++
    }

    count := 0
    for _, v := range a {
        target := X - v
        if j, ok := m[target]; ok {
            count += j
        }
        // 如果 (v, v) 本身可以构成一对，则需要减去一个

    }

    // 每对数被计算了两次，除以2得到正确的数量
    fmt.Println(count)
}