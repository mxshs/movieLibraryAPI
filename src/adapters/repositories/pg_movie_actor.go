package repository_adapter

import (
	"mxshs/movieLibrary/src/domain"
	_ "github.com/lib/pq"
)

func (pdb *PgDB) GetActorMovies(aid int) ([]*domain.Movie, error) {
	q, err := pdb.db.Query(
		`SELECT mid FROM movie_actors
        WHERE aid = $1;`,
		aid,
	)
	if err != nil {
		return nil, err
	}

	result := []*domain.Movie{}
	for q.Next() {
		var movie domain.Movie

		err = q.Scan(&movie)
		if err != nil {
			return nil, err
		}

		result = append(result, &movie)
	}

	return result, nil
}

func (pdb *PgDB) GetMovieActors(mid int) ([]*domain.Actor, error) {
	q, err := pdb.db.Query(
		`SELECT actors.aid, actors.name, actors.gender, actors.birthdate FROM movie_actors
        LEFT JOIN actors ON actors.aid = movie_actors.aid
        WHERE movie_actors.mid = $1;`,
		mid,
	)
	if err != nil {
		return nil, err
	}

	result := []*domain.Actor{}
	for q.Next() {
		var actor domain.Actor

		err = q.Scan(&actor.Id, &actor.Name, &actor.Gender, &actor.Birthdate)
		if err != nil {
			return nil, err
		}

		result = append(result, &actor)
	}

	return result, nil
}

func (pdb *PgDB) CreateMovieActor(mid, aid int) error {
	_, err := pdb.db.Query(
		`INSERT INTO movie_actors
        VALUES ($1, $2);`,
		mid,
		aid,
	)

	return err
}

func (pdb *PgDB) DeleteMovieActor(mid, aid int) error {
	_, err := pdb.db.Query(
		`DELETE FROM movie_actors
        WHERE mid = $1 AND aid = $2;`,
		mid,
		aid,
	)

	return err
}
