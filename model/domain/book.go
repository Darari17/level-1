package domain

import "time"

type Book struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement;type:bigint"`
	Title     string    `gorm:"column:title;type:varchar;size:255;not null"`
	Author    string    `gorm:"column:author;type:varchar;size:255;not null"`
	Year      int       `gorm:"column:year;type:int;not null"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (b *Book) TableName() string {
	return "books"
}
