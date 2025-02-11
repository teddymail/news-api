package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"news/internal/router"
	"os"
)

var db *gorm.DB
var err error

func init() {
	// 加载.env文件
	err = godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	// 修改为从环境变量中读取数据库连接信息
	dsn := os.Getenv("DB_DSN")
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
	}
}

func main() {
	r := router.SetupRouter(db)

	// 从环境变量中读取PORT参数
	port := os.Getenv("PORT")
	if port == "" {
		port = "8099" // 默认端口
	}

	// 启动服务
	r.Run(":" + port) // 使用从环境变量中读取的端口
}
