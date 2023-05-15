package main

import "fmt"

type Goods struct {
	Kind string // 商品种类
	Fact bool   //商品真伪
}

// 抽象购物主题
type Shopping interface {
	Buy(goods *Goods) // 任务执行函数
}

// 实现
// 具体的购物主题， 实现了shopping， 去韩国购物
type KoreaShopping struct{}

func (ks *KoreaShopping) Buy(goods *Goods) {
	fmt.Println("去韩国进行了购物, 买了 ", goods.Kind)
}

//具体的购物主题， 实现了shopping， 去非洲购物
type AfrikaShopping struct{}

func (as *AfrikaShopping) Buy(goods *Goods) {
	fmt.Println("去非洲进行了购物, 买了 ", goods.Kind)
}

// 海外代理 代购
type OverseasProxy struct {
	shopping Shopping // 代理某个主题，这里是抽象类型
}

//创建一个代理,并且配置关联被代理的主题
func NewProxy(shopping Shopping) Shopping {
	return &OverseasProxy{shopping}
}

func (op *OverseasProxy) Buy(goods *Goods) {
	// 1. 先验货
	if op.distinguish(goods) == true {
		//2. 进行购买
		op.shopping.Buy(goods) //调用原被代理的具体主题任务
		//3 海关安检
		op.check(goods)
	}
}

//验货流程
func (op *OverseasProxy) distinguish(goods *Goods) bool {
	fmt.Println("对[", goods.Kind, "]进行了辨别真伪.")
	if goods.Fact == false {
		fmt.Println("发现假货", goods.Kind, ", 不应该购买。")
	}
	return goods.Fact
}

//安检流程
func (op *OverseasProxy) check(goods *Goods) {
	fmt.Println("对[", goods.Kind, "] 进行了海关检查， 成功的带回祖国")
}
func main() {
	g1 := Goods{
		Kind: "韩国面膜",
		Fact: true,
	}

	g2 := Goods{
		Kind: "CET4证书",
		Fact: false,
	}

	//如果不使用代理来完成从韩国购买任务
	var shopping Shopping
	shopping = new(KoreaShopping) //具体的购买主题

	//1-先验货
	if g1.Fact == true {
		fmt.Println("对[", g1.Kind, "]进行了辨别真伪.")
		//2-去韩国购买
		shopping.Buy(&g1)
		//3-海关安检
		fmt.Println("对[", g1.Kind, "] 进行了海关检查， 成功的带回祖国")
	}

	fmt.Println("---------------以下是 使用 代理模式-------")
	var overseasProxy Shopping
	overseasProxy = NewProxy(shopping)
	overseasProxy.Buy(&g1)
	overseasProxy.Buy(&g2)
}
