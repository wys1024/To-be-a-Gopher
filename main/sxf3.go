package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// backtrack 函数用于递归生成所有可能的出售顺序
func backtrack(input string, pushIndex int, stack []rune, output []rune, results *[]string) {
	// 如果所有字母都已处理且栈为空，则将当前输出序列添加到结果中
	if pushIndex == len(input) && len(stack) == 0 {
		*results = append(*results, string(output))
		return
	}

	// 如果还有字母未压入栈中，可以选择将下一个字母压入栈
	if pushIndex < len(input) {
		// 创建新的栈副本，并将当前字母压入栈中
		newStack := append(stack, rune(input[pushIndex]))
		// 递归调用，处理下一个字母
		backtrack(input, pushIndex+1, newStack, output, results)
	}

	// 如果栈中有字母，可以选择弹出栈顶字母进行出售
	if len(stack) > 0 {
		// 获取栈顶字母
		top := stack[len(stack)-1]
		// 创建新的栈副本，移除栈顶字母
		newStack := stack[:len(stack)-1]
		// 创建新的输出序列，添加被弹出的字母
		newOutput := append(output, top)
		// 递归调用，继续处理
		backtrack(input, pushIndex, newStack, newOutput, results)
	}
}

func main() {
	// 使用 bufio 读取用户输入
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入字母序列（例如 abc）：")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("读取输入时出错:", err)
		return
	}
	// 去除输入中的换行符和空格
	input = strings.TrimSpace(input)

	// 检查输入是否为空
	if len(input) == 0 {
		fmt.Println("输入不能为空。")
		return
	}

	// 初始化结果切片
	var results []string

	// 调用 backtrack 函数生成所有可能的出售顺序
	backtrack(input, 0, []rune{}, []rune{}, &results)

	// 输出所有可能的出售顺序
	fmt.Println("所有可能的出售顺序如下：")
	for _, seq := range results {
		fmt.Println(seq)
	}
}
