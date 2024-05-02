package database

import (
	"github.com/aadi-1024/auth-micro/pkg/jwtUtil"
	"github.com/aadi-1024/auth-micro/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

func (d *Database) VerifyNgoLogin(ngo models.Ngo, j *jwtUtil.JwtConfig) (string, error) {
	u := &models.Ngo{}
	res := d.Pool.First(u).Where("email = ?", ngo.Email)
	if res.Error != nil {
		return "", res.Error
	}

	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(ngo.Password))
	if err != nil {
		return "", err
	}
	//if no err up till this point, password has been verified as correct
	token, err := j.GenerateToken(u.Id, "ngo")
	if err != nil {
		return "", err
	}
	return token, nil
}

func (d *Database) RegisterNgo(ngo models.Ngo) (int, error) {
	passHash, err := bcrypt.GenerateFromPassword([]byte(ngo.Password), -1)
	if err != nil {
		return -1, err
	}

	ngo.Password = string(passHash)

	res := d.Pool.Create(&ngo)
	return ngo.Id, res.Error
}
