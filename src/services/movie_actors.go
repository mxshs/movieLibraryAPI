package services

import (
	"mxshs/movieLibrary/src/domain"
	"mxshs/movieLibrary/src/ports/repositories"
)

type MovieActorService struct {
	movieActorRepo repositories.MovieActorRepository
}

func NewMovieActorService(movieActorRepo repositories.MovieActorRepository) *MovieActorService {
	return &MovieActorService{movieActorRepo}
}

func (mac *MovieActorService) GetMovieActors(id int) ([]*domain.Actor, error) {
	return mac.movieActorRepo.GetMovieActors(id)
}

func (mac *MovieActorService) GetActorMovies(id int) ([]*domain.Movie, error) {
	return mac.movieActorRepo.GetActorMovies(id)
}

func (mac *MovieActorService) CreateMovieActor(movieId, actorId int) error {
	return mac.movieActorRepo.CreateMovieActor(movieId, actorId)
}

func (mac *MovieActorService) DeleteMovieActor(movieId, actorId int) error {
	return mac.movieActorRepo.DeleteMovieActor(movieId, actorId)
}
