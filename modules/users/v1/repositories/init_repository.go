package repositories

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type (
	userRepository struct {
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
	schema = `users`
)

func NewUserRepository(db *sqlx.DB) UserRepository {
	r := &userRepository{
		db: db,
	}
	r.initStmt()

	return r
}

func (r *userRepository) initStmt() {
	var (
		err     error
		stmtMap = make(map[string]*sqlx.Stmt)
		stmts   []statement
	)

	stmts = append(stmts, userStmts...)

	for _, v := range stmts {
		stmtMap[v.query], err = r.db.Preparex(v.query)
		if err != nil {
			log.Fatalf("Failed to initialize user statement key %v, err : %v", v.key, err)
		}
	}

	r.stmt = stmtMap
}
