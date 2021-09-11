package dataBase

import (
	_ "embed"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//go:embed pass.txt
var pass string

var DB *gorm.DB

func Connect() error {
	var err error
	dns := fmt.Sprintf("host=localhost user=kuppuch password=%s dbname=parser port=5432 sslmode=disable", pass)
	DB, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		return err
	}
	migrate()
	return nil
}
