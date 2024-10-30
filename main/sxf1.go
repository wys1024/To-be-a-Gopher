package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param st string字符串
 * @return string字符串
 */
func get_substr(st string) string {
    // 定义一个辅助函数，用于消除连续3个及以上的相同字符
    removeSequences := func(s string) (string, bool) {
        var result []rune
        n := len(s)
        if n == 0 {
            return s, false
        }
        var changed bool
        count := 1
        for i := 1; i < n; i++ {
            if rune(s[i]) == rune(s[i-1]) {
                count++
            } else {
                if count >= 3 {
                    changed = true
                    // 不将连续的字符加入结果中
                } else {
                    for j := 0; j < count; j++ {
                        result = append(result, rune(s[i-1]))
                    }
                }
                count = 1
            }
        }
        // 处理最后一组字符
        if count >= 3 {
            changed = true
        } else {
            for j := 0; j < count; j++ {
                result = append(result, rune(s[n-1]))
            }
        }
        return string(result), changed
    }

    // 反复消除，直到无法再消除
    for {
        newStr, changed := removeSequences(st)
        if !changed {
            break
        }
        st = newStr
    }
    return st
}

// 以下是用于测试的主函数，可以根据需要进行修改
func main() {
    examples := []string{
        "aabbaa",          // 原始密文
        "aaabbaa",         // 插入了"aaa"
        "aabaaabbaa",      // 插入了"aaa"和"bbb"
        "aaabbbcccaaa",    // 多次插入
        "aabbaabb",        // 无需消除
        "",                // 空字符串
        "aaabbaaabbb",     // 多个需要消除的部分
    }

    for _, ex := range examples {
        fmt.Printf("混淆后: %s -> 恢复后: %s\n", ex, get_substr(ex))
    }
}
