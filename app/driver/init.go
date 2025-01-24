package driver

import (
	"github.com/deanfirdianza/dauth-be-go/app/config"
	"github.com/deanfirdianza/dauth-be-go/app/env"
)

var (
	Conf, ErrConf = env.Init()

	DBSqlx, ErrDBSqlx = config.Connect(&Conf)
)
