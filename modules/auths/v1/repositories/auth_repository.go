package repositories

import (
	"database/sql"
	"time"

	"github.com/deanfirdianza/dauth-be-go/modules/auths/v1/models"
)

// ...implement repository methods...

var (
	authStmts = []statement{
		{"SelectAuth", qAuthSelect},
		{"InsertAuth", qAuthInsert},
		{"RevokeAuth", qAuthRevoke},
	}
)

const (
	qAuthMainField = `id, account_id, refresh_token, created_at, expires_at, revoked`
	qAuthSelect    = `SELECT ` + qAuthMainField + ` FROM ` + schema + `.auth_tokens WHERE account_id = $1 and revoked = false ORDER BY created_at DESC LIMIT 1`
	qAuthInsert    = `INSERT INTO ` + schema + `.auth_tokens(account_id, refresh_token, expires_at, revoked) VALUES ($1, $2, $3, $4)`
	qAuthRevoke    = `UPDATE ` + schema + `.auth_tokens SET revoked = true WHERE account_id = $1`
	qAuthDeleteOld = `DELETE FROM ` + schema + `.auth_tokens WHERE account_id = $1 AND id NOT IN (
		SELECT id FROM ` + schema + `.auth_tokens WHERE account_id = $1 ORDER BY created_at DESC LIMIT 2
	)`
)

type AuthRepository interface {
	// ...define repository methods...
	InsertAuth(accountID string, refreshToken string, expiresAt time.Time, revoked bool) (sql.Result, error)
	SelectAuth(accountID string) (*models.AuthToken, error)
	RevokeAuth(accountID string) (sql.Result, error)
	DeleteOldAuths(accountID string) (sql.Result, error)
}

func (r *authRepository) DeleteOldAuths(accountID string) (sql.Result, error) {
	stmt, err := r.db.Prepare(qAuthDeleteOld)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(accountID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *authRepository) RevokeAuth(accountID string) (sql.Result, error) {
	stmt, err := r.db.Prepare(qAuthRevoke)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(accountID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *authRepository) InsertAuth(
	accountID string,
	refreshToken string,
	expiresAt time.Time,
	revoked bool,
) (sql.Result, error) {
	stmt, err := r.db.Prepare(qAuthInsert)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		accountID,
		refreshToken,
		expiresAt,
		revoked,
	)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *authRepository) SelectAuth(accountID string) (*models.AuthToken, error) {
	stmt, err := r.db.Prepare(qAuthSelect)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var authToken models.AuthToken
	err = stmt.QueryRow(accountID).Scan(
		&authToken.ID,
		&authToken.AccountID,
		&authToken.RefreshToken,
		&authToken.CreatedAt,
		&authToken.ExpiresAt,
		&authToken.Revoked,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &authToken, nil
}
