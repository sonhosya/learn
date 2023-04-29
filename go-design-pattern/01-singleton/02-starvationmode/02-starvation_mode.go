package starvationmode

import "fmt"

type Singleton struct{ val int }

var instance *Singleton

// var instance = new(Singleton)
// 提高对比效果使用 init（）
func init() {
	fmt.Printf("instance: %p, not ready.\n", instance)
	instance = new(Singleton)
	fmt.Printf("instance: %p, init by init()\n", instance)
}
func GetInstancByStavationMode() *Singleton {
	return instance
}
