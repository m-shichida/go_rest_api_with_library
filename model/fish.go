package model

import (
	"time"
)

type Fish struct {
	ID int `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Classification int8 `db:"classification" json:"classification"`
	Description string `db:"description" json:"description"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
