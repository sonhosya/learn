package main

import (
	"fmt"
	"sync"
)

type Singelton struct{}

func (s *Singelton) SomeThing() {
	fmt.Println("单例对象的某方法")
}

var once sync.Once
var singelton *Singelton // 声明变量

func GetInstance() *Singelton {
	// 懒汉模式--使用时初始化
	once.Do(func() {
		// sync.Once 保障只初始化一次 该方式是线程安全的
		singelton = new(Singelton) // 变量赋值
	})
	return singelton
}
func main() {
	s := GetInstance()
	s.SomeThing()
}
