package services

import (
	"mxshs/movieLibrary/src/domain"
	"mxshs/movieLibrary/src/ports/repositories"
	"mxshs/movieLibrary/src/utils"
	"sync"
)

type ActorService struct {
	actorRepo      repositories.ActorRepository
	movieActorRepo repositories.MovieActorRepository
}

func NewActorService(actorRepo repositories.ActorRepository, movieActorRepo repositories.MovieActorRepository) *ActorService {
	return &ActorService{actorRepo, movieActorRepo}
}

func (as *ActorService) Create(name string, gender string, bd utils.Date) (*domain.Actor, error) {
	return as.actorRepo.CreateActor(name, gender, bd)
}

func (as *ActorService) GetActor(id int) (*domain.ActorDetail, error) {
	actor, err := as.actorRepo.GetActor(id)
	if err != nil {
		return nil, err
	}

	if movies, err := as.movieActorRepo.GetActorMovies(id); err == nil {
		actor.Movies = movies
	}

	return actor, nil
}

func (as *ActorService) GetActors() ([]*domain.ActorDetail, error) {
	actors, err := as.actorRepo.GetActors()
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup

	for i := range len(actors) {
		wg.Add(1)
		go func(id int) {
			movies, err := as.movieActorRepo.GetActorMovies(id)
			if err != nil {
				return
			}

			actors[i].Movies = movies
			wg.Done()
		}(i + 1)
	}

	wg.Wait()

	return actors, nil
}

func (as *ActorService) Update(id int, name string, gender string, bd utils.Date) (*domain.Actor, error) {
	return as.actorRepo.UpdateActor(id, name, gender, bd)
}

func (as *ActorService) Delete(id int) (int, error) {
	err := as.actorRepo.DeleteActor(id)

	return id, err
}
