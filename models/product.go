package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	NamaProduct string `gorm:"type:varchar(300);column:namaproduct" json:"namaproduct"`
	Deskripsi   string `gorm:"type:text;column:deskripsi" json:"deskripsi"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
