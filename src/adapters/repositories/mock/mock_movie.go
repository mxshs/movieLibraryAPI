package mock_db

import (
	"fmt"
	"mxshs/movieLibrary/src/domain"
	"mxshs/movieLibrary/src/utils"
	"strings"
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

func (mdb *MockDB) GetMovie(id int) (*domain.MovieDetail, error) {
	movie, ok := mdb.movies[id]
	if !ok {
		return nil, fmt.Errorf("entity with id %d does not exist", id)
	}

	return &domain.MovieDetail{
		Id:          movie.Id,
		Title:       movie.Title,
		Description: movie.Description,
		ReleaseDate: movie.ReleaseDate,
		Rating:      movie.Rating,
		Stars:       nil,
	}, nil
}

func (mdb *MockDB) GetMovies() ([]*domain.MovieDetail, error) {
	movies := make([]*domain.MovieDetail, len(mdb.movies))

	for idx, movie := range mdb.movies {
		movies[idx] = &domain.MovieDetail{
			Id:          movie.Id,
			Title:       movie.Title,
			Description: movie.Description,
			ReleaseDate: movie.ReleaseDate,
			Rating:      movie.Rating,
			Stars:       nil,
		}
	}

	return movies, nil
}

func (mdb *MockDB) UpdateMovie(id int, title, description string, releaseDate utils.Date, rating uint8) (*domain.Movie, error) {
	movie, ok := mdb.movies[id]
	if !ok {
		return nil, fmt.Errorf("entity with id %d does not exist", id)
	}

	if len(title) > 0 {
		movie.Title = title
	}

	if len(description) > 0 {
		movie.Description = description
	}

	if !releaseDate.IsZero() {
		movie.ReleaseDate = releaseDate
	}

	if rating != 0 {
		movie.Rating = rating
	}

	return movie, nil
}

func (mdb *MockDB) DeleteMovie(id int) error {
	if _, ok := mdb.movies[id]; !ok {
		return fmt.Errorf("entity with id %d does not exist", id)
	}

	delete(mdb.movies, id)

	return nil
}

func (mdb *MockDB) SearchMoviesByTitle(sortKey, sortOrder, title string) ([]*domain.MovieDetail, error) {
	result := []*domain.MovieDetail{}

	for _, movie := range mdb.movies {
		if strings.Contains(movie.Title, title) {
			m := &domain.MovieDetail{
				Id:          movie.Id,
				Title:       movie.Title,
				Description: movie.Description,
				ReleaseDate: movie.ReleaseDate,
				Rating:      movie.Rating,
				Stars:       nil,
			}

			stars, err := mdb.GetMovieActors(movie.Id)
			if err != nil {
				m.Stars = stars
			}

			result = append(result, m)
		}
	}

	return result, nil
}

func (mdb *MockDB) SearchMovies(sortKey, sortOrder, title, actor string) ([]*domain.MovieDetail, error) {
	result := []*domain.MovieDetail{}

	for _, movie := range mdb.movies {
		if strings.Contains(movie.Title, title) {
			stars, err := mdb.GetMovieActors(movie.Id)
			if err != nil {
				return nil, err
			}

			if len(actor) != 0 {
				for _, star := range stars {
					if strings.Contains(star.Name, actor) {
						m := &domain.MovieDetail{
							Id:          movie.Id,
							Title:       movie.Title,
							Description: movie.Description,
							ReleaseDate: movie.ReleaseDate,
							Rating:      movie.Rating,
							Stars:       stars,
						}

						result = append(result, m)
						break
					}
				}
			} else {
				m := &domain.MovieDetail{
					Id:          movie.Id,
					Title:       movie.Title,
					Description: movie.Description,
					ReleaseDate: movie.ReleaseDate,
					Rating:      movie.Rating,
					Stars:       stars,
				}

				result = append(result, m)
			}
		}
	}

	return result, nil
}
