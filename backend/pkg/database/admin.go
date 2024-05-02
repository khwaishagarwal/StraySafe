package database

import "github.com/aadi-1024/StraySafe/backend/pkg/models"

func (d *Database) DeleteUser(id int) error {
	return d.Pool.Where("id = ?", id).Delete(&models.User{}).Error
}

func (d *Database) MakeAdmin(id int) error {
	return d.Pool.Model(&models.User{}).Where("id = ?", id).UpdateColumn("admin", true).Error
}

func (d *Database) DeleteNgo(id int) error {
	return d.Pool.Where("id = ?", id).Delete(&models.Ngo{}).Error
}
