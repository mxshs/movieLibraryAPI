package repository_adapter

import (
	"fmt"
	"mxshs/movieLibrary/src/domain"
	"time"

	_ "github.com/lib/pq"
)

func (pdb *PgDB) CreateMovie(title, description string, releaseDate time.Time, rating uint8) (*domain.Movie, error) {
	q, err := pdb.db.Query(
		`INSERT INTO movies(title, description, release_date, rating)
        VALUES ($1, $2, $3, $4)
        RETURNING mid;`,
		title,
		description,
		releaseDate,
		rating,
	)
	if err != nil {
		return nil, err
	}

	var mid int

	q.Scan(&mid)

	return &domain.Movie{
		Id:          mid,
		Title:       title,
		Description: description,
		ReleaseDate: releaseDate,
		Rating:      rating,
	}, nil
}

func (pdb *PgDB) GetMovie(mid int) (*domain.MovieDetail, error) {
	q, err := pdb.db.Query(
		`SELECT * FROM movies
        WHERE mid = $1`,
		mid,
	)
	if err != nil {
		return nil, err
	}

	if !q.Next() {
		return nil, fmt.Errorf(
			"movie with id %d not found",
			mid,
		)
	}

	var movie domain.MovieDetail

	err = q.Scan(&movie.Id, &movie.Title, &movie.Description, &movie.ReleaseDate, &movie.Rating)
	if err != nil {
		return nil, err
	}

	return &movie, nil
}

func (pdb *PgDB) GetMovies() ([]*domain.MovieDetail, error) {
	q, err := pdb.db.Query(
		`SELECT * FROM movies;`,
	)
	if err != nil {
		return nil, err
	}

	movies := []*domain.MovieDetail{}

	for q.Next() {
		var movie domain.MovieDetail

		err = q.Scan(&movie.Id, &movie.Title, &movie.Description, &movie.ReleaseDate, &movie.Rating)
		if err != nil {
			return nil, err
		}

		movies = append(movies, &movie)
	}

	return movies, nil
}

func (pdb *PgDB) SearchMoviesByTitle(sortKey, sortOrder, title string) ([]*domain.MovieDetail, error) {
	q, err := pdb.db.Query(
		`SELECT * FROM movies
        WHERE title SIMILAR TO $1;`,
		fmt.Sprintf("%%%s%%", title),
	)
	if err != nil {
		return nil, err
	}

	movies := []*domain.MovieDetail{}

	for q.Next() {
		var movie domain.MovieDetail

		err = q.Scan(&movie.Id, &movie.Title, &movie.Description, &movie.ReleaseDate, &movie.Rating)
		if err != nil {
			return nil, err
		}

		movies = append(movies, &movie)
	}

	return movies, nil
}

func (pdb *PgDB) SearchMovies(sortKey, sortOrder, title, actor string) ([]*domain.MovieDetail, error) {
	q, err := pdb.db.Query(
		`
        SELECT * FROM movies
        WHERE mid IN (
        SELECT movies.mid FROM movies
        LEFT JOIN movie_actors ON movies.mid=movie_actors.mid
        WHERE title SIMILAR TO $1
        INTERSECT
        SELECT mid FROM actors
        LEFT JOIN movie_actors ON actors.aid=movie_actors.aid
        WHERE name SIMILAR TO $2);
        `,
		fmt.Sprintf("%%%s%%", title),
		fmt.Sprintf("%%%s%%", actor),
	)
	if err != nil {
		return nil, err
	}

	movies := []*domain.MovieDetail{}

	for q.Next() {
		var movie domain.MovieDetail

		err = q.Scan(&movie.Id, &movie.Title, &movie.Description, &movie.ReleaseDate, &movie.Rating)
		if err != nil {
			return nil, err
		}

		movies = append(movies, &movie)
	}

	return movies, nil
}

func (pdb *PgDB) UpdateMovie(mid int, title, description string, releaseDate time.Time, rating uint8) (*domain.Movie, error) {
	q, err := pdb.db.Query(
		`UPDATE movies SET
        title = COALESCE($2, title),
        description = COALESCE($3, description),
        release_date = COALESCE($4, release_date),
        rating = COALESCE($5, rating),
        WHERE mid = $1
        RETURNING mid, title, description, release_date, rating;`,
		mid,
		title,
		description,
		releaseDate,
		rating,
	)
	if err != nil {
		return nil, err
	}

	var movie domain.Movie

	err = q.Scan(&movie)
	if err != nil {
		return nil, err
	}

	return &movie, nil
}

func (pdb *PgDB) DeleteMovie(mid int) error {
	_, err := pdb.db.Query(
		`DELETE FROM movies
        WHERE mid = $1;`,
		mid,
	)

	return err
}
