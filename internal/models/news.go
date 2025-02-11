package models

import (
	"gorm.io/gorm"
	"time"
)

// News 定义新闻模型
type News struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string    `gorm:"not null" json:"title"`
	Category    string    `gorm:"not null" json:"category"`
	Content     string    `gorm:"type:text" json:"content"`
	PublishTime time.Time `gorm:"not null" json:"publish_time"`
	Source      string    `gorm:"not null" json:"source"`
	Link        string    `gorm:"not null" json:"link"`
	FeatureCode string    `gorm:"not null;unique" json:"feature_code"`
}

// CreateNews 创建新闻
func CreateNews(db *gorm.DB, news *News) error {
	return db.Create(news).Error
}

// GetNews 获取所有新闻
func GetNews(db *gorm.DB) ([]News, error) {
	var newsList []News
	if err := db.Find(&newsList).Error; err != nil {
		return nil, err
	}
	return newsList, nil
}

// GetNewsByID 根据ID获取新闻
func GetNewsByID(db *gorm.DB, id uint) (*News, error) {
	var news News
	if err := db.First(&news, id).Error; err != nil {
		return nil, err
	}
	return &news, nil
}

// GetNewsBySourceAndCategory 根据source和category获取新闻
func GetNewsBySourceAndCategory(db *gorm.DB, source, category string, offset, pageSize int) ([]News, error) {
	var newsList []News
	query := db.Model(&News{})

	if source != "" {
		query = query.Where("source = ?", source)
	}

	if category != "" {
		query = query.Where("category = ?", category)
	}

	if err := query.Order("publish_time DESC").Offset(offset).Limit(pageSize).Find(&newsList).Error; err != nil {
		return nil, err
	}

	return newsList, nil
}

// GetNewsCountBySourceAndCategory 根据source和category获取新闻总数
func GetNewsCountBySourceAndCategory(db *gorm.DB, source, category string) (int64, error) {
	var count int64
	query := db.Model(&News{})

	if source != "" {
		query = query.Where("source = ?", source)
	}

	if category != "" {
		query = query.Where("category = ?", category)
	}

	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// UpdateNews 更新新闻
func UpdateNews(db *gorm.DB, news *News) error {
	return db.Save(news).Error
}

// DeleteNews 删除新闻
func DeleteNews(db *gorm.DB, id uint) error {
	return db.Delete(&News{}, id).Error
}

// GetOffset 计算offset
func GetOffset(page, pageSize int) int {
	if page <= 0 {
		page = 1
	}
	return (page - 1) * pageSize
}
