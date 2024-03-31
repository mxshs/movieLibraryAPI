package services_test

import (
	mock_db "mxshs/movieLibrary/src/adapters/repositories/mock"
	"mxshs/movieLibrary/src/services"
	"mxshs/movieLibrary/src/utils"
	"time"
)

var as *services.ActorService
var db *mock_db.MockDB
var ms *services.MovieService
var us *services.UserService

func init() {
	db = mock_db.NewDB()
	as = services.NewActorService(db, db)
	ms = services.NewMovieService(db, db)
	us = services.NewUserService(db)
}

func dateHelper(ds string) utils.Date {
	date, _ := time.Parse("02.01.2006", ds)

	return utils.Date{Time: date}
}
