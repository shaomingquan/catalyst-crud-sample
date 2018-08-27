package store

import (
	"time"
)

type Test struct {
	ID        int       `gorm:"type:int(10)" json:"id"`
	CreatedAt time.Time `gorm:"type:datetime" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:datetime" json:"updated_at"`
	Name      string    `gorm:"type:varchar(255)" json:"name"`
	Score     int       `gorm:"type:int(11)" json:"score"`
}

func (table Test) TableName() string {
	return "test"
}
