package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v1"
)

var dbmap *gorp.DbMap
var user, pwd, db string

func SetConnectionConfig(u, p, d string) {
	user = u
	pwd = p
	db = d
}

// Singleton for access to the DB config
func GetDb() *gorp.DbMap {
	if dbmap != nil {
		return dbmap
	}
	dbmap = NewDatabase()
	return dbmap
}

func NewDatabase() *gorp.DbMap {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(mariadb.mariadb.svc.cluster.local)/%s?parseTime=true", user, pwd, db))

	if err != nil {
		log.Fatal(err)
		return dbmap
	}

	// Needs to be tuned according the available infrastructure
	// Allowing 10 connection maximum
	// Having a pool of 3 ready to
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(3)
	db.SetConnMaxLifetime(5 * time.Minute)

	// construct a gorp DbMap
	dbmap = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	dbmap.TraceOn("[gorp]", log.New(os.Stdout, "jobs:", log.Lmicroseconds))

	// Register and create the tables
	// This should go outside the startup process. Not to mess with databases.
	if err := CreateScheme(); err != nil {
		log.Fatalf("Error creating the scheme and tables: %v", err)
		return dbmap
	}

	return dbmap
}

func Close() error {
	var err error
	if dbmap != nil {
		err = dbmap.Db.Close()
	}
	return err
}

func Clean() {
	DropTables()
	dbmap = nil
}
