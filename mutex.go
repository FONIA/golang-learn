package main

import (
	"fmt"
	"sync"
)

type SafeDict struct {
	data  map[string]int
	mutex *sync.Mutex
}

// 读写锁,多读少写
// *sync.RWMutex
// .RLock()
// .RUnLock

// type SafeDict2 struct {
// 	data        map[string]int
// 	*sync.Mutex //匿名锁
// }

//初始化：&SafeDict{
//		 data,
//		 &sync.Mutex{}, //初始化
//	}

func NewSafeDict(data map[string]int) *SafeDict {
	return &SafeDict{
		data:  data,
		mutex: &sync.Mutex{}, //初始化
	}
}

func (d *SafeDict) Len() int {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	return len(d.data)
}

func (d *SafeDict) Put(key string, val int) (int, bool) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	old_val, ok := d.data[key]
	d.data[key] = val
	return old_val, ok
}

func (d *SafeDict) Test() {
	d.mutex.Lock()
	len := len(d.data)
	d.mutex.Unlock() //手动释放
	fmt.Println(len)
}
