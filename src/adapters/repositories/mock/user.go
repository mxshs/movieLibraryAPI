package mock_db

import (
	"fmt"
	"mxshs/movieLibrary/src/domain"
)

func (mdb *MockDB) CreateUser(username, password string) (*domain.User, error) {
	user := &domain.UserDetail{
		Id:       int(mdb.uid.Add(1)),
		Username: username,
		Password: password,
		Role:     domain.USR,
	}

	mdb.users[user.Id] = user

	return &domain.User{
		Id:       user.Id,
		Username: username,
		Role:     user.Role,
	}, nil
}

func (mdb *MockDB) GetUser(id int) (*domain.User, error) {
	if user, ok := mdb.users[id]; !ok {
		return nil, fmt.Errorf("entity with id %d does not exist", id)
	} else {
		return &domain.User{
			Id:       user.Id,
			Username: user.Username,
			Role:     user.Role,
		}, nil
	}
}

func (mdb *MockDB) UpdateUser(id int, newUsername, newPassword string) (*domain.User, error) {
	user, ok := mdb.users[id]
	if !ok {
		return nil, fmt.Errorf("entity with id %d does not exist", id)
	}

	if len(newUsername) > 0 {
		user.Username = newUsername
	}

	if len(newPassword) > 0 {
		user.Password = newPassword
	}

	return &domain.User{
		Id:       user.Id,
		Username: user.Username,
		Role:     user.Role,
	}, nil
}

func (mdb *MockDB) DeleteUser(id int) error {
	if _, ok := mdb.users[id]; !ok {
		return fmt.Errorf("entity with id %d does not exist", id)
	}

	delete(mdb.users, id)

	return nil
}

func (mdb *MockDB) LoginUser(username, password string) error {
	for _, user := range mdb.users {
		if user.Username == username {
			if user.Password != password {
				return fmt.Errorf("invalid username/password combination")
			}

			return nil
		}
	}

	return fmt.Errorf("invalid username/password combination")
}
