package model

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/zanshin/atempo/pkg/config"
	l "github.com/zanshin/atempo/pkg/logger"
)

var (
	ctx        context.Context
	cancelfunc context.CancelFunc
)

// Build the Data Source Name (DSN)
func dsn(dbname string, dbc config.DbConfig) string {
	l.Info.Printf("Data source name: %s:%s@tcp(%s:%d)/%s\n", dbc.User, "********", dbc.Host, dbc.Port, dbname)
	if len(dbname) == 0 {
		l.Warning.Println("No database name provided, only connecting to MySQL server.")
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbc.User, dbc.Pass, dbc.Host, dbc.Port, dbname)
}

// Connect to a database, if empty string is passed, jsut connect to MySQL
func DBConnect(dbname string, dbc config.DbConfig) (*sql.DB, error) {
	l.Info.Println("Entering DBConnect")
	db, err := sql.Open("mysql", dsn(dbname, dbc))
	if err != nil {
		l.Error.Printf("Connection to MySQL failed. Reason: %s\n", err)
		return nil, err
	}

	// defer db.Close()

	return db, nil
}

// Create database, unless it already exists
func DBCreate(db *sql.DB, dbname string) (sql.Result, error) {
	l.Info.Println("Entering DBCreate")
	ctx, cancelfunc = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	res, err := db.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+dbname)
	if err != nil {
		l.Error.Printf("Creating database %q failed. Reason: %s\n", dbname, err)
		return nil, err
	}

	l.Info.Printf("Database connection established to %q", dbname)
	return res, nil
}

// Ping the database to verify the connection
func DBPing(db *sql.DB) error {
	l.Info.Println("Entering DBPing")
	ctx, cancelfunc = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	err := db.PingContext(ctx)
	if err != nil {
		l.Error.Printf("Unable to ping database. Reason: %s\n", err)
		return err
	}

	return nil
}

func Db(dbConfig config.DbConfig) *sql.DB {
	fmt.Println("reached func DB in persist")
	db, err := sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
			dbConfig.User, dbConfig.Pass, dbConfig.Host, dbConfig.Port, dbConfig.Name))

	fmt.Println("sql.Open completed in persist")
	defer db.Close()

	// fmt.Sprintf("user=%s password=%s host=%s dbname=%s", dbConfig.User,
	// 	dbConfig.Pass, dbConfig.Host, dbConfig.Name))
	if err != nil {
		log.Fatal("Unable to connect to the database: ", err)
	}
	return db
}

func createTable(db *sql.DB, name string, query string) {

}

func SetPageViews(db *sql.DB, pageViews []PageView) {
	fmt.Println("reached func SetPageViews in persist")
	if len(pageViews) < 1 {
		return
	}
	tx, _ := db.Begin()
	stmt, err := db.Prepare("INSERT INTO page_view(timestamp, url, ip_address, user_agent, screen_height, screen_width) VALUES (NOW(), $1, $2, $3, $4, $5)")
	if err != nil {
		log.Println("Unable to prepare statment for PageView: ", err)
	}
	for k := range pageViews {
		_, err = tx.Stmt(stmt).Exec(
			pageViews[k].URL,
			pageViews[k].IPAddress,
			pageViews[k].UserAgent,
			pageViews[k].ScreenHeight,
			pageViews[k].ScreenWidth,
		)
		if err != nil {
			log.Fatal(err)
		}
	}
	tx.Commit()
}

func SetHrefClicks(db *sql.DB, hrefClicks []HrefClick) {
	fmt.Println("reached func SetHrefClicks in persist")
	if len(hrefClicks) < 1 {
		return
	}
	tx, _ := db.Begin()
	stmt, err := db.Prepare("INSERT INTO href_click(timestamp, url, ip_address, href, href_rectangle) VALUES (NOW(), $1, $2, $3, box(point($4,$5), point($6,$7)))")
	if err != nil {
		log.Println("Unable to prepare statment for HrefClick: ", err)
	}
	for k := range hrefClicks {
		tx.Stmt(stmt).Exec(
			hrefClicks[k].URL,
			hrefClicks[k].IPAddress,
			hrefClicks[k].Href,
			hrefClicks[k].HrefTop,
			hrefClicks[k].HrefRight,
			hrefClicks[k].HrefBottom,
			hrefClicks[k].HrefLeft,
		)
	}
	tx.Commit()
}
