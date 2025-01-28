package repositories

import (
	"database/sql"
	"time"

	"github.com/deanfirdianza/dauth-be-go/modules/users/v1/models"
)

// ...implement repository methods...

var (
	userStmts = []statement{
		{"qUserSelect", qUserSelect},
		{"qUserInsert", qUserInsert},
	}
)

const (
	qUserMainField = `username, password, email, salt, created_at, updated_at`
	qUserSelect    = `SELECT ` + qUserMainField + ` FROM ` + schema + `.accounts WHERE username = $1`
	qUserInsert    = `INSERT INTO ` + schema + `.accounts(username, password, email, salt, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`
)

type UserRepository interface {
	InsertUser(username string, password string, email string, salt string) (sql.Result, error)
	SelectUser(username string) (*models.Accounts, error)
}

func (r *userRepository) InsertUser(
	username string,
	password string,
	email string,
	salt string,
) (sql.Result, error) {
	stmt, err := r.db.Prepare(qUserInsert)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		username,
		password,
		email,
		salt,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *userRepository) SelectUser(username string) (*models.Accounts, error) {
	stmt, err := r.db.Prepare(qUserSelect)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var user models.Accounts
	err = stmt.QueryRow(username).Scan(
		&user.Username,
		&user.Password,
		&user.Email,
		&user.Salt,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
