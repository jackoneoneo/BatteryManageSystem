package tool

import (
	"sync"
	"time"
)

// 栅栏
type CountDownLatch struct {
	Count int //线程的数量
	sync.Mutex
}

func (c *CountDownLatch) CountDown() {
	c.Lock()
	c.Count -= 1
	c.Unlock()
}

func (c *CountDownLatch) Await() {
	for {
		c.Lock()
		value := c.Count
		c.Unlock()
		if value != 0 {
			time.Sleep(time.Duration(20) * time.Millisecond)
		} else { //退出等待
			break
		}
	}

}
