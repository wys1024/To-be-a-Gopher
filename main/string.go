package main

//58运维面试题
// 1.最长公共前缀(如下代码)
// 2.怎么防止写一个已关闭的channel（请写出代码
//通过布尔变量和互斥锁 sync.Mutex 控制 channel 状态;
// 3. redis 主从复制
// 4.redis 哨兵模式，集群模式的区别，各自的优缺点，怎么保证数据一致性
//sql:一个表 两个字段性别，成绩。请写一个sql统计一下女生各个成绩有多少人
// SELECT score, COUNT(*) AS count
// FROM students
// WHERE gender = '女'
// GROUP BY score
// ORDER BY score;

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0{
		return ""
	}

	for i := 0; i < len(strs[0]); i++ {
		char := strs[0][i]

		for j := 1; j < len(strs); j++ {
			if strs[j][i]!= char || i == len(strs[j]){
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}



func main() {
	strs := []string{"apple", "application", "april"}
	fmt.Println(longestCommonPrefix(strs)) // 输出: "ap"
}
