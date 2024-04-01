package services_test

import (
	"mxshs/movieLibrary/src/domain"
	"testing"
)

func TestGetMovieActors(t *testing.T) {
	tests := []struct {
		mid      int
		err      bool
		expected []*domain.Actor
	}{
		{
			1,
			false,
			[]*domain.Actor{
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
		{
			50,
			true,
			nil,
		},
	}

	for _, test := range tests {
		actors, err := ma.GetMovieActors(test.mid)
		if err != nil {
			if test.err {
				return
			}

			t.Fail()
		}

		for idx, actor := range actors {
			if actor.Name != test.expected[idx].Name {
				t.Fail()
			}

			if actor.Gender != test.expected[idx].Gender {
				t.Fail()
			}

			if actor.Birthdate != test.expected[idx].Birthdate {
				t.Fail()
			}
		}
	}
}

func TestGetActorMovies(t *testing.T) {
	tests := []struct {
		aid      int
		err      bool
		expected []*domain.Movie
	}{
		{
			1,
			false,
			[]*domain.Movie{
				{
					Title:       "The Wolf of Wall Street",
					Description: "Amazing movie",
					ReleaseDate: dateHelper("17.12.2013"),
					Rating:      12,
				},
			},
		},
		{
			50,
			true,
			nil,
		},
	}

	for _, test := range tests {
		movies, err := ma.GetActorMovies(test.aid)
		if err != nil {
			if test.err {
				return
			}

			t.Fail()
		}

		for idx, movie := range movies {
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

func TestCreateMovieActor(t *testing.T) {
	tests := []struct {
		mid      int
		aid      int
		err      bool
		expected []*domain.Actor
	}{
		{
			2,
			2,
			false,
			[]*domain.Actor{
				{
					Name:      "Matthew McConnaughey",
					Gender:    "male",
					Birthdate: dateHelper("04.11.1969"),
				},
			},
		},
		{
			50,
			2,
			true,
			nil,
		},
	}

	for _, test := range tests {
		err := ma.CreateMovieActor(test.mid, test.aid)
		if err != nil {
			if test.err {
				return
			}

			t.Fail()
		}

		actors, err := ma.GetMovieActors(test.mid)
		if err != nil {
			panic(err)
		}

		for idx, actor := range actors {
			if actor.Name != test.expected[idx].Name {
				t.Fail()
			}

			if actor.Gender != test.expected[idx].Gender {
				t.Fail()
			}

			if actor.Birthdate != test.expected[idx].Birthdate {
				t.Fail()
			}
		}
	}
}

func TestDeleteMovieActor(t *testing.T) {
	tests := []struct {
		mid      int
		aid      int
		err      bool
		expected []*domain.Actor
	}{
		{
			2,
			2,
			false,
			[]*domain.Actor{
				{
					Name:      "Matthew McConaughey",
					Gender:    "male",
					Birthdate: dateHelper("04.11.1969"),
				},
			},
		},
		{
			50,
			2,
			true,
			nil,
		},
	}

	for _, test := range tests {
		err := ma.DeleteMovieActor(test.mid, test.aid)
		if err != nil {
			if test.err {
				return
			}

			t.Fail()
		}

		actors, err := ma.GetMovieActors(test.mid)
		if err != nil {
			panic(err)
		}

		for idx, actor := range actors {
			if actor.Name != test.expected[idx].Name {
				t.Fail()
			}

			if actor.Gender != test.expected[idx].Gender {
				t.Fail()
			}

			if actor.Birthdate != test.expected[idx].Birthdate {
				t.Fail()
			}
		}
	}
}
