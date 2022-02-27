package model

import (
	"time"
)

type Fish struct {
	Id int `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Classification int8 `db:"classification" json:"classification"`
	Description string `db:"description" json:"description"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

// swagger 用。POST, PATCH するときのリクエストボディの指定がわからん
type FishParameter struct {
	Name string
	Classification int8
	Description string
}
