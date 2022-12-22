package database

import (
	"backend-engineer-test-privy/config"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func InitDBMysql() *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.MYSQL_USER, config.MYSQL_PASSWORD, config.MYSQL_HOST, config.MYSQL_PORT, config.MYSQL_DATABASE)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	sqlMigration := `CREATE TABLE IF NOT EXISTS cakes (
		id          INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
		title       VARCHAR(255) NOT NULL,
		description TEXT NOT NULL,
		rating      FLOAT NOT NULL,
		image       VARCHAR(255) NOT NULL,
		created_at  DATETIME NOT NULL,
		updated_at  DATETIME NOT NULL,
		deleted_at  DATETIME
	  );`

	_, err = db.Exec(sqlMigration)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
