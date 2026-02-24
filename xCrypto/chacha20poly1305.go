package xCrypto

import (
	"github.com/tink-crypto/tink-go/v2/aead/subtle"
)

// ChaCha20Poly1305 key 为加密数据, associatedData 为任意关联数据(可以为空 []byte{})
type ChaCha20Poly1305 struct {
	Key            []byte
	AssociatedData []byte
}

// NewChaCha20Poly1305 新建 ChaCha20-Poly1305 加密
func NewChaCha20Poly1305(key []byte, associatedData []byte) *ChaCha20Poly1305 {
	return &ChaCha20Poly1305{
		Key:            key,
		AssociatedData: associatedData,
	}
}

func (c *ChaCha20Poly1305) Encrypt(plaintext []byte) ([]byte, error) {
	ca, err := subtle.NewChaCha20Poly1305(c.Key)
	if err != nil {
		return nil, err
	}
	return ca.Encrypt(plaintext, c.AssociatedData)
}

func (c *ChaCha20Poly1305) Decrypt(ciphertext []byte) ([]byte, error) {
	ca, err := subtle.NewChaCha20Poly1305(c.Key)
	if err != nil {
		return nil, err
	}
	return ca.Decrypt(ciphertext, c.AssociatedData)
}
