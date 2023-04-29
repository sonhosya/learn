package singleton

import (
	"fmt"
	"sync"
)

type Singleton struct{ val int }

var instance *Singleton
var locker = sync.Mutex{}

func GetInstancByLazyReload() *Singleton {
	// 保证初始化值时 并发安全
	locker.Lock()
	defer locker.Unlock()
	if instance == nil {
		fmt.Printf("instance: %p, not ready.\n", instance)
		instance = new(Singleton)
		fmt.Printf("instance: %p, init by Lazy reload.\n", instance)
	}
	return instance
}
