package model

import (
	"time"
	"gorm.io/gorm"
)

type Cart struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	UserID    uint           `json:"user_id" gorm:"not null"`
	ProductID uint           `json:"product_id" gorm:"not null"`
	Quantity  int            `json:"quantity" gorm:"not null;default:1"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	
	// 关联
	Product Product `json:"product" gorm:"foreignKey:ProductID"`
}

// TableName 指定表名
func (Cart) TableName() string {
	return "cart"
} 