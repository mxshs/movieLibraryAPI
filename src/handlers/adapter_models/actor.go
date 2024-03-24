package adaptermodels

import "mxshs/movieLibrary/src/handlers/utils"

type BaseActor struct {
	Id        int        `json:"id" example:"0" format:"int64"`
	Name      string     `json:"name" example:"Leonardo DiCaprio" validate:"required"`
	Gender    string     `json:"gender" example:"male" validate:"required"`
	Birthdate utils.Date `json:"birthdate" example:"11.11.1974" format:"date"`
}
