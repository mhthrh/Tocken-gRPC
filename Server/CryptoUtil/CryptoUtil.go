package CryptoUtil

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

type CryptoInterface interface {
	Decrypt()
	Encrypt()
	Sha256()
	Md5Sum()
}
type Crypto struct {
	key      string
	FilePath string
	Text     string
	Result   string
}

func NewKey() *Crypto {
	c := new(Crypto)
	c.key = "AnKoloft@~delNazok!12345" // key parameter must be 16, 24 or 32,
	return c
}

func (k *Crypto) Encrypt() error {
	key := []byte(k.key)
	plaintext := []byte(k.Text)
	c, err := aes.NewCipher(key)
	if err != nil {
		return err

	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return err

	}

	k.Result = Byte64(gcm.Seal(nonce, nonce, plaintext, nil))
	return nil
}

func (k *Crypto) Decrypt() error {
	key := []byte(k.key)
	bb, _ := base64.StdEncoding.DecodeString(k.Text)
	ciphertext := []byte(bb)
	c, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return err
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	t, e := gcm.Open(nil, nonce, ciphertext, nil)
	if e != nil {
		return err
	}
	k.Result = BytesToString(t)
	return nil

}
func Byte64(msg []byte) string {
	return base64.StdEncoding.EncodeToString(msg)
}

func BytesToString(b []byte) string {
	return bytes.NewBuffer(b).String()
}
