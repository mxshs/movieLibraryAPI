package services_test

import (
	"mxshs/movieLibrary/src/domain"
	"mxshs/movieLibrary/src/utils"
	"testing"
)

func TestCreateActor(t *testing.T) {
	tests := []struct {
		name      string
		gender    string
		birthdate utils.Date
		err       bool
		expected  domain.Actor
	}{
		{
			"Leonardo DiCaprio",
			"male",
			dateHelper("11.11.1974"),
			false,
			domain.Actor{
				Id:        1,
				Name:      "Leonardo DiCaprio",
				Gender:    "male",
				Birthdate: dateHelper("11.11.1974"),
			},
		},
		{
			"Kate Winslet",
			"female",
			dateHelper("05.10.1975"),
			false,
			domain.Actor{
				Id:        2,
				Name:      "Kate Winslet",
				Gender:    "female",
				Birthdate: dateHelper("05.10.1975"),
			},
		},
	}

	for _, test := range tests {
		res, err := as.CreateActor(test.name, test.gender, test.birthdate)
		if err != nil {
			if test.err {
				return
			}

			t.Fail()
		}

		if res.Id != test.expected.Id {
			t.Fail()
		}

		if res.Name != test.expected.Name {
			t.Fail()
		}

		if res.Gender != test.expected.Gender {
			t.Fail()
		}

		if res.Birthdate != test.expected.Birthdate {
			t.Fail()
		}
	}
}

func TestUpdateActor(t *testing.T) {
	tests := []struct {
		id           int
		newName      string
		newGender    string
		newBirthdate utils.Date
		err          bool
		expected     *domain.Actor
	}{
		{
			2,
			"Christian Bale",
			"male",
			dateHelper("30.01.1974"),
			false,
			&domain.Actor{
				Id:        2,
				Name:      "Christian Bale",
				Gender:    "male",
				Birthdate: dateHelper("30.01.1974"),
			},
		},
		{
			1,
			"",
			"",
			utils.Date{},
			false,
			&domain.Actor{
				Id:        1,
				Name:      "Leonardo DiCaprio",
				Gender:    "male",
				Birthdate: dateHelper("11.11.1974"),
			},
		},
		{
			3,
			"",
			"",
			utils.Date{},
			true,
			nil,
		},
	}

	for _, test := range tests {
		res, err := as.UpdateActor(test.id, test.newName, test.newGender, test.newBirthdate)
		if err != nil {
			if test.err {
				return
			}

			t.Fail()
		}

		if res.Id != test.expected.Id {
			t.Fail()
		}

		if res.Name != test.expected.Name {
			t.Fail()
		}

		if res.Gender != test.expected.Gender {
			t.Fail()
		}

		if res.Birthdate != test.expected.Birthdate {
			t.Fail()
		}
	}
}

func TestGetActor(t *testing.T) {
	tests := []struct {
		id       int
		err      bool
		expected *domain.ActorDetail
	}{
		{
			2,
			false,
			&domain.ActorDetail{
				Id:        2,
				Name:      "Christian Bale",
				Gender:    "male",
				Birthdate: dateHelper("30.01.1974"),
			},
		},
		{
			1,
			false,
			&domain.ActorDetail{
				Id:        1,
				Name:      "Leonardo DiCaprio",
				Gender:    "male",
				Birthdate: dateHelper("11.11.1974"),
			},
		},
		{
			3,
			true,
			nil,
		},
	}

	for _, test := range tests {
		res, err := as.GetActor(test.id)
		if err != nil {
			if test.err {
				return
			}

			t.Fail()
		}

		if res.Id != test.expected.Id {
			t.Fail()
		}

		if res.Name != test.expected.Name {
			t.Fail()
		}

		if res.Gender != test.expected.Gender {
			t.Fail()
		}

		if res.Birthdate != test.expected.Birthdate {
			t.Fail()
		}
	}
}

func TestGetActors(t *testing.T) {
	tests := []struct {
		expected []*domain.ActorDetail
	}{
		{
			[]*domain.ActorDetail{
				{
					Id:        1,
					Name:      "Leonardo DiCaprio",
					Gender:    "male",
					Birthdate: dateHelper("11.11.1974"),
				},
				{
					Id:        2,
					Name:      "Christian Bale",
					Gender:    "male",
					Birthdate: dateHelper("30.01.1974"),
				},
			},
		},
	}

	for _, test := range tests {
		res, err := as.GetActors()
		if err != nil {
			t.Fail()
		}

		for idx, actor := range res {
			if actor.Id != test.expected[idx].Id {
				t.Fail()
			}
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

func TestDeleteActor(t *testing.T) {
	tests := []struct {
		id  int
		err bool
	}{
		{
			1,
			false,
		},
		{
			3,
			true,
		},
	}

	for _, test := range tests {
		_, err := as.DeleteActor(test.id)
		if err != nil {
			if test.err {
				return
			}

			t.Fail()
		}

		actor, err := as.GetActor(test.id)
		if err != nil && actor == nil {
			return
		}

		t.Fail()
	}
}
