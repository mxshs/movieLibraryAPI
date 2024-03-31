package repository_adapter

import (
	"fmt"
	"mxshs/movieLibrary/src/domain"

	_ "github.com/lib/pq"
)

func (pdb *PgDB) CreateUser(username, password string, role domain.Role) (*domain.User, error) {
	q, err := pdb.db.Query(
		`INSERT INTO users(username, password, role)
        VALUES ($1, $2, $3)
        RETURNING uid;`,
		username,
		password,
		role,
	)
	if err != nil {
		return nil, err
	}

	if !q.Next() {
		return nil, fmt.Errorf(
			"unexpected empty return after succesful insertion")
	}

	var userId int

	err = q.Scan(&userId)
	if err != nil {
		return nil, err
	}

	return &domain.User{
		Id:       userId,
		Username: username,
		Password: password,
		Role:     role,
	}, nil
}

func (pdb *PgDB) GetUserByUsername(username string) (*domain.User, error) {
	q, err := pdb.db.Query(
		`SELECT * FROM users
        WHERE username = $1;`,
		username,
	)
	if err != nil {
		return nil, err
	}

	if !q.Next() {
		return nil, fmt.Errorf("user with username %s does not exist", username)
	}

	var user domain.User

	err = q.Scan(&user.Id, &user.Username, &user.Password, &user.Role)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (pdb *PgDB) GetUser(uid int) (*domain.User, error) {
	q, err := pdb.db.Query(
		`SELECT * FROM users
        WHERE uid = $1;`,
		uid,
	)
	if err != nil {
		return nil, err
	}

	if !q.Next() {
		return nil, fmt.Errorf("user with id %d does not exist", uid)
	}

	var user domain.User

	err = q.Scan(&user.Id, &user.Username, &user.Password, &user.Role)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (pdb *PgDB) GetUsers() ([]*domain.User, error) {
	q, err := pdb.db.Query(
		`SELECT * FROM users;`,
	)
	if err != nil {
		return nil, err
	}

	users := []*domain.User{}

	for q.Next() {
		var user domain.User

		err = q.Scan(&user.Id, &user.Username, &user.Password, &user.Role)
		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}

func (pdb *PgDB) UpdateUser(id int, newUsername, newPassword string, newRole domain.Role) (*domain.User, error) {
	q, err := pdb.db.Query(
		`UPDATE users SET
        username = COALESCE($2, username)
        password = COALESCE($3, password)
        role = COALESCE($4, role)
        WHERE id = $1
        RETURNING uid, username, password, role;`,
		id,
		newUsername,
		newPassword,
		newRole,
	)
	if err != nil {
		return nil, err
	}

	if !q.Next() {
		return nil, fmt.Errorf("user with id %d does not exist", id)
	}

	var user domain.User

	err = q.Scan(&user.Id, &user.Username, &user.Password, &user.Role)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (pdb *PgDB) DeleteUser(id int) error {
	_, err := pdb.db.Exec(
		`DELETE FROM users
        WHERE uid = $1;`,
		id,
	)

	return err
}

func (pdb *PgDB) LoginUser(username, password string) (*domain.User, error) {
	q, err := pdb.db.Query(
		`SELECT * FROM users
        WHERE username = $1;`,
		username,
	)
	if err != nil {
		return nil, err
	}

	if !q.Next() {
		return nil, fmt.Errorf("user with given credentials does not exist")
	}

	var user domain.UserDetail

	err = q.Scan(&user.Id, &user.Username, &user.Password, &user.Role)
	if err != nil {
		return nil, err
	}

	if user.Password != password {
		return nil, fmt.Errorf("user with given credentials does not exist")
	}

	return &domain.User{
		Id:       user.Id,
		Username: user.Username,
		Password: user.Password,
		Role:     user.Role,
	}, nil
}
