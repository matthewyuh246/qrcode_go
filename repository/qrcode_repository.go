package repository

import (
	"github.com/matthewyuh246/qrcode_go/model"
	"gorm.io/gorm"
)

type IQRCodeRepository interface {
	Save(qrCode *model.QRCode) error
	FindRecent(limit int, userId uint) ([]model.QRCode, error)
	FindFavorite(limit int, userId uint) ([]model.QRCode, error)
}

type qrcodeRepository struct {
	db *gorm.DB
}

func NewQRCodeRepository(db *gorm.DB) IQRCodeRepository {
	return &qrcodeRepository{db}
}

func (qr *qrcodeRepository) Save(qrCode *model.QRCode) error {
	if err := qr.db.Create(qrCode).Error; err != nil {
		return err
	}
	return nil
}

func (qr *qrcodeRepository) FindRecent(limit int, userId uint) ([]model.QRCode, error) {
	var qrCodes []model.QRCode
	if err := qr.db.Joins("User").Where("user_id=?", userId).Order("created_at desc").Limit(limit).Find(&qrCodes).Error; err != nil {
		return nil, err
	}
	return qrCodes, nil
}

func (qr *qrcodeRepository) FindFavorite(limit int, userId uint) ([]model.QRCode, error) {
	var qrCodes []model.QRCode
	if err := qr.db.Joins("User").Where("user_id=?", userId).Where("is_favorite = ?", true).Order("created_at desc").Limit(limit).Find(&qrCodes).Error; err != nil {
		return nil, err
	}
	return qrCodes, nil
}
