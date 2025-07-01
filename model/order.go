package model

import (
	"time"
	"gorm.io/gorm"
)

type Order struct {
	ID         uint           `json:"id" gorm:"primarykey"`
	UserID     uint           `json:"user_id" gorm:"not null"`
	OrderNo    string         `json:"order_no" gorm:"uniqueIndex;size:64;not null"`
	TotalPrice float64        `json:"total_price" gorm:"type:decimal(10,2);not null"`
	Status     int            `json:"status" gorm:"default:0"` // 0:待付款 1:待发货 2:待收货 3:已完成 4:已取消
	AddressID  uint           `json:"address_id" gorm:"not null"`
	Remark     string         `json:"remark" gorm:"size:255"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`
	
	// 关联
	User    User         `json:"user" gorm:"foreignKey:UserID"`
	Address Address      `json:"address" gorm:"foreignKey:AddressID"`
	Items   []OrderItem  `json:"items" gorm:"foreignKey:OrderID"`
}

// TableName 指定表名
func (Order) TableName() string {
	return "order"
}

type OrderItem struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	OrderID   uint           `json:"order_id" gorm:"not null"`
	ProductID uint           `json:"product_id" gorm:"not null"`
	Quantity  int            `json:"quantity" gorm:"not null"`
	Price     float64        `json:"price" gorm:"type:decimal(10,2);not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	
	// 关联
	Product Product `json:"product" gorm:"foreignKey:ProductID"`
}

// TableName 指定表名
func (OrderItem) TableName() string {
	return "order_item"
}

type Address struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	UserID    uint           `json:"user_id" gorm:"not null"`
	Name      string         `json:"name" gorm:"size:64;not null"`
	Phone     string         `json:"phone" gorm:"size:20;not null"`
	Province  string         `json:"province" gorm:"size:64;not null"`
	City      string         `json:"city" gorm:"size:64;not null"`
	District  string         `json:"district" gorm:"size:64;not null"`
	Detail    string         `json:"detail" gorm:"size:255;not null"`
	IsDefault bool           `json:"is_default" gorm:"default:false"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// TableName 指定表名
func (Address) TableName() string {
	return "address"
} 