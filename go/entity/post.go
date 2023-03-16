package entity

import "time"

type Post struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Title     string    `gorm:"type:varchar(200)" json:"title"`
	Content   string    `gorm:"type:text" json:"content"`
	Category  string    `gorm:"type:varchar(100)" json:"category"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Status    string    `gorm:"type:varchar(100)" json:"status"`
}
