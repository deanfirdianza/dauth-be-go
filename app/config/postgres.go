package config

import (
	"log"
	"time"

	"github.com/deanfirdianza/dauth-be-go/app/env"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect(conf *env.Conf) (*sqlx.DB, error) {
	var (
		dsn       string
		connValue int
	)

	switch conf.App.Mode {
	case "production":
		dsn = "host=" + conf.DB_Prod.Host +
			" user=" + conf.DB_Prod.User +
			" password=" + conf.DB_Prod.Pass +
			" dbname=" + conf.DB_Prod.Name +
			" port=" + conf.DB_Prod.Port +
			" sslmode=disable TimeZone=Asia/Jakarta"
		log.Println(conf.App.Name, "runing on", conf.App.Mode, "mode")
		connValue = 800
	default:
		dsn = "host=" + conf.DB.Host +
			" user=" + conf.DB.User +
			" password=" + conf.DB.Pass +
			" dbname=" + conf.DB.Name +
			" port=" + conf.DB.Port +
			" sslmode=disable TimeZone=Asia/Jakarta"
		log.Println(conf.App.Name, "runing on", conf.App.Mode, "mode")
		connValue = 150
	}

	// SQLX setup
	sqlxDB, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to connect database using sqlx: %v", err)
	} else {
		log.Println("database connected using sqlx")
	}

	//set max connection pooling
	err = sqlxDB.Ping()
	if err != nil {
		log.Fatalf("Error while checking connection to database")
	}

	sqlxDB.SetMaxIdleConns(10)
	sqlxDB.SetMaxOpenConns(connValue)
	sqlxDB.SetConnMaxIdleTime(time.Minute)

	return sqlxDB, err

}
