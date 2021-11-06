package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"go-todo-app/config"
	"log"

	"github.com/google/uuid"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

var err error

const (
	tableNameUser = "users"
)

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}

	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(id INTEGER PRIMARY KEY AUTOINCREMENT,uuid STRING NOT NULL UNIQUE,name STRING,email STRING,password STRING,created_at DATETIME)`, tableNameUser)
	Db.Exec(cmdU)
}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj = uuid.New()
	return uuidobj
}

func Encrypt(platintext string) (crypttext string) {
	crypttext = fmt.Sprintf("%x", sha1.Sum([]byte(platintext)))
	return crypttext
}