package repository

import (
	"database/sql"
	"errors"
	"jwt-api/internernal/models"
)

type User struct {
	DB *sql.DB
}

func NewUser(db *sql.DB) *User {
	return &User{DB: db}
}

func (r *User) CreateUser(user *models.User) error {
	query := "INSERT INTO users (name, email, Pass) VALUES ($1, $2, $3)"
	result, err := r.DB.Exec(query, user.Name, user.Email, user.Pass)
	if err != nil {
		return err
	}
	if rowsAffected, err := result.RowsAffected(); err != nil {
		return err
	} else if rowsAffected == 0 {
		return errors.New("no rows affected")
	}
	return nil
}
func (r *User) UpdateUser(user *models.User) error {
	query := "UPDATE users SET name = $1, email = $2, Pass = $3 WHERE id = $4"
	result, err := r.DB.Exec(query, user.Name, user.Email, user.Pass, user.ID)
	if err != nil {
		return err
	}
	if rowsAffected, err := result.RowsAffected(); err != nil {
		return err
	} else if rowsAffected == 0 {
		return errors.New("no rows affected")
	}
	return nil
}
func (r *User) DeleteUser(user *models.User) error {
	query := "DELETE FROM users WHERE id = $1"
	result, err := r.DB.Exec(query, user.ID)
	if err != nil {
		return err
	}
	if rowsAffected, err := result.RowsAffected(); err != nil {
		return err
	} else if rowsAffected == 0 {
		return errors.New("no rows affected")
	}
	return nil
}
func (r *User) GetUser(id uint64) (*models.User, error) {
	var user *models.User
	query := "SELECT * FROM users WHERE id = $1"
	row := r.DB.QueryRow(query, id)
	if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Pass); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *User) GetUserByEmail(email string) (*models.User, error) {
	var user *models.User
	query := "SELECT * FROM users WHERE email = $1"
	row := r.DB.QueryRow(query, email)
	if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Pass); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *User) GetAll() ([]models.User, error) {
	var users []models.User
	query := "SELECT * FROM users"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Pass); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
