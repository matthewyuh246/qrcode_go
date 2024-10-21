package usecase

import (
	"github.com/matthewyuh246/qrcode_go/model"
	"github.com/matthewyuh246/qrcode_go/repository"
	"github.com/skip2/go-qrcode"
)

type IQRCodeUsecase interface {
	GenerateQRCode(text string) (*model.QRCode, error)
	SaveQRCode(qrCode *model.QRCode) error
	GetRecentQRCodes(limit int, userId uint) ([]model.QRCode, error)
}

type qrcodeUsecase struct {
	qr repository.IQRCodeRepository
}

func NewQRCodeUsecase(qr repository.IQRCodeRepository) IQRCodeUsecase {
	return &qrcodeUsecase{qr}
}

func (qu *qrcodeUsecase) GenerateQRCode(text string) (*model.QRCode, error) {
	qr, err := qrcode.New(text, qrcode.Medium)
	if err != nil {
		return nil, err
	}

	data, err := qr.PNG(256)
	if err != nil {
		return nil, err
	}

	return &model.QRCode{
		Text:  text,
		Image: data,
	}, nil
}

func (qu *qrcodeUsecase) SaveQRCode(qrCode *model.QRCode) error {
	return qu.qr.Save(qrCode)
}

func (qu *qrcodeUsecase) GetRecentQRCodes(limit int, userId uint) ([]model.QRCode, error) {
	return qu.qr.FindRecent(limit, userId)
}
