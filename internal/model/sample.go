package model

import "time"

// Sample represents the "sample" table.
type Sample struct {
	Id        string    `gorm:"column:id" json:"id"`                 // id
	Name      string    `gorm:"column:name" json:"name"`             // name
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"` // created_at
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"` // updated_at
}
