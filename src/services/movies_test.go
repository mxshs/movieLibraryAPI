package services_test

import (
	"mxshs/movieLibrary/src/domain"
	"mxshs/movieLibrary/src/utils"
	"testing"
)

func TestCreateMovie(t *testing.T) {
	tests := []struct {
		title       string
		description string
		releaseDate utils.Date
		rating      uint8
		err         bool
		expected    domain.Movie
	}{
		{
			"The Wolf of Wall Street",
			"amazing movie",
			dateHelper("17.12.2013"),
			12,
			false,
			domain.Movie{
				Id:          1,
				Title:       "The Wolf of Wall Street",
				Description: "amazing movie",
				ReleaseDate: dateHelper("17.12.2013"),
				Rating:      12,
			},
		},
		{
			"American Psycho",
			"also good",
			dateHelper("14.04.2000"),
			12,
			false,
			domain.Movie{
				Id:          2,
				Title:       "American Psycho",
				Description: "also good",
				ReleaseDate: dateHelper("14.04.2000"),
				Rating:      12,
			},
		},
	}

	for _, test := range tests {
		res, err := ms.CreateMovie(test.title, test.description, test.releaseDate, test.rating)
		if err != nil {
			if test.err {
				return
			}

			t.Fail()
		}

		if res.Id != test.expected.Id {
			t.Fail()
		}

		if res.Title != test.expected.Title {
			t.Fail()
		}

		if res.Description != test.expected.Description {
			t.Fail()
		}

		if res.ReleaseDate != test.expected.ReleaseDate {
			t.Fail()
		}

		if res.Rating != test.expected.Rating {
			t.Fail()
		}
	}
}

func TestUpdateMovie(t *testing.T) {
	tests := []struct {
		id          int
		title       string
		description string
		releaseDate utils.Date
		rating      uint8
		err         bool
		expected    domain.Movie
	}{
		{
			1,
			"Titanic",
			"good movie",
			dateHelper("19.12.1997"),
			11,
			false,
			domain.Movie{
				Id:          1,
				Title:       "Titanic",
				Description: "good movie",
				ReleaseDate: dateHelper("19.12.1997"),
				Rating:      11,
			},
		},
		{
			2,
			"",
			"",
			utils.Date{},
			12,
			false,
			domain.Movie{
				Id:          2,
				Title:       "American Psycho",
				Description: "also good",
				ReleaseDate: dateHelper("14.04.2000"),
				Rating:      12,
			},
		},
	}

	for _, test := range tests {
		res, err := ms.UpdateMovie(test.id, test.title, test.description, test.releaseDate, test.rating)
		if err != nil {
			if test.err {
				return
			}

			t.Fail()
		}

		if res.Id != test.expected.Id {
			t.Fail()
		}

		if res.Title != test.expected.Title {
			t.Fail()
		}

		if res.Description != test.expected.Description {
			t.Fail()
		}

		if res.ReleaseDate != test.expected.ReleaseDate {
			t.Fail()
		}

		if res.Rating != test.expected.Rating {
			t.Fail()
		}
	}
}

func TestGetMovie(t *testing.T) {
	tests := []struct {
		id       int
		err      bool
		expected *domain.MovieDetail
	}{
		{
			2,
			false,
			&domain.MovieDetail{
				Id:          2,
				Title:       "American Psycho",
				Description: "also good",
				ReleaseDate: dateHelper("14.04.2000"),
				Rating:      12,
			},
		},
		{
			1,
			false,
			&domain.MovieDetail{
				Id:          1,
				Title:       "Titanic",
				Description: "good movie",
				ReleaseDate: dateHelper("19.12.1997"),
				Rating:      11,
			},
		},
		{
			3,
			true,
			nil,
		},
	}

	for _, test := range tests {
		res, err := ms.GetMovie(test.id)
		if err != nil {
			if test.err {
				return
			}

			t.Fail()
		}

		if res.Id != test.expected.Id {
			t.Fail()
		}

		if res.Title != test.expected.Title {
			t.Fail()
		}

		if res.Description != test.expected.Description {
			t.Fail()
		}

		if res.ReleaseDate != test.expected.ReleaseDate {
			t.Fail()
		}

		if res.Rating != test.expected.Rating {
			t.Fail()
		}
	}
}

func TestGetMovies(t *testing.T) {
	db.CreateMovieActor(2, 2)
	tests := []struct {
		title    string
		actor    string
		err      bool
		expected []*domain.MovieDetail
	}{
		{
			"sycho",
			"",
			false,
			[]*domain.MovieDetail{
				{
					Id:          2,
					Title:       "American Psycho",
					Description: "also good",
					ReleaseDate: dateHelper("14.04.2000"),
					Rating:      12,
				},
			},
		},
		{
			"",
			"ale",
			false,
			[]*domain.MovieDetail{
				{
					Id:          2,
					Title:       "American Psycho",
					Description: "also good",
					ReleaseDate: dateHelper("14.04.2000"),
					Rating:      12,
				},
			},
		},
	}

	for _, test := range tests {
		res, err := ms.GetMovies("", "", test.title, test.actor)
		if err != nil {
			if test.err {
				return
			}

			t.Fail()
		}

		for idx, movie := range res {
			if movie.Id != test.expected[idx].Id {
				t.Fail()
			}

			if movie.Title != test.expected[idx].Title {
				t.Fail()
			}

			if movie.Description != test.expected[idx].Description {
				t.Fail()
			}

			if movie.ReleaseDate != test.expected[idx].ReleaseDate {
				t.Fail()
			}

			if movie.Rating != test.expected[idx].Rating {
				t.Fail()
			}
		}
	}
}
