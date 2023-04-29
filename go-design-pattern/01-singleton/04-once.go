package singleton

import (
	"fmt"
	"sync"
)

var once sync.Once

func GetInstancByOnce() *Singleton {
	once.Do(func() {
		fmt.Printf("instance: %p, not ready.\n", instance)
		instance = new(Singleton)
		fmt.Printf("instance: %p, init by sync.once.\n", instance)
	})
	return instance
}
