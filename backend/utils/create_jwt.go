package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Sub         int `json:"sub"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

func CreateJwt(secret string, data Payload) (string, error) {
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}
	headerBytArr, err := json.Marshal(header)

	if err != nil {
		return "", err
	}

	headerBase64 := base64UrlEncode(headerBytArr)

	dataBytArr, err := json.Marshal(data)

	if err != nil {
		return "", err
	}

	payloadBase64 := base64UrlEncode(dataBytArr)

	secretBytArr := []byte(secret)

	message := headerBase64 + "." + payloadBase64
	messageBytArr := []byte(message)

	hashed := hmac.New(sha256.New, secretBytArr)
	hashed.Write(messageBytArr)

	signature := hashed.Sum(nil)
	signatureBase64 := base64UrlEncode(signature)

	jwt := headerBase64 + "." + payloadBase64 + "." + signatureBase64

	return jwt, nil
}

func base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
