package storage

import (
	"database/sql"

	"github.com/SteMak/vanilla/config"
	"github.com/SteMak/vanilla/out"

	// sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

var (
	Users usersTable
)

func Init() {
	out.Info("\nstorage connection         ")

	var err error
	db, err = sql.Open(config.Storage.Driver, config.Storage.Connection)
	if err != nil {
		out.Infoln("[FAIL]")
		out.Fatal(err)
	}

	db.SetMaxOpenConns(1)
	out.Infoln("[OK]")
}
