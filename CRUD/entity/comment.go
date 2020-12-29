package entity

type Comment struct {
	Id uint `gorm:"primaryKey"`
	Title string `gorm:"size:45"`
	Content string `gorm:"size:255"`

	Post Post `gorm:"foreignKey:PostId"`
	PostId uint

	User User `gorm:"foreignKey:PostUserEmail"`
	PostUserEmail string
}

type Comments []Comment
