package main

import "fmt"

func main() {
    // 创建一个 map
    m := make(map[string]int)
	

    // 插入元素
    m["apple"] = 2
    m["banana"] = 3

    // 读取元素
    fmt.Println("Apple count:", m["apple"])

    // 删除元素
    delete(m, "banana")

    // 检查元素是否存在
    if val, exists := m["banana"]; exists {
        fmt.Println("Banana count:", val)
    } else {
        fmt.Println("Banana not found")
    }
}
