package repository

import (
	"database/sql"
	"go_rest_api/model"
)

func FishList(name string) ([]*model.Fish, error) {
	query := `SELECT * FROM fishes WHERE name LIKE ?;`

	var fishes []*model.Fish
	if err := db.Select(&fishes, query, "%" + name + "%"); err != nil {
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

func FishGetById(id int) (*model.Fish, error) {
	query := `SELECT * FROM fishes WHERE id = ?;`

	var fish model.Fish

	if err := db.Get(&fish, query, id); err != nil {
		return nil, err
	}

	return &fish, nil
}

func FishUpdate(fish *model.Fish) (sql.Result, error) {
	updateQuery :=
		`UPDATE fishes
		    SET name = :name,
				    classification = :classification,
						description = :description
			WHERE id = :id`

	tx := db.MustBegin()
	res, err := tx.NamedExec(updateQuery, &fish);
	if err != nil {
		tx.Rollback()

		return res, err
	}

	tx.Commit()

	return res, nil
}

func FishDestroy(id int) (sql.Result, error) {
	query := `DELETE FROM fishes where id = ?`

	tx := db.MustBegin()
	res, err := tx.Exec(query, id);
	if err != nil {
		tx.Rollback()

		return res, err
	}

	tx.Commit()

	return res, nil
}
