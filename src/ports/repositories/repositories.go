package repositories

import (
	"mxshs/movieLibrary/src/domain"
	"time"
)

type ActorRepository interface {
	CreateActor(name string, gender string, bd time.Time) (*domain.Actor, error)
	GetActor(id int) (*domain.Actor, error)
	GetActors() ([]*domain.Actor, error)
	UpdateActor(id int, name, gender string, bd time.Time) (*domain.Actor, error)
	DeleteActor(id int) error
}

type MovieRepository interface {
	CreateMovie(title, description string, releaseDate time.Time, rating uint8) (*domain.Movie, error)
	GetMovie(id int) (*domain.Movie, error)
	GetMovies() ([]*domain.Movie, error)
	SearchMoviesByTitle(sortKey, sortOrder, title string) ([]*domain.Movie, error)
	SearchMovies(sortKey, sortOrder, title, actor string) ([]*domain.Movie, error)
	UpdateMovie(id int, title, description string, releaseDate time.Time, rating uint8) (*domain.Movie, error)
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
	GetUserByUsername(username string) (*domain.User, error)
	LoginUser(username, password string) (*domain.User, error)
	UpdateUser(username, newPassword string) (*domain.User, error)
	DeleteUser(username string) error
}
