package mock_db

import (
	"fmt"
	"mxshs/movieLibrary/src/domain"
	"mxshs/movieLibrary/src/utils"
)

func (mdb *MockDB) CreateActor(name string, gender string, bd utils.Date) (*domain.Actor, error) {
	id := mdb.aid.Add(1)
	actor := &domain.Actor{
		Id:        int(id),
		Name:      name,
		Gender:    gender,
		Birthdate: bd,
	}

	mdb.actors[int(id)] = actor

	return actor, nil
}

func (mdb *MockDB) GetActor(id int) (*domain.ActorDetail, error) {
	actor, ok := mdb.actors[id]
	if !ok {
		return nil, fmt.Errorf("entity with id %d does not exist", id)
	}

	return &domain.ActorDetail{
		Id:        actor.Id,
		Name:      actor.Name,
		Gender:    actor.Gender,
		Birthdate: actor.Birthdate,
	}, nil
}

func (mdb *MockDB) GetActors() ([]*domain.ActorDetail, error) {
	actors := make([]*domain.ActorDetail, len(mdb.actors))

	idx := 0
	for _, actor := range mdb.actors {
		actors[idx] = &domain.ActorDetail{
			Id:        actor.Id,
			Name:      actor.Name,
			Gender:    actor.Gender,
			Birthdate: actor.Birthdate,
		}
		idx++
	}

	return actors, nil
}

func (mdb *MockDB) UpdateActor(id int, name, gender string, bd utils.Date) (*domain.Actor, error) {
	actor, ok := mdb.actors[id]
	if !ok {
		return nil, fmt.Errorf("entity with id %d does not exist", id)
	}

	if len(name) > 0 {
		actor.Name = name
	}

	if len(gender) > 0 {
		actor.Gender = gender
	}

	if !bd.IsZero() {
		actor.Birthdate = bd
	}

	return actor, nil
}

func (mdb *MockDB) DeleteActor(id int) error {
	if _, ok := mdb.actors[id]; !ok {
		return fmt.Errorf("entity with id %d does not exist", id)
	}

	delete(mdb.actors, id)

	return nil
}
