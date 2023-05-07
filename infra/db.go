package infra

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	// blank import for MySQL driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/nokin-all-of-career/career-web-backend/configs"
)

// DB : DB
var DB *sql.DB

// NewDBConnection : connect db
func NewDBConnection() error {
	/* ===== connect datebase ===== */
	// user
	user := os.Getenv("MYSQL_USER")
	// password
	password := os.Getenv("MYSQL_PASSWORD")
	// connection database
	database := os.Getenv("MYSQL_DATABASE")
	// connection instancename
	instanceConnectionName := os.Getenv("INSTANCE_CONNECTION_NAME")

	// designate the connection destination information as follows
	var dbURI string
	dbURI = fmt.Sprintf("%s:%s@unix(/cloudsql/%s)/%s?parseTime=true", user, password, instanceConnectionName, database)

	// DB is the pool of database connections.
	var err error
	DB, err = setupDB("mysql", dbURI)
	if err != nil {
		return fmt.Errorf("sql.Open: %v", err)
	}

	return err
}

// NewLocalDBConnection : connect local db
func NewLocalDBConnection() error {
	/* ===== connect datebase ===== */
	user := configs.Config.User
	if user == "" {
		user = os.Getenv("MYSQL_USER")
	}
	password := configs.Config.Password
	if password == "" {
		password = os.Getenv("MYSQL_PASSWORD")
	}
	host := configs.Config.Host
	if host == "" {
		host = os.Getenv("MYSQL_HOST")
	}
	port := configs.Config.Port
	if port == "" {
		port = os.Getenv("MYSQL_PORT")
	}
	database := configs.Config.Database
	if database == "" {
		database = os.Getenv("MYSQL_DATABASE")
	}

	// designate the connection destination information as follows
	// user:password@tcp(host:port)/database
	// TODO:ORMでコネクションプールしておく
	var err error
	DB, err = setupDB("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, database))
	if err != nil {
		return fmt.Errorf("sql.Open: %v", err)
	}

	return err
}

// this function is a function for connection pooling
func setupDB(dbDriver string, dsn string) (*sql.DB, error) {
	db, err := sql.Open(dbDriver, dsn)
	if err != nil {
		return nil, err
	}
	// コネクションプールの最大接続数を設定。
	db.SetMaxIdleConns(100)
	// 接続の最大数を設定。 nに0以下の値を設定で、接続数は無制限。
	db.SetMaxOpenConns(100)
	// 接続の再利用が可能な時間を設定。dに0以下の値を設定で、ずっと再利用可能。
	db.SetConnMaxLifetime(100 * time.Second)

	return db, err
}
