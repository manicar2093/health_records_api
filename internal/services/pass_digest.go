package services

import "golang.org/x/crypto/bcrypt"

type (
	HashPassword interface {
		Digest(string) (string, error)
	}
	PasswordComparable interface {
		Compare(hashed, password string) error
	}
	BcryptImpl struct{}
)

func NewBcryptImpl() *BcryptImpl {
	return &BcryptImpl{}
}

func (c *BcryptImpl) Digest(p string) (string, error) {
	passAsBytes := []byte(p)
	digested, err := bcrypt.GenerateFromPassword(passAsBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(digested), nil
}

func (c *BcryptImpl) Compare(hashed, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
}
