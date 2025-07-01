package model

import (
	"time"
	"gorm.io/gorm"
)

type User struct {
	ID         uint           `json:"id" gorm:"primarykey"`
	OpenID     string         `json:"openid" gorm:"uniqueIndex;size:64;not null"`
	Nickname   string         `json:"nickname" gorm:"size:64"`
	Avatar     string         `json:"avatar" gorm:"size:255"`
	Phone      string         `json:"phone" gorm:"size:20"`
	Gender     int            `json:"gender" gorm:"default:0"` // 0:未知 1:男 2:女
	Country    string         `json:"country" gorm:"size:64"`
	Province   string         `json:"province" gorm:"size:64"`
	City       string         `json:"city" gorm:"size:64"`
	Language   string         `json:"language" gorm:"size:16"`
	Status     int            `json:"status" gorm:"default:1"` // 1:正常 0:禁用
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`
}

// TableName 指定表名
func (User) TableName() string {
	return "user"
} 