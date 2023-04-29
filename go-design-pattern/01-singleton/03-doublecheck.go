package singleton

import "fmt"

// 相较于 GetInstancByLazyReload()
// 保证并发安全的前提下，锁的范围减小，提高了程序性能
func GetInstancByDoubleCheck() *Singleton {
	if instance == nil {
		locker.Lock() // 保证初始化值时，并发安全
		if instance == nil {
			fmt.Printf("instance: %p, not ready.\n", instance)
			instance = new(Singleton)
			fmt.Printf("instance: %p, init by double check.\n", instance)
		}
		locker.Unlock()

	}
	return instance
}
