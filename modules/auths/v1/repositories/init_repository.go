package repositories

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type (
	authRepository struct {
		// ...database connection or other dependencies...
		db   *sqlx.DB
		stmt map[string]*sqlx.Stmt
	}

	statement struct {
		key   string
		query string
	}
)

const (
	schema = `auths`
)

func NewAuthRepository(db *sqlx.DB) AuthRepository {
	r := &authRepository{
		db: db,
	}
	r.initStmt()

	return r
}

func (r *authRepository) initStmt() {
	var (
		err     error
		stmtMap = make(map[string]*sqlx.Stmt)
		stmts   []statement
	)

	stmts = append(stmts, authStmts...)

	for _, v := range stmts {
		stmtMap[v.query], err = r.db.Preparex(v.query)
		if err != nil {
			log.Fatalf("Failed to initialize auth statement key %v, err : %v", v.key, err)
		}
	}

	r.stmt = stmtMap
}
