package db

import (
	"fmt"

	"gorm.io/gorm"
)

type Link struct{
	ID string `gorm:"primarykey;unique"`
	Name string `gorm:"notnull"`
	URL string `gorm:"notnull"`

}

func Migrations(db *gorm.DB) error {
	if err := db.AutoMigrate(&Link{}); err != nil{
		return fmt.Errorf("error al migrar las tablas: %w", err)
	}
	fmt.Println("Migraciones de tablas completadas.")
	return nil
}