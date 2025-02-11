package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

// GetPageSize 计算分页大小
func GetPageSize(c *gin.Context) (int, int) {
	pageParam, _ := c.GetQuery("page")
	page, _ := strconv.Atoi(pageParam)
	pageSizeParam, _ := c.GetQuery("pageSize")
	pageSize, _ := strconv.Atoi(pageSizeParam)
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10 // 默认每页大小为10
	}
	offset := (page - 1) * pageSize
	return offset, pageSize
}
