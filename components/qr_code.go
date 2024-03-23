package components

import (
	"encoding/base64"
	"fmt"
)

type Cryptocurrency string

const (
	Bitcoin  Cryptocurrency = "bitcoin"
	Ethereum Cryptocurrency = "ethereum"
)

type QrCode struct {
	Id   ComponentId `json:"id"`
	Data string      `json:"data"`
}

func NewQrCode(data []byte) *QrCode {
	return &QrCode{
		Id:   ComponentIdQrCode,
		Data: base64.StdEncoding.EncodeToString(data),
	}
}

func QrCodeEmail(email string) *QrCode {
	return NewQrCode([]byte(fmt.Sprintf("mailto:%s", email)))
}

func QrCodeURL(url string) *QrCode {
	return NewQrCode([]byte(url))
}

func QrCodePhone(phone string) *QrCode {
	return NewQrCode([]byte(fmt.Sprintf("tel:%s", phone)))
}

func QrCodeSMS(text string) *QrCode {
	return NewQrCode([]byte(fmt.Sprintf("smsto:%s", text)))
}

func QrCodeCoordinates(lat, lng float64) *QrCode {
	return NewQrCode([]byte(fmt.Sprintf("geo:%f,%f", lat, lng)))
}

func QrCodeCryptocurrencyAddress(crypto Cryptocurrency, address string) *QrCode {
	return NewQrCode([]byte(fmt.Sprintf("%s:%s", crypto, address)))
}
