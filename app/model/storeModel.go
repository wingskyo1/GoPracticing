package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Store struct {
	gorm.Model
	Name     string     `gorm:"unique" json:"name"`
	Open     bool       `json:"open"`
	Menulist []Menulist `gorm:"ForeignKey:StoreID" json:"menulist"`
}

type Menulist struct {
	gorm.Model
	Name        string `gorm:"unique" json:"name"`
	Description string `json:"description"`
	StoreID     uint   `json:"store_id"`
}

type order struct {
	gorm.Model
	Name        string     `gorm:"unique" json:"name"`
	Description string     `json:"description"`
	OrderList   []Menulist `gorm:"ForeignKey:OrderID" json:"orderList"`
}

type orderList struct {
	gorm.Model
	Name        string `gorm:"unique" json:"name"`
	Description string `json:"description"`
	OrderID     uint   `json:"order_id"`
}
