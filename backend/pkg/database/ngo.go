package database

import (
	"errors"
	"github.com/aadi-1024/StraySafe/backend/pkg/models"
	"gorm.io/gorm/clause"
)

func (d *Database) MarkResolved(incident models.Incident, ngoId int) error {
	res := d.Pool.Model(&incident).Where("resolved = ?", false).Updates(&models.Incident{
		Resolved:   true,
		ResolverId: ngoId,
	})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("invalid incident id")
	}
	return nil
}

func (d *Database) GetResolvedCases(ngoId int) (int, error) {
	var numResolved int

	res := d.Pool.Table("incidents").Select("COUNT(*)").Where("resolverid = ?", ngoId).Scan(&numResolved)

	return numResolved, res.Error
}

func (d *Database) GetNearestCases(latitude, longitude float32, limit int) ([]*models.Incident, error) {
	var inc []*models.Incident
	res := d.Pool.Table("incidents").Clauses(clause.OrderBy{
		Expression: clause.Expr{SQL: "SQRT(POWER(latitude - ?,2) + POWER(longitude - ?,2))", Vars: []interface{}{[]float32{latitude}, []float32{longitude}}},
	}).Limit(limit).Where("resolved = ?", false).Find(&inc)
	return inc, res.Error
}

func (d *Database) GetIncident(id int) (models.Incident, error) {
	inc := models.Incident{}

	res := d.Pool.First(&inc, id)
	return inc, res.Error
}
