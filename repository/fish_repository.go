package repository

import (
	"go_rest_api/model"
)

func FishList() ([]*model.Fish, error) {
	query := `SELECT * FROM fishes;`

	var fishes []*model.Fish
	if err := db.Select(&fishes, query); err != nil {
		return nil, err
	}

	return fishes, nil
}
