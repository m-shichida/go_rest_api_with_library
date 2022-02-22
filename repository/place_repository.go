package repository

import (
	"go_rest_api/model"
)

func PlaceList() ([]*model.Place, error) {
	query := `SELECT * FROM places;`

	var places []*model.Place
	if err := db.Select(&places, query); err != nil {
		return nil, err
	}

	return places, nil
}
