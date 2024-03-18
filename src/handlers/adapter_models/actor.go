package adaptermodels

import "mxshs/movieLibrary/src/handlers/utils"

type BaseActor struct {
	Id        int        `json:"id"`
	Name      string     `json:"name"`
	Gender    string     `json:"gender"`
	Birthdate utils.Date `json:"birthdate"`
}
