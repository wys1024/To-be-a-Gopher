package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// readServices 读取一行输入并将其解析为整数切片
func readServices(scanner *bufio.Scanner) ([]int, bool) {
	if !scanner.Scan() {
		return nil, false
	}
	line := strings.TrimSpace(scanner.Text())
	if line == "" {
		return []int{}, true
	}
	parts := strings.Fields(line)
	services := make([]int, 0, len(parts))
	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil || num <= 0 || num >= 1024 { // 确保为正整数且小于1024
			return nil, false
		}
		services = append(services, num)
	}
	return services, true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// 读取服务器A的服务
	A, okA := readServices(scanner)
	if !okA {
		fmt.Println("fail")
		return
	}

	// 读取服务器B的服务
	B, okB := readServices(scanner)
	if !okB {
		fmt.Println("fail")
		return
	}

	// 检查是否有额外的非空输入行
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			fmt.Println("fail")
			return
		}
	}

	// 计算服务器A和B的内存总和
	sumA, sumB := 0, 0
	for _, val := range A {
		sumA += val
	}
	for _, val := range B {
		sumB += val
	}

	// 如果总和已经相等，不需要交换
	if sumA == sumB {
		fmt.Println("fail")
		return
	}

	diff := sumA - sumB
	halfDiff := diff / 2
	halfDiff2 := diff / 2

	// 如果差值不是偶数，无法通过交换一次服务使总和相等
	if diff > 0 {
		if diff%2 != 0 {
			fmt.Println("fail")
			return
		}  
		// 构建服务器B的服务内存到出现次数的映射
		bMap := make(map[int]int)
		for _, val := range B {
			bMap[val]++
		}
	
		// 遍历服务器A的每个服务，查找是否存在符合条件的服务在服务器B
		for _, a := range A {
			b := a - halfDiff
			if b > 0 {
				if count, exists := bMap[b]; exists && count > 0 {
					fmt.Printf("%d %d\n", a, b)
					return
				}
			}
		}

	}else{
		if diff%2!= 0 {
			fmt.Println("fail")
			return
		}
		// 构建服务器A的服务内存到出现次数的映射
		aMap := make(map[int]int)
		for _, val := range A {
			aMap[val]++
		}

		// 遍历服务器B的每个服务，查找是否存在符合条件的服务在服务器A
		for _, b := range B {
			a := b - halfDiff2
			if a > 0 {
				if count, exists := aMap[a]; exists && count > 0 {
					fmt.Printf("%d %d\n", a, b)
					return
				}
			}
		}

	}

	// 如果没有找到符合条件的交换组合
	fmt.Println("fail")
}
