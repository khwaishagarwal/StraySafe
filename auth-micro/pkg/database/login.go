package database

import (
	"github.com/aadi-1024/auth-micro/pkg/jwtUtil"
	"github.com/aadi-1024/auth-micro/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

// VerifyLogin verifies whether the credentials are correct and if so returns a JWT token
func (d *Database) VerifyLogin(user models.User, j *jwtUtil.JwtConfig) (string, error) {
	u := &models.User{}
	res := d.Pool.First(u, "email = ?", user.Email)
	if res.Error != nil {
		return "", res.Error
	}

	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}
	//if no err up till this point, password has been verified as correct
	//token, err := j.GenerateToken(u.Id, "usr")
	var token string
	if u.Admin {
		token, err = j.GenerateToken(u.Id, "adm")
	} else {
		token, err = j.GenerateToken(u.Id, "usr")
	}
	if err != nil {
		return "", err
	}
	return token, nil
}

func (d *Database) RegisterUser(user models.User) error {
	passHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), -1)
	if err != nil {
		return err
	}

	res := d.Pool.Create(&models.User{
		Username: user.Username,
		Email:    user.Email,
		Password: string(passHash),
	})
	return res.Error
}

func (d *Database) ResetPassword(user models.User, newPass []byte) error {
	u := &models.User{}
	res := d.Pool.First(u, "email = ?", user.Email)
	if res.Error != nil {
		return res.Error
	}

	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Password))
	if err != nil {
		return err
	}

	hash, err := bcrypt.GenerateFromPassword(newPass, -1)
	if err != nil {
		return err
	}

	return d.Pool.Model(u).Update("password", string(hash)).Error
}
