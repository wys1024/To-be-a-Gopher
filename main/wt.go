package main

import (
	"bufio"
	"fmt"
	"os"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 读取初始位置和目标位置
	var a, b, c, d int
	fmt.Fscan(reader, &a, &b, &c, &d)

	// 读取瓶子数量
	var n int
	fmt.Fscan(reader, &n)

	bottles := make([][2]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &bottles[i][0], &bottles[i][1])
	}

	// 选择第一个瓶子
	minCost := int(1e18)
	for i := 0; i < n; i++ {
		xi, yi := bottles[i][0], bottles[i][1]
		cost := abs(xi-a) + abs(yi-b) + abs(xi-c) + abs(yi-d)
		if cost < minCost {
			minCost = cost
		}
	}

	// 从目标位置搬运剩余瓶子
	totalCost := minCost
	for i := 0; i < n; i++ {
		xi, yi := bottles[i][0], bottles[i][1]
		totalCost += abs(xi-c) + abs(yi-d)
	}

	// 输出总代价
	fmt.Fprintln(writer, totalCost)
}