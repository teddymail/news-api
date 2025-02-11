package migrations

import (
	"gorm.io/gorm"
	"news/internal/models"
)

func Migrate(db *gorm.DB) error {
	// 创建News表
	if err := db.AutoMigrate(&models.News{}); err != nil {
		return err
	}
	return nil
}