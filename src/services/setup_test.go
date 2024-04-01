package services_test

import (
	mock_db "mxshs/movieLibrary/src/adapters/repositories/mock"
	"mxshs/movieLibrary/src/domain"
	"mxshs/movieLibrary/src/services"
	"mxshs/movieLibrary/src/utils"
	"time"
)

var as *services.ActorService
var db *mock_db.MockDB
var ms *services.MovieService
var us *services.UserService
var ma *services.MovieActorService

func init() {
	db = mock_db.NewDB()
	as = services.NewActorService(db, db)
	ms = services.NewMovieService(db, db)
	us = services.NewUserService(db)
	ma = services.NewMovieActorService(db)

	setupFixtures()
}

func dateHelper(ds string) utils.Date {
	date, _ := time.Parse("02.01.2006", ds)

	return utils.Date{Time: date}
}

func setupFixtures() {
	_, err := db.CreateActor("Leonardo DiCaprio", "male", dateHelper("11.11.1974"))
	if err != nil {
		panic(err)
	}

	_, err = db.CreateActor("Matthew McConnaughey", "male", dateHelper("04.11.1969"))
	if err != nil {
		panic(err)
	}

	_, err = db.CreateMovie("The Wolf of Wall Street", "Amazing movie", dateHelper("17.12.2013"), 12)
	if err != nil {
		panic(err)
	}

	_, err = db.CreateMovie("Interstellar", "Insane", dateHelper("26.10.2014"), 12)
	if err != nil {
		panic(err)
	}

	err = db.CreateMovieActor(1, 1)
	if err != nil {
		panic(err)
	}

	err = db.CreateMovieActor(1, 2)
	if err != nil {
		panic(err)
	}

	_, err = db.CreateUser("test", "testpassword", domain.USR)
	if err != nil {
		panic(err)
	}
}
