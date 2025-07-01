package model

import (
	"time"
	"gorm.io/gorm"
)

type Product struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	Name        string         `json:"name" gorm:"size:128;not null"`
	CategoryID  uint           `json:"category_id" gorm:"not null"`
	Price       float64        `json:"price" gorm:"type:decimal(10,2);not null"`
	Stock       int            `json:"stock" gorm:"not null;default:0"`
	CoverImg    string         `json:"cover_img" gorm:"size:255"`
	Images      string         `json:"images" gorm:"type:text"` // JSON格式存储多张图片
	Description string         `json:"description" gorm:"type:text"`
	Status      int            `json:"status" gorm:"default:1"` // 1:上架 0:下架
	Sort        int            `json:"sort" gorm:"default:0"`
	Sales       int            `json:"sales" gorm:"default:0"` // 销量
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	
	// 关联
	Category Category `json:"category" gorm:"foreignKey:CategoryID"`
}

// TableName 指定表名
func (Product) TableName() string {
	return "product"
}

type Category struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	Name      string         `json:"name" gorm:"size:64;not null"`
	Sort      int            `json:"sort" gorm:"default:0"`
	Status    int            `json:"status" gorm:"default:1"` // 1:启用 0:禁用
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// TableName 指定表名
func (Category) TableName() string {
	return "category"
} 