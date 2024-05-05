package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type BaseModel struct {
	gorm.Model

	Id int `gorm:"primary_key"`

	CreatedBy int `gorm:"not null"`

	UpdatedBy *sql.NullInt64 `gorm:"null"`

	DeletedBy *sql.NullInt64 `gorm:"null"`
}

func (m *BaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	value := tx.Statement.Context.Value("Id")
	var userId = &sql.NullInt64{Valid: false}
	if value != nil {
		userId = &sql.NullInt64{Valid: true, Int64: int64(value.(float64))}
	}
	m.UpdatedBy = userId
	return
}

func (m *BaseModel) BeforeDelete(tx *gorm.DB) (err error) {
	value := tx.Statement.Context.Value("UserId")
	var userId = &sql.NullInt64{Valid: false}
	if value != nil {
		userId = &sql.NullInt64{Valid: true, Int64: int64(value.(float64))}
	}
	m.DeletedBy = userId
	return
}
