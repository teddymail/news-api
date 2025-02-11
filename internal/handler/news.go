package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"news/internal/models"
	"news/internal/utils"
	"strconv"
)

var db *gorm.DB

func SetDB(database *gorm.DB) {
	db = database
}

// GetNewsByID 根据ID获取新闻
func GetNewsByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	news, err := models.GetNewsByID(db, uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "News not found"})
		return
	}

	c.JSON(http.StatusOK, news)
}

// GetNews 根据source和category获取新闻
func GetNews(c *gin.Context) {
	source := c.Query("source")
	category := c.Query("category")
	offset, pageSize := utils.GetPageSize(c) // 获取pageSize参数

	newsList, err := models.GetNewsBySourceAndCategory(db, source, category, offset, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get news"})
		return
	}

	// 增加对新闻总数的查询
	totalCount, err := models.GetNewsCountBySourceAndCategory(db, source, category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get news count"})
		return
	}

	// 返回新闻列表和总数
	c.JSON(http.StatusOK, gin.H{
		"news":       newsList,
		"totalCount": totalCount,
	})
}
