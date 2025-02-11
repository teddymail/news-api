package router

import (
	"github.com/gin-contrib/cors" // 引入cors包
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"news/internal/handler"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// 添加CORS中间件配置
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // 根据实际情况修改允许的源
	r.Use(cors.New(config))

	// 定义路由
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// 设置主db
	handler.SetDB(db)

	// 创建/news分组
	newsGroup := r.Group("/v1")
	{
		// 新闻模块
		newsGroup.GET("news/:id", handler.GetNewsByID)
		newsGroup.GET("news/list", handler.GetNews)
	}

	return r
}
