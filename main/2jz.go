package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param n int整型
 * @return int整型
 */
func bitwiseComplement( n int ) int {
    // 将n转换为二进制字符串
    binaryStr := fmt.Sprintf("%b", n)

    // 翻转二进制字符串中的每一位
    complementStr := ""
    for i := 0; i < len(binaryStr); i++ {
        if binaryStr[i] == '0' {
            complementStr += "1"
        } else {
            complementStr += "0"
        }
    }
    // 将翻转后的二进制字符串转换为十进制整数
    complement, _ := fmt.Sscanf(complementStr, "%b")
    return complement
}