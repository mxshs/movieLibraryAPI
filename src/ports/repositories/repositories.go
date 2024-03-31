package repositories

import (
	"mxshs/movieLibrary/src/domain"
	"mxshs/movieLibrary/src/utils"
)

type ActorRepository interface {
	CreateActor(name string, gender string, bd utils.Date) (*domain.Actor, error)
	GetActor(id int) (*domain.ActorDetail, error)
	GetActors() ([]*domain.ActorDetail, error)
	UpdateActor(id int, name, gender string, bd utils.Date) (*domain.Actor, error)
	DeleteActor(id int) error
}

type MovieRepository interface {
	CreateMovie(title, description string, releaseDate utils.Date, rating uint8) (*domain.Movie, error)
	GetMovie(id int) (*domain.MovieDetail, error)
	GetMovies() ([]*domain.MovieDetail, error)
	SearchMoviesByTitle(sortKey, sortOrder, title string) ([]*domain.MovieDetail, error)
	SearchMovies(sortKey, sortOrder, title, actor string) ([]*domain.MovieDetail, error)
	UpdateMovie(id int, title, description string, releaseDate utils.Date, rating uint8) (*domain.Movie, error)
	DeleteMovie(id int) error
}

type MovieActorRepository interface {
	GetMovieActors(mid int) ([]*domain.Actor, error)
	GetActorMovies(aid int) ([]*domain.Movie, error)
	CreateMovieActor(movieId, actorId int) error
	DeleteMovieActor(movieId, actorId int) error
}

type UserRepository interface {
	CreateUser(username, password string, role domain.Role) (*domain.User, error)
	GetUser(id int) (*domain.User, error)
	GetUsers() ([]*domain.User, error)
	LoginUser(username, password string) (*domain.User, error)
	UpdateUser(id int, newUsername, newPassword string, role domain.Role) (*domain.User, error)
	DeleteUser(id int) error
}
