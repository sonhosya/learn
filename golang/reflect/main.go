package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
}

func main() {
	// 非指针类型变量 reflect
	u := User{Name: "sonhosya", Age: 18} // 实例化user
	uType := reflect.TypeOf(u)           // 利用反射获取变量u的类型
	fmt.Printf("打印变量'u'的类型: '%s'\n", uType.Name())

	for i := 0; i < uType.NumField(); i++ { // 遍历变量u的类型的字段
		f := uType.Field(i)
		fmt.Printf("打印变量'u'包含字段，字段序号: %d, 字段名: '%s', 本字段类型: '%s'\n", i, f.Name, f.Type)
	}

	// 指针类型变量 reflect
	uu := &u
	uuType := reflect.TypeOf(uu)
	fmt.Printf("打印变量'uu'的类型: '%s'\n", uuType.Kind())
	uuElem := uuType.Elem() // 获取指针变量指向的数据（变量）
	fmt.Println("")
	for i := 0; i < uuElem.NumField(); i++ { // 遍历变量u的类型的字段
		f := uuElem.Field(i)
		fmt.Printf("打印变量'u'包含字段，字段序号: %d, 字段名: '%s', 本字段类型: '%s'\n", i, f.Name, f.Type)
	}
}
