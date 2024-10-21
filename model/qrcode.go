package model

import "time"

type QRCode struct {
	ID        int64     `gorm:"primaryKey" json:"id"`
	Title     string    `json:"Title"`
	Text      string    `json:"text" gorm:"not null"`
	Image     []byte    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User      `json:"user"`
	UserId    uint      `json:"user_id" gorm:"not null"`
}

type QRCodeResponse struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	TTitle    string    `json:"Title"`
	Text      string    `json:"text" gorm:"not null"`
	Image     []byte    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}