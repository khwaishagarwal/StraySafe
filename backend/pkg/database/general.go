package database

import "github.com/aadi-1024/StraySafe/backend/pkg/models"

func (d *Database) GetNgoById(id int) (*models.Ngo, error) {
	ngo := &models.Ngo{}

	res := d.Pool.Select("name", "email", "about", "latitude", "longitude").First(ngo, id)
	return ngo, res.Error
}
