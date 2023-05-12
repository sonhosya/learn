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

// 新增一个"日本苹果"
// 突出工厂方法 符合开闭原则
type JapanApple struct {
}

func (jp *JapanApple) Show() {
	fmt.Println("我是日本苹果")
}

// 梨 具体水果类
type Pear struct {
	Fruit
}

func (pear *Pear) Show() {
	fmt.Println("我是梨")
}

// 工厂类(抽象接口)
type AbstractFactory interface {
	CreateFruit() Fruit //生产水果类(抽象)的生产器方法
}

// 苹果工厂
type AppleFactory struct{}

func (af *AppleFactory) CreateFruit() Fruit {
	return &Apple{}
}

// 新增一个"日本苹果工厂"
// 突出工厂方法 符合开闭原则
type JapanAppleFactory struct{}

func (jaf *JapanAppleFactory) CreateFruit() Fruit {
	return &JapanApple{}
}

// 香蕉工厂
type BananaFactory struct{}

func (bf *BananaFactory) CreateFruit() Fruit {
	return &Banana{}
}

// 梨工厂
type PearFactory struct{}

func (pf *PearFactory) CreateFruit() Fruit {
	return &Pear{}
}
func main() {
	var appleFac AbstractFactory // 先定义后使用 突出依赖倒转原则
	appleFac = new(AppleFactory)
	var apple Fruit // 先定义后使用 突出依赖倒转原则
	apple = appleFac.CreateFruit()
	apple.Show()

	var bananaFac AbstractFactory
	bananaFac = new(BananaFactory)
	var banana Fruit
	banana = bananaFac.CreateFruit()
	banana.Show()

	var pearFac AbstractFactory
	pearFac = new(PearFactory)
	var pear Fruit
	pear = pearFac.CreateFruit()
	pear.Show()

	var japanAppleFac AbstractFactory
	japanAppleFac = new(JapanAppleFactory)
	var japanApple Fruit
	japanApple = japanAppleFac.CreateFruit()
	japanApple.Show()

}
