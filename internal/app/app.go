package app

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/zanshin/atempo/internal/config"
	"github.com/zanshin/atempo/internal/database"
	l "github.com/zanshin/atempo/internal/logger"
	"github.com/zanshin/atempo/internal/model"
)

var (
	hrefClicks []model.HrefClick
	pageViews  []model.PageView
)

func RunSetup(conf config.Config) {
	// Connect to the database
	db := dbConnect(conf)
	l.Info.Println("Successfully connected to MySQL")

	// Create database if it doesn't already exist
	dbname := "atempo"
	dbResult := dbCreate(db, dbname)

	// Test result of creation for number of row affected.
	if dbConfirm(dbResult) {
		// Set connection pool options
		db.SetMaxOpenConns(20)
		db.SetMaxIdleConns(20)
		db.SetConnMaxLifetime(time.Minute * 5)
	} else {
		l.Error.Fatal("Database creation failed")
	}

	// Verify the connection
	if dbVerify(db) {
		l.Info.Printf("Connected to DB %s successfully\n", dbname)
	} else {
		l.Error.Fatal("Unable to ping database. Exiting")
	}
}

func Run(conf config.Config) {
	// Connect to the database
	db := dbConnect(conf)

	seconds := time.Duration(conf.BatchInsertSeconds) * time.Second

	l.Info.Println("About to listenForRecords")
	go listenForRecords(db, seconds)

	// Create the handlers for page-view/ and href-click/ POSTs
	l.Info.Println("How about some HandleFuncs")
	http.HandleFunc("/page-views/", makeHandler(pageViewsHandler))
	http.HandleFunc("/href-click/", makeHandler(hrefClickHandler))

	l.Info.Println("Listening and serving")
	http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), nil)
}

func listenForRecords(db *sql.DB, seconds time.Duration) {
	// Run every x seconds.
	for range time.Tick(seconds) {
		// Handle page views.
		newPageViews := make([]model.PageView, len(pageViews))
		copy(newPageViews, pageViews)
		go database.SetPageViews(db, newPageViews)
		pageViews = pageViews[0:0]

		// Handle href clicks.
		newHrefClicks := make([]model.HrefClick, len(hrefClicks))
		copy(newHrefClicks, hrefClicks)
		go database.SetHrefClicks(db, newHrefClicks)
		hrefClicks = hrefClicks[0:0]
	}
}

func IPAddress(remoteAddr string) string {
	arr := strings.Split(remoteAddr, ":")
	return arr[0]
}

func hrefClickHandler(w http.ResponseWriter, r *http.Request, body []byte) {
	hrefClick := model.HrefClick{}
	if err := json.Unmarshal(body, &hrefClick); err != nil {
		l.Error.Println("Unable to unmarshal hrefClick: ", err)
	}
	// Get ip address from http request
	hrefClick.IPAddress = IPAddress(r.RemoteAddr)
	hrefClicks = append(hrefClicks, hrefClick)
	w.WriteHeader(201)
}

func pageViewsHandler(w http.ResponseWriter, r *http.Request, body []byte) {
	pageView := model.PageView{}
	if err := json.Unmarshal(body, &pageView); err != nil {
		l.Error.Println("Unable to unmarshal pageView: ", err)
	}
	// Get ip address from http request
	pageView.IPAddress = IPAddress(r.RemoteAddr)
	pageViews = append(pageViews, pageView)
	w.WriteHeader(201)
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, []byte)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Headers", "x-requested-with, x-requested-by, Content-Type")
		w.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		if r.Method == "OPTIONS" {
			w.WriteHeader(200)
			return
		}
		if r.Method != "POST" {
			w.WriteHeader(405)
			return
		}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			l.Warning.Println("Unable to read requeset body: ", err)
		}
		fn(w, r, body)
	}
}

func dbConnect(conf config.Config) *sql.DB {
	db, err := database.DBConnect("", conf.DbConfig)
	if err != nil {
		l.Error.Fatal("Connection to MySQL failed. Exiting.")
	}

	return db
}

func dbCreate(db *sql.DB, dbname string) sql.Result {
	res, err := database.DBCreate(db, dbname)
	if err != nil {
		l.Error.Fatal("Database creation failed. Exiting")
	}

	return res

}

func dbConfirm(dbResult sql.Result) bool {
	no, err := dbResult.RowsAffected()
	if err != nil {
		l.Error.Printf("Error %s when fetching affected rows\n", err)
		l.Error.Fatal("Unable to get affected rows. Exiting")
		return false
	}

	l.Info.Printf("Rows affected: %d\n", no)
	return true

}

func dbVerify(db *sql.DB) bool {
	err := database.DBPing(db)
	if err != nil {
		l.Error.Fatal("Unable to ping database. Exiting")
		return false
	}

	return true

}
