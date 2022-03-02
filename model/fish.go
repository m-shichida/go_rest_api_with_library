package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Fish struct {
	Id int `db:"id" json:"id"`
	Name string `db:"name" json:"name" validate:"required,max=100"`
	Classification int8 `db:"classification" json:"classification" validate:"required,gte=0,lte=2"`
	Description string `db:"description" json:"description" validate:"required"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

// swagger 用。POST, PATCH するときのリクエストボディの指定がわからん
type FishParameter struct {
	Name string
	Classification int8
	Description string
}

func (fish *Fish) ValidationMessages(err error) []string {
	var messages []string

	for _, err := range err.(validator.ValidationErrors) {
		switch (err.Field()) {
			case "Name":
				switch (err.Tag()) {
					case "required":
						messages = append(messages, "名前は必須項目です")
					case "max":
						messages = append(messages, "名前は101文字以上入力できません")
				}
			case "Classification":
				switch (err.Tag()) {
					case "required":
						messages = append(messages, "分類は必須項目です")
					case "lte":
						messages = append(messages, "分類は0から2の数値を指定してください")
					case "gte":
						messages = append(messages, "分類は0から2の数値を指定してください")
			}
			case "Description":
				if err.Tag() == "required" {
					messages = append(messages, "説明は必須項目です")
				}
		}
	}

	return messages
}
