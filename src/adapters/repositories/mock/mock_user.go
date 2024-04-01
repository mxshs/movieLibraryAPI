package mock_db

import (
	"fmt"
	"mxshs/movieLibrary/src/domain"
)

func (mdb *MockDB) CreateUser(username, password string, role domain.Role) (*domain.User, error) {
	user := &domain.User{
		Id:       int(mdb.uid.Add(1)),
		Username: username,
		Password: password,
		Role:     role,
	}

	mdb.users[user.Id] = user

	return &domain.User{
		Id:       user.Id,
		Username: username,
		Password: user.Password,
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
			Password: user.Password,
			Role:     user.Role,
		}, nil
	}
}

func (mdb *MockDB) GetUsers() ([]*domain.User, error) {
	result := make([]*domain.User, len(mdb.users))

	idx := 0
	for _, user := range mdb.users {
		result[idx] = user
		idx++
	}

	return result, nil
}

func (mdb *MockDB) UpdateUser(id int, newUsername, newPassword string, newRole domain.Role) (*domain.User, error) {
	user, ok := mdb.users[id]
	if !ok {
		return nil, fmt.Errorf("entity with id %d does not exist", id)
	}

	user.Username = newUsername
	user.Password = newPassword
	user.Role = newRole

	return &domain.User{
		Id:       user.Id,
		Username: user.Username,
		Password: user.Password,
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

func (mdb *MockDB) LoginUser(username, password string) (*domain.User, error) {
	for _, user := range mdb.users {
		if user.Username == username {
			if user.Password != password {
				return nil, fmt.Errorf("invalid username/password combination")
			}

			return &domain.User{
				Id:       user.Id,
				Username: user.Username,
				Password: user.Password,
				Role:     user.Role,
			}, nil
		}
	}

	return nil, fmt.Errorf("invalid username/password combination")
}
