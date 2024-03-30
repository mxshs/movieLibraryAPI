package mock_db

import (
	"fmt"
	"mxshs/movieLibrary/src/domain"
	"mxshs/movieLibrary/src/utils"
)

func (mdb *MockDB) CreateMovie(title, description string, releaseDate utils.Date, rating uint8) (*domain.Movie, error) {
	id := mdb.mid.Add(1)
	movie := &domain.Movie{
		Id:          int(id),
		Title:       title,
		Description: description,
		ReleaseDate: releaseDate,
		Rating:      rating,
	}

	mdb.movies[int(id)] = movie

	return movie, nil
}

func (mdb *MockDB) GetMovie(id int) (*domain.Movie, error) {
	movie, ok := mdb.movies[id]
	if !ok {
		return nil, fmt.Errorf("Entity with id %d does not exist", id)
	}

	return movie, nil
}

func (mdb *MockDB) GetMovies() ([]*domain.Movie, error) {
	movies := make([]*domain.Movie, len(mdb.movies))

	for _, movie := range mdb.movies {
		movies = append(movies, movie)
	}

	return movies, nil
}

func (mdb *MockDB) UpdateMovie(id int, title, description string, releaseDate utils.Date, rating uint8) (*domain.Movie, error) {
	movie, ok := mdb.movies[id]
	if !ok {
		return nil, fmt.Errorf("Entity with id %d does not exist", id)
	}

	if len(title) > 0 {
		movie.Title = title
	}

	if len(description) > 0 {
		movie.Description = description
	}

	if releaseDate.Unix() != 0 {
		movie.ReleaseDate = releaseDate
	}

	if rating != 0 {
		movie.Rating = rating
	}

	return movie, nil
}

func (mdb *MockDB) DeleteMovie(id int) (int, error) {
	if _, ok := mdb.movies[id]; !ok {
		return 0, fmt.Errorf("entity with id %d does not exist", id)
	}

	delete(mdb.movies, id)

	return id, nil
}
