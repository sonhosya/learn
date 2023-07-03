package main

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);index:idx_name,unique"`
	Age  int    `gorm:"type:tinyint unsigned"`
}

// user 实际表名
func (m *User) TableName() string {
	return "my_user"
}

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:123456@tcp(127.0.0.1:3306)/my_db?charset=utf8mb4&parseTime=True&loc=Local"
	var newLogger = logger.New(
		log.New(os.Stdout, "", log.LstdFlags),
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
	// 迁移 schema  直接生成表
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Println(err)
		return
	}

	// Create
	db.Create(&User{Name: "张三", Age: 18})
	// Read
	var user User
	user.ID = 1
	log.Printf("查询前：%#v", user)

	db.First(&user)
	log.Printf("查询结果：%#v", user)
	db.Delete(&user)
}
