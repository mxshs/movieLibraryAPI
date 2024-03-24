package repository_adapter

import (
	"fmt"
	"mxshs/movieLibrary/src/domain"
	"time"

	_ "github.com/lib/pq"
)

func (pdb *PgDB) CreateActor(name, gender string, bd time.Time) (*domain.Actor, error) {
	q, err := pdb.db.Query(
		`INSERT INTO actors(name, gender, birthdate)
        VALUES ($1, $2, $3)
        RETURNING aid;`,
		name,
		gender,
		bd,
	)
	if err != nil {
		return nil, err
	}

	if !q.Next() {
		return nil, fmt.Errorf(
			"unexpected empty return after succesful insertion")
	}

	var actorId int

	err = q.Scan(&actorId)
	if err != nil {
		return nil, err
	}

	return &domain.Actor{
		Id:        actorId,
		Name:      name,
		Gender:    gender,
		Birthdate: bd,
	}, nil
}

func (pdb *PgDB) GetActor(aid int) (*domain.ActorDetail, error) {
	q, err := pdb.db.Query(
		`SELECT * FROM actors
        WHERE aid = $1;`,
		aid,
	)
	if err != nil {
		return nil, err
	}

	if !q.Next() {
		return nil, fmt.Errorf(
			"actor with id %d not found",
			aid,
		)
	}

	var actor domain.ActorDetail

	err = q.Scan(&actor.Id, &actor.Name, &actor.Gender, &actor.Birthdate)
	if err != nil {
		return nil, err
	}

	return &actor, nil
}

func (pdb *PgDB) GetActors() ([]*domain.ActorDetail, error) {
	q, err := pdb.db.Query(
		`SELECT * FROM actors;`,
	)
	if err != nil {
		return nil, err
	}

	actors := []*domain.ActorDetail{}

	for q.Next() {
		var actor domain.ActorDetail

		err = q.Scan(&actor.Id, &actor.Name, &actor.Gender, &actor.Birthdate)
		if err != nil {
			return nil, err
		}

		actors = append(actors, &actor)
	}

	return actors, nil
}

func (pdb *PgDB) UpdateActor(aid int, name, gender string, bd time.Time) (*domain.Actor, error) {
	q, err := pdb.db.Query(
		`UPDATE actors SET
        name = COALESCE($2, name),
        gender = COALESCE($3, gender),
        birthdate = COALESCE($4, birthdate)
        WHERE aid = $1
        RETURNING id, name, gender, birthdate;`,
		aid,
		name,
		gender,
		bd,
	)
	if err != nil {
		return nil, err
	}

	var actor domain.Actor

	q.Scan(&actor)

	return &actor, nil
}

func (pdb *PgDB) DeleteActor(aid int) error {
	_, err := pdb.db.Query(
		`DELETE FROM actors
        WHERE aid = $1`,
		aid,
	)

	return err
}
