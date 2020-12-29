package entity

import "time"

type Post struct {
	Id uint `gorm:"primaryKey"`
	Title string `gorm:"size:45"`
	Content string `gorm:"size:255"`
	CreatedAt time.Time

	User User `gorm:"foreignKey:UserEmail"`
	UserEmail string
}

type Posts []Post
