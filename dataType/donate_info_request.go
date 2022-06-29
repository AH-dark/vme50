package dataType

import (
	"mime/multipart"
)

type DonateInfoReq struct {
	Name    string                `form:"name" json:"name" validate:"required"`
	Email   string                `form:"email" json:"email" validate:"required"`
	Payment string                `form:"payment" json:"payment" validate:"required"`
	QRCode  *multipart.FileHeader `form:"qrcode" json:"-" validate:"required"`
}
