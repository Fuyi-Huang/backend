package library

import (
	"database/sql"

	"github.com/fuyi-huang/backend/config"
	extend "github.com/fuyi-huang/backend/ext"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectMysql() uint {
	db_str := config.Db_master["user"] + ":" + config.Db_master["passwd"] + "@tcp(" + config.Db_master["host"] + ")/" + config.Db_master["db"]
	var err error
	DB, err = sql.Open("mysql", db_str)
	if err != nil {
		return extend.MYSQL_OP_ERR
	}

	if err = DB.Ping(); err != nil {
		return extend.MYSQL_OP_ERR
	}

	return extend.SUCCESS
}
