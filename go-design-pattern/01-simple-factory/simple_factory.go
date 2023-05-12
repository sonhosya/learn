package main

import "fmt"

// 水果类 抽象接口
type Fruit interface {
	Show()
}

// 苹果 具体水果类
type Apple struct {
}

func (a *Apple) Show() {
	fmt.Println("我是苹果")
}

// 香蕉 具体水果类
type Banana struct {
}

func (b *Banana) Show() {
	fmt.Println("我是香蕉")
}

// 梨 具体水果类
type Pear struct {
	Fruit
}

func (pear *Pear) Show() {
	fmt.Println("我是梨")
}

type SimpleFactory struct {
}

func (sf *SimpleFactory) CreateFruit(kind string) Fruit {
	switch kind {
	case "apple":
		return &Apple{}
	case "banana":
		return &Banana{}
	case "pear":
		return &Pear{}
	}
	return nil
}
func main() {
	factory := SimpleFactory{}
	apple := factory.CreateFruit("apple")
	apple.Show()
	banana := factory.CreateFruit("banana")
	banana.Show()
	pear := factory.CreateFruit("pear")
	pear.Show()
}
