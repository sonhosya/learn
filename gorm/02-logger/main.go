package main

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:123456@tcp(127.0.0.1:3306)/my_db?charset=utf8mb4&parseTime=True&loc=Local"
	var newLogger = logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info, // 设置日志打印等级
			Colorful:      true,        // 设置是否开启彩色打印
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("db connection success.")
	db.Debug()
	// 迁移 schema  直接生成表
	db.AutoMigrate(&Product{})
	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	product.ID = 1
	log.Printf("查询前：%#v", product)

	db.First(&product)
	log.Printf("查询结果：%#v", product)
	db.Delete(&product)
}
