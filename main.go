package main

import (
	"github.com/matthewyuh246/qrcode_go/controller"
	"github.com/matthewyuh246/qrcode_go/db"
	"github.com/matthewyuh246/qrcode_go/repository"
	"github.com/matthewyuh246/qrcode_go/router"
	"github.com/matthewyuh246/qrcode_go/usecase"
	"github.com/matthewyuh246/qrcode_go/validator"
)

func main() {
	db := db.NewDB()
	userValidator := validator.NewUserValidator()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	userController := controller.NewUserController(userUsecase)
	qrcodeRepository := repository.NewQRCodeRepository(db)
	qrcodeUsecase := usecase.NewQRCodeUsecase(qrcodeRepository)
	qrcodeController := controller.NewQRCodeController(qrcodeUsecase)
	e := router.NewRouter(userController, qrcodeController)
	e.Logger.Fatal(e.Start(":8080"))
}
