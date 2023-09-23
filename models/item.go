package models

import "time"

type Item struct{
	Id int `gorm:"primaryKey"`
	ItemCode string `gorm:"not null; type:varchar(120)"`
	CustomerName string `gorm:"not null; type:varchar(120)"`
	Description string `gorm:"not null; type:varchar(255)"`
	Quantity int `gorm:"not null"`
	OrderId int

	CreatedAt time.Time
	UpdatedAt time.Time
}