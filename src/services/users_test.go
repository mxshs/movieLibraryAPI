package services_test

import (
	"mxshs/movieLibrary/src/domain"
	"testing"
)

func TestGetUser(t *testing.T) {
	tests := []struct {
		id       int
		err      bool
		expected *domain.User
	}{
		{
			1,
			false,
			&domain.User{
				Username: "test",
				Password: "testpassword",
				Role:     domain.USR,
			},
		},
		{
			2,
			true,
			nil,
		},
	}

	for _, test := range tests {
		res, err := us.GetUser(test.id)
		if err != nil {
			if test.err {
				return
			}

			t.Fail()
		}

		if res.Username != test.expected.Username {
			t.Fail()
		}

		if res.Password != test.expected.Password {
			t.Fail()
		}

		if res.Role != test.expected.Role {
			t.Fail()
		}
	}
}

func TestCreateUser(t *testing.T) {
	tests := []struct {
		username string
		password string
		role     domain.Role
		err      bool
		expected *domain.User
	}{
		{
			"test1",
			"testpassword1",
			domain.USR,
			false,
			&domain.User{
				Username: "test1",
				Password: "testpassword1",
				Role:     domain.USR,
			},
		},
		{
			"test2",
			"testpassword2",
			domain.ADM,
			false,
			&domain.User{
				Username: "test2",
				Password: "testpassword2",
				Role:     domain.ADM,
			},
		},
	}

	for _, test := range tests {
		res, err := us.CreateUser(test.username, test.password, test.role)
		if err != nil {
			if test.err {
				return
			}

			t.Fail()
		}

		user, err := us.GetUser(res.Id)
		if err != nil {
			panic(err)
		}

		if user.Username != test.expected.Username {
			t.Fail()
		}

		if user.Password != test.expected.Password {
			t.Fail()
		}

		if user.Role != test.expected.Role {
			t.Fail()
		}
	}
}

func TestGetUsers(t *testing.T) {
	tests := []struct {
		err      bool
		expected []*domain.User
	}{
		{
			false,
			[]*domain.User{
				{
					Username: "test",
					Password: "testpassword",
					Role:     domain.USR,
				},
				{
					Username: "test1",
					Password: "testpassword1",
					Role:     domain.USR,
				},
				{
					Username: "test2",
					Password: "testpassword2",
					Role:     domain.ADM,
				},
			},
		},
	}

	for _, test := range tests {
		res, err := us.GetUsers()
		if err != nil {
			if test.err {
				return
			}

			t.Fail()
		}

		for idx, user := range res {
			if user.Username != test.expected[idx].Username {
				t.Fail()
			}

			if user.Password != test.expected[idx].Password {
				t.Fail()
			}

			if user.Role != test.expected[idx].Role {
				t.Fail()
			}
		}
	}
}

func TestUpdateUser(t *testing.T) {
	tests := []struct {
		id          int
		newUsername string
		newPassword string
		newRole     domain.Role
		err         bool
		expected    *domain.User
	}{
		{
			2,
			"test11",
			"testpassword11",
			domain.ADM,
			false,
			&domain.User{
				Username: "test11",
				Password: "testpassword11",
				Role:     domain.ADM,
			},
		},
		{
			50,
			"",
			"",
			domain.ADM,
			true,
			nil,
		},
	}

	for _, test := range tests {
		res, err := us.UpdateUser(test.id, test.newUsername, test.newPassword, test.newRole)
		if err != nil {
			if test.err {
				return
			}

			t.Fail()
		}

		user, err := us.GetUser(res.Id)
		if err != nil {
			panic(err)
		}

		if user.Username != test.expected.Username {
			t.Fail()
		}

		if user.Password != test.expected.Password {
			t.Fail()
		}

		if user.Role != test.expected.Role {
			t.Fail()
		}
	}
}

func TestLoginUser(t *testing.T) {
	tests := []struct {
		username string
		password string
		err      bool
		expected *domain.User
	}{
		{
			"test",
			"testpassword",
			false,
			&domain.User{
				Id:       1,
				Username: "test",
				Password: "testpassword",
				Role:     domain.USR,
			},
		},
		{
			"itwillerror",
			"itwillerror",
			true,
			nil,
		},
	}

	for _, test := range tests {
		res, err := us.LoginUser(test.username, test.password)
		if err != nil {
			if test.err {
				return
			}

			t.Fail()
		}

		if res.Username != test.expected.Username {
			t.Fail()
		}

		if res.Password != test.expected.Password {
			t.Fail()
		}

		if res.Role != test.expected.Role {
			t.Fail()
		}
	}
}

func TestDeleteUser(t *testing.T) {
	tests := []struct {
		id       int
		err      bool
		expected []*domain.User
	}{
		{
			2,
			false,
			[]*domain.User{
				{
					Username: "test",
					Password: "testpassword",
					Role:     domain.USR,
				},
				{
					Username: "test2",
					Password: "testpassword2",
					Role:     domain.ADM,
				},
			},
		},
		{
			3,
			false,
			[]*domain.User{
				{
					Username: "test",
					Password: "testpassword",
					Role:     domain.USR,
				},
			},
		},
		{
			4,
			true,
			nil,
		},
	}

	for _, test := range tests {
		err := us.DeleteUser(test.id)
		if err != nil {
			if test.err {
				return
			}

			t.Fail()
		}

		res, err := us.GetUsers()
		if err != nil {
			panic(err)
		}

		for idx, user := range res {
			if user.Username != test.expected[idx].Username {
				t.Fail()
			}

			if user.Password != test.expected[idx].Password {
				t.Fail()
			}

			if user.Role != test.expected[idx].Role {
				t.Fail()
			}
		}
	}
}
