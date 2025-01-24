package repositories

// ...implement repository methods...

var (
	authStmts = []statement{
		// {"select", qAuthSelect},
	}
)

const (
	qAuthMainField = `username, password, email, salt, created_at, updated_at`
	qAuthSelect    = `SELECT ` + qAuthMainField + ` FROM ` + schema + ` WHERE username = ?`
)
