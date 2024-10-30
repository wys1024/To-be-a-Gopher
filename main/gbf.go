package main

// 场景：在一个高并发的web服务器中，要限制IP的频繁访问。现模拟100个IP同时并发访问服务器，每个IP要重复访问1000次。
// 每个IP三分钟之内只能访问一次。修改以下代码完成该过程，要求能成功输出 success:100
import (
	"fmt"
	"sync"
	"time"
)
 
type Ban struct {
	visitIPs map[string]time.Time
	mu   sync.Mutex
}
 
func NewBan() *Ban {

	return &Ban{visitIPs: make(map[string]time.Time)}
}
func (o *Ban) visit(ip string) bool {
	o.mu.Lock()
	defer o.mu.Unlock()

	now := time.Now()
	// 检查是否存在
	if lastVisit, ok := o.visitIPs[ip]; ok {
		if now.Sub(lastVisit) < 3*time.Minute {
			return true
		}

	}
	o.visitIPs[ip] = now
	return false
}
func main() {
	success := 0
	var wg sync.WaitGroup
	var musucess sync.Mutex
	wg.Add(1000*100)

	ban := NewBan()
	for i := 0; i < 1000; i++ {
		for j := 0; j < 100; j++ {
			go func(j int) {
				defer wg.Done()
				ip := fmt.Sprintf("192.168.1.%d", j)
				if !ban.visit(ip) {
					musucess.Lock()
					success++
					musucess.Unlock()
					
				}
			}(j)
		}
 
	}
	wg.Wait()
	fmt.Println("success:", success)
}
