package main

import (
	"fmt"

	"github.com/matthewyuh246/qrcode_go/db"
	"github.com/matthewyuh246/qrcode_go/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.QRCode{})
}
