package repository

import (
	"database/sql"
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

func FishCreate(fish *model.Fish) (sql.Result, error) {
	query :=
		`INSERT INTO fishes(name, classification, description)
		      VALUES (:name, :classification, :description)`

	tx := db.MustBegin()
	result, err := tx.NamedExec(query, &fish);
	if err != nil {
		tx.Rollback()

		return nil, err
	}

	tx.Commit()

	return result, nil
}
