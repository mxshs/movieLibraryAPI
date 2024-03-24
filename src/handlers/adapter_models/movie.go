package adaptermodels

import (
	"mxshs/movieLibrary/src/handlers/utils"
)

type BaseMovie struct {
	Id          int        `json:"id" example:"0" format:"int64"`
	Title       string     `json:"title" example:"The Wolf of Wall-Street"`
	Description string     `json:"description" example:"Movie about some stuff" validate:"required" maxLength:"1500"`
	ReleaseDate utils.Date `json:"release_date" example:"25.12.2013" format:"date" validate:"required"`
	Rating      uint8      `json:"rating" example:"8" format:"uint8" validate:"required" minimum:"0" maximum:"10"`
}
