package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// 并发加锁 sync 互斥锁/读写锁；atomic 原子操作；sync.Once 单例
// 双重检查
func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(singletonA())
			// fmt.Println(singletonB())
			// fmt.Println(singletonC())
			// fmt.Println(singletonD())
		}()
	}
	wg.Wait()
}

type Object struct {
	Val int
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var (
	rwm  sync.RWMutex
	sm   sync.Mutex
	flag int32
	sy   sync.Once
	o    *Object
)

func singletonA() *Object {

	rwm.RLock()
	if nil == o {
		rwm.RUnlock()
		rwm.Lock()
		o = &Object{
			Val: rand.Int(),
		}
		rwm.Unlock()
		return o
	}
	rwm.RUnlock()
	return o
}

func singletonA1() *Object {

	// rwm.RLock()
	if nil == o {
		// rwm.RUnlock()
		rwm.Lock()
		if nil == o {
			o = &Object{
				Val: rand.Int(),
			}
		}
		rwm.Unlock()
		return o
	}
	// rwm.RUnlock()
	return o
}

func singletonB() *Object {
	if atomic.CompareAndSwapInt32(&flag, 0, 1) {
		o = &Object{
			Val: rand.Int(),
		}
	}
	return o
}

func singletonC() *Object {
	sy.Do(func() {
		o = &Object{
			Val: rand.Int(),
		}
	})
	return o
}

func singletonD() *Object {
	if o == nil {
		sm.Lock()
		defer sm.Unlock()
		if o == nil {
			o = &Object{
				Val: rand.Int(),
			}
		}
	}
	return o
}
