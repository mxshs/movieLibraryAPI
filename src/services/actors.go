package services

import (
	"mxshs/movieLibrary/src/domain"
	"mxshs/movieLibrary/src/ports/repositories"
	"time"
)

type ActorService struct {
	actorRepo repositories.ActorRepository
}

func NewActorService(actorRepo repositories.ActorRepository) *ActorService {
	return &ActorService{actorRepo}
}

func (as *ActorService) Create(name string, gender string, bd time.Time) (*domain.Actor, error) {
	return as.actorRepo.CreateActor(name, gender, bd)
}

func (as *ActorService) GetActor(id int) (*domain.Actor, error) {
	return as.actorRepo.GetActor(id)
}

func (as *ActorService) GetActors() ([]*domain.Actor, error) {
	return as.actorRepo.GetActors()
}

func (as *ActorService) Update(id int, name string, gender string, bd time.Time) (*domain.Actor, error) {
	return as.actorRepo.UpdateActor(id, name, gender, bd)
}

func (as *ActorService) Delete(id int) (int, error) {
	err := as.actorRepo.DeleteActor(id)

	return id, err
}
