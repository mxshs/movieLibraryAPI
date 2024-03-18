package adaptermodels

import (
	"mxshs/movieLibrary/src/handlers/utils"
)

type BaseMovie struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	ReleaseDate utils.Date `json:"releaseDate"`
	Rating      uint8      `json:"rating"`
}

type DetailMovie struct {
	Id          int         `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	ReleaseDate utils.Date  `json:"releaseDate"`
	Rating      uint8       `json:"rating"`
	Stars       []BaseActor `json:"stars"`
}
