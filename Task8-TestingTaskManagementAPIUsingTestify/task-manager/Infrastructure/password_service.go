package Infrastructure

import "golang.org/x/crypto/bcrypt"

type PasswordService struct{}

func NewPassword() *PasswordService {
	return &PasswordService{}
}

func (*PasswordService) ValidatePasswordHash(existingUserPassword string, claimPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(existingUserPassword), []byte(claimPassword))
}
func (*PasswordService) GeneratePasswordHash(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}
