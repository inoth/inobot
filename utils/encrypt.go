package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)

var DefulaKey = "inobotscriptencruptkey"

func Md5(data string) string {
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}

//默认加密
func DefukaEncrypt(data string) (string, bool) {
	ret, err := Encrypt(DefulaKey, data)
	if err != nil {
		return "", false
	}
	return ret, true
}

//默认解密
func DefukaDecrypt(data string) (string, bool) {
	ret, err := Decrypt(DefulaKey, data)
	if err != nil {
		return "", false
	}
	return ret, true
}

func getKeyBytes(key string) []byte {
	keyBytes := []byte(key)
	switch l := len(keyBytes); {
	case l < 16:
		keyBytes = append(keyBytes, make([]byte, 16-l)...)
	case l > 16:
		keyBytes = keyBytes[:16]
	}
	return keyBytes
}

func encrypt(key string, origData []byte) ([]byte, error) {
	keyBytes := getKeyBytes(key)
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, keyBytes[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func decrpt(key string, crypted []byte) ([]byte, error) {
	keyBytes := getKeyBytes(key)
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, keyBytes[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func Encrypt(key string, val string) (string, error) {
	origData := []byte(val)
	crypted, err := encrypt(key, origData)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(crypted), nil
}

func Decrypt(key string, val string) (string, error) {
	crypted, err := base64.URLEncoding.DecodeString(val)
	if err != nil {
		return "", err
	}
	origData, err := decrpt(key, crypted)
	if err != nil {
		return "", err
	}
	return string(origData), nil
}
