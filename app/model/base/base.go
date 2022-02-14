package base

import (
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type BaseIDModel struct {
	// Id of as primary key
	// in: int64
	ID uint64 `gorm:"primary_key:auto_increment" json:"-"`

	// UID of the product
	// in: int64
	UID string `gorm:"uniqueIndex" json:"uid"`

	IsDeleted bool   `json:"is_deleted"`
	CreatedBy string `json:"-"`
	CreatedAt string `json:"-"`
	UpdatedBy string `json:"-"`
	UpdatedAt string `json:"-"`
}

func (base *BaseIDModel) BeforeCreate(tx *gorm.DB) error {
	uid, _ := gonanoid.New()
	tx.Statement.SetColumn("UID", uid)
	tx.Statement.SetColumn("IsDeleted", false)
	tx.Statement.SetColumn("CreatedAt", time.Now().UTC().Format("2006-01-02 15:04:05"))
	tx.Statement.SetColumn("UpdatedAt", time.Now().UTC().Format("2006-01-02 15:04:05"))
	return nil
}

func (base *BaseIDModel) BeforeUpdate(tx *gorm.DB) error {
	tx.Statement.SetColumn("UpdatedAt", time.Now().UTC().Format("2006-01-02 15:04:05"))
	return nil
}
