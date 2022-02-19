package model

import (
	"time"
)

type Place struct {
	ID int `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Address string `db:"address" json:"address"`
	Description string `db:"description" json:"description"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
