package util

import (
	"encoding/base64"
)

// const Passphrase = "abc"

// func Encrypt(input string, pass string) string {

// 	data := []byte(input)
// 	block, _ := aes.NewCipher([]byte(createHash(pass)))
// 	gcm, err := cipher.NewGCM(block)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	nonce := make([]byte, gcm.NonceSize())
// 	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
// 		panic(err.Error())
// 	}
// 	ciphertext := gcm.Seal(nonce, nonce, data, nil)
// 	return base64.StdEncoding.EncodeToString(ciphertext)
// }

// func createHash(key string) string {
// 	hasher := md5.New()
// 	hasher.Write([]byte(key))
// 	return hex.EncodeToString(hasher.Sum(nil))
// }

func Encode(b string) string {
	return base64.StdEncoding.EncodeToString([]byte(b))
}

func Decode(s string) string {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return string(data)
}

// func Decrypt(input string, Passphrase string) string {
// 	block, err := aes.NewCipher([]byte(Passphrase))
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	cipherText := Decode(input)
// 	cfb := cipher.NewCFBDecrypter(block, []byte(cipherText))
// 	plainText := make([]byte, len(cipherText))
// 	cfb.XORKeyStream(plainText, []byte(cipherText))
// 	return string(plainText)
// }
