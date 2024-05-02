package database

import (
	"github.com/aadi-1024/StraySafe/backend/pkg/models"
	"gorm.io/gorm/clause"
)

//
// user related access functions
//

func (d *Database) NewIncident(incident models.Incident) error {
	res := d.Pool.Omit("resolverid").Create(&incident)
	return res.Error
}

func (d *Database) GetNearestNgo(latitude, longitude float32, num int) ([]*models.Ngo, error) {
	var ngos []*models.Ngo
	res := d.Pool.Clauses(clause.OrderBy{
		Expression: clause.Expr{SQL: "SQRT(POWER(latitude - ?,2) + POWER(longitude - ?,2))", Vars: []interface{}{[]float32{latitude}, []float32{longitude}}},
	}).Select("name", "about", "latitude", "longitude").Limit(num).Find(&ngos)
	return ngos, res.Error
}

func (d *Database) GetIncidents(uid int) ([]*models.Incident, error) {
	var incs []*models.Incident
	res := d.Pool.Table("incidents").Find(&incs, "uid = ?", uid).Limit(50)
	return incs, res.Error
}
