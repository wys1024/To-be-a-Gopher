/*
场景：在一个高并发的web服务器中，要限制IP的频繁访问。
以下代码模拟100个IP同时并发访问服务器，每个IP会重复访问1000次。限制每个IP三分钟之内只能访问一次。

修改以下代码完成该过程，要求能成功输出 success:100
*/

package main
 
import (
	"fmt"
	"time"
)
 
type Ban struct {
	visitIPs map[string]time.Time
}
 
func NewBan() *Ban {
	return &Ban{visitIPs: make(map[string]time.Time)}
}

// true 表示三分钟内访问过 限制访问，false表示三分钟内没有访问过
func (o *Ban) visit(ip string) bool {
	if _, ok := o.visitIPs[ip]; ok {
		return true
	}
	o.visitIPs[ip] = time.Now()
	return false
}

func main() {
	success := 0
	ban := NewBan()
	// 每个 i 表示一次访问
	for i := 0; i < 1000; i++ {
        // 每个 j 表示一个单独的IP
		for j := 0; j < 100; j++ {
			go func() {
				ip := fmt.Sprintf("192.168.1.%d", j)
				if !ban.visit(ip) {
					success++
					// 模拟耗时操作
					time.Sleep(time.Millisecond * 200)
				}



			}()
		}
 
	}
	fmt.Println("success:", success)
}