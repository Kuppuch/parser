package dataBase

import (
	"github.com/Kuppuch/parser/structs"
)

func migrate() {
	DB.AutoMigrate(&structs.Line{})
}
