package repository_adapter

import (
	"fmt"
	"mxshs/movieLibrary/src/domain"
	pq "github.com/lib/pq"
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
		Role:     role,
	}, nil
}

func (pdb *PgDB) GetUserByUsername(username string) (*domain.User, error) {
	q, err := pdb.db.Query(
		`SELECT uid, username, role FROM users
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

	err = q.Scan(&user.Id, &user.Username, &user.Role)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (pdb *PgDB) GetUserById(uid int) (*domain.User, error) {
	q, err := pdb.db.Query(
		`SELECT uid, username, role FROM users
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

	err = q.Scan(&user.Id, &user.Username, &user.Role)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (pdb *PgDB) GetUsersById(uids ...int) ([]*domain.User, error) {
	q, err := pdb.db.Query(
		`SELECT uid, username, role FROM movies
        WHERE mid IN $1;`,
		pq.Array(uids),
	)
	if err != nil {
		return nil, err
	}

	users := make([]*domain.User, len(uids))

	for id := range uids {
		if !q.Next() {
			break
		}

		var user domain.User

		err = q.Scan(&user.Id, &user.Username, &user.Role)
		if err != nil {
			return nil, err
		}

		users[id] = &user
	}

	return users, nil
}

func (pdb *PgDB) GetUsers() ([]*domain.User, error) {
	q, err := pdb.db.Query(
		`SELECT uid, username, role FROM users;`,
	)
	if err != nil {
		return nil, err
	}

	users := []*domain.User{}

	for q.Next() {
		var user domain.User

		err = q.Scan(&user.Id, &user.Username, &user.Role)
		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}

func (pdb *PgDB) UpdateUser(username, newPassword string) (*domain.User, error) {
	q, err := pdb.db.Query(
		`UPDATE users SET
        password = COALESCE($2, password)
        WHERE username = $1
        RETURNING uid, username, role;`,
		username,
		newPassword,
	)
	if err != nil {
		return nil, err
	}

	if !q.Next() {
		return nil, fmt.Errorf("user with username %s does not exist", username)
	}

	var user domain.User

	err = q.Scan(&user.Id, &user.Username, &user.Role)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (pdb *PgDB) DeleteUser(username string) error {
	_, err := pdb.db.Exec(
		`DELETE FROM users
        WHERE username = $1;`,
		username,
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
		return nil, fmt.Errorf("user with username %s does not exist", username)
	}

	var user domain.User

	err = q.Scan(&user.Id, &user.Username, &user.Password, &user.Role)
	if err != nil {
		return nil, err
	}

	if user.Password != password {
		return nil, fmt.Errorf("invalid password")
	}

	return &user, nil
}

