package models

type Book struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	Title       string `gorm:"type:varchar(500)" json:"title"`
	Author      string `gorm:"type:varchar(500)" json:"author"`
	Description string `gorm:"type:text" json:"description"`
	PublishDate string `gorm:"type:date" json:"publish_date"`
}
