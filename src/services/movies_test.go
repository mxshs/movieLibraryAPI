package services_test

import (
	"mxshs/movieLibrary/src/domain"
	"mxshs/movieLibrary/src/utils"
	"testing"
)

func TestGetMovie(t *testing.T) {
	tests := []struct {
		id       int
		err      bool
		expected *domain.MovieDetail
	}{
		{
			1,
			false,
			&domain.MovieDetail{
				Title:       "The Wolf of Wall Street",
				Description: "Amazing movie",
				ReleaseDate: dateHelper("17.12.2013"),
				Rating:      12,
				Stars: []*domain.Actor{
					{
						Name:      "Leonardo DiCaprio",
						Gender:    "male",
						Birthdate: dateHelper("11.11.1974"),
					},
					{
						Name:      "Matthew McConnaughey",
						Gender:    "male",
						Birthdate: dateHelper("04.11.1969"),
					},
				},
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
			"American Psycho",
			"Also good",
			dateHelper("14.04.2000"),
			12,
			false,
			domain.Movie{
				Title:       "American Psycho",
				Description: "Also good",
				ReleaseDate: dateHelper("14.04.2000"),
				Rating:      12,
			},
		},
		{
			"Titanic",
			"Good movie",
			dateHelper("19.12.1997"),
			11,
			false,
			domain.Movie{
				Title:       "Titanic",
				Description: "Good movie",
				ReleaseDate: dateHelper("19.12.1997"),
				Rating:      11,
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

		movie, err := ms.GetMovie(res.Id)
		if err != nil {
			panic(err)
		}

		if movie.Title != test.expected.Title {
			t.Fail()
		}

		if movie.Description != test.expected.Description {
			t.Fail()
		}

		if movie.ReleaseDate != test.expected.ReleaseDate {
			t.Fail()
		}

		if movie.Rating != test.expected.Rating {
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
			3,
			"",
			"",
			utils.Date{},
			12,
			false,
			domain.Movie{
				Title:       "American Psycho",
				Description: "Also good",
				ReleaseDate: dateHelper("14.04.2000"),
				Rating:      12,
			},
		},
		{
			4,
			"Rogue One",
			"Great",
			dateHelper("10.12.2016"),
			12,
			false,
			domain.Movie{
				Title:       "Rogue One",
				Description: "Great",
				ReleaseDate: dateHelper("10.12.2016"),
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

		movie, err := ms.GetMovie(res.Id)
		if err != nil {
			panic(err)
		}

		if movie.Title != test.expected.Title {
			t.Fail()
		}

		if movie.Description != test.expected.Description {
			t.Fail()
		}

		if movie.ReleaseDate != test.expected.ReleaseDate {
			t.Fail()
		}

		if movie.Rating != test.expected.Rating {
			t.Fail()
		}
	}
}

func TestGetMovies(t *testing.T) {
	tests := []struct {
		title    string
		actor    string
		err      bool
		expected []*domain.MovieDetail
	}{
		{
			"treet",
			"",
			false,
			[]*domain.MovieDetail{
				{
					Title:       "The Wolf of Wall Street",
					Description: "Amazing movie",
					ReleaseDate: dateHelper("17.12.2013"),
					Rating:      12,
					Stars: []*domain.Actor{
						{
							Name:      "Leonardo DiCaprio",
							Gender:    "male",
							Birthdate: dateHelper("11.11.1974"),
						},
						{
							Name:      "Matthew McConnaughey",
							Gender:    "male",
							Birthdate: dateHelper("04.11.1969"),
						},
					},
				},
			},
		},
		{
			"",
			"Capri",
			false,
			[]*domain.MovieDetail{
				{
					Title:       "The Wolf of Wall Street",
					Description: "Amazing movie",
					ReleaseDate: dateHelper("17.12.2013"),
					Rating:      12,
					Stars: []*domain.Actor{
						{
							Name:      "Leonardo DiCaprio",
							Gender:    "male",
							Birthdate: dateHelper("11.11.1974"),
						},
						{
							Name:      "Matthew McConnaughey",
							Gender:    "male",
							Birthdate: dateHelper("04.11.1969"),
						},
					},
				},
			},
		},
		{
			"asdasdasd",
			"",
			false,
			[]*domain.MovieDetail{},
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

		if len(res) != len(test.expected) {
			t.Fail()
		}

		for idx, movie := range res {
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

			for aidx, actor := range movie.Stars {
				if actor.Name != test.expected[idx].Stars[aidx].Name {
					t.Fail()
				}

				if actor.Gender != test.expected[idx].Stars[aidx].Gender {
					t.Fail()
				}

				if actor.Birthdate != test.expected[idx].Stars[aidx].Birthdate {
					t.Fail()
				}
			}
		}
	}
}
