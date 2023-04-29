# 单例模式
## 懒加载模式
01-lazy_reload.go

## 饥饿模式
02-starvation_mode.go

## 双重检查模式
03-doublecheck.go

## sync.Once模式
04-once.go

## 测试
```shell
go test .\01-singleton\ .run `TEST_NAME` -v
go test .\01-singleton\02-starvationmode -v
```