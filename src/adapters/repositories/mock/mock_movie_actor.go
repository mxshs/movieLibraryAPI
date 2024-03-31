package mock_db

import (
	"fmt"
	"mxshs/movieLibrary/src/domain"
)

func (mdb *MockDB) GetMovieActors(mid int) ([]*domain.Actor, error) {
	if _, ok := mdb.movies[mid]; !ok {
		return nil, fmt.Errorf("entity with id %d does not exist", mid)
	}

	result := []*domain.Actor{}

	ll, ok := mdb.movieActors[mid]
	if !ok {
		return result, nil
	}

	for aid := ll.head; aid != nil; aid = aid.next {
		actor := mdb.actors[aid.id]
		result = append(result, actor)
	}

	return result, nil
}

func (mdb *MockDB) GetActorMovies(aid int) ([]*domain.Movie, error) {
	if _, ok := mdb.actors[aid]; !ok {
		return nil, fmt.Errorf("entity with id %d does not exist", aid)
	}

	result := []*domain.Movie{}

	ll, ok := mdb.actorMovies[aid]
	if !ok {
		return result, nil
	}

	for mid := ll.head; mid != nil; mid = mid.next {
		movie := mdb.movies[mid.id]
		result = append(result, movie)
	}

	return result, nil
}

func (mdb *MockDB) CreateMovieActor(movieId, actorId int) error {
	if _, ok := mdb.actors[actorId]; !ok {
		return fmt.Errorf("entity with id %d does not exist", actorId)
	}
	if _, ok := mdb.movies[movieId]; !ok {
		return fmt.Errorf("entity with id %d does not exist", movieId)
	}

	if ll, ok := mdb.movieActors[movieId]; !ok {
		ll = &LL{}
		ll.Add(actorId)
		mdb.movieActors[movieId] = ll
	} else {
		ll.Add(actorId)
	}

	if ll, ok := mdb.actorMovies[actorId]; !ok {
		ll = &LL{}
		ll.Add(movieId)
		mdb.actorMovies[actorId] = ll
	} else {
		ll.Add(movieId)
	}

	return nil
}

func (mdb *MockDB) DeleteMovieActor(movieId, actorId int) error {
	if _, ok := mdb.actors[actorId]; !ok {
		return fmt.Errorf("entity with id %d does not exist", actorId)
	}
	if _, ok := mdb.movies[movieId]; !ok {
		return fmt.Errorf("entity with id %d does not exist", movieId)
	}

	if ll, ok := mdb.movieActors[movieId]; ok {
		ll.Remove(actorId)
	}

	if ll, ok := mdb.actorMovies[actorId]; !ok {
		ll.Remove(movieId)
	}

	return nil
}
