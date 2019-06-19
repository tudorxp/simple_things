package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
	"time"
)

func die_if_error(err error) {
	if err != nil {
		panic(err)
	}
}

func handle_get_count(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	rows := 0
	err := db.QueryRow("select count(*) from kv").Scan(&rows)
	die_if_error(err)
	fmt.Fprintf(w, "rows: %d\n", rows)
}

func handle_count(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// check HTTP method
		switch r.Method {
		case "GET":
			handle_get_count(db, w, r)
		default:
			w.Header().Set("Allow", "GET")
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func handle_get_kv(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// TODO implement
}
func handle_post_kv(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// TODO implement
}
func handle_patch_kv(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// TODO implement
}

func handle_kv(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// check HTTP method
		switch r.Method {
		case "GET":
			handle_get_kv(db, w, r)
		case "POST":
			handle_post_kv(db, w, r)
		case "PATCH":
			handle_patch_kv(db, w, r)
		default:
			w.Header().Set("Allow", "GET, POST, PATCH")
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func main() {

	const (
		db_maxconn         = 10
		db_idleconn        = 2
		db_default_connstr = "postgres://test:test@localhost:5432/test?sslmode=disable"
		bind_address       = "localhost:8080"
	)

	// parse arguments
	db_connstr := flag.String("db", db_default_connstr, "PostgreSQL connection string")
	flag.Parse()

	// open DB connection
	db, err := sql.Open("postgres", *db_connstr)
	die_if_error(err)
	db.SetMaxOpenConns(db_maxconn)
	db.SetMaxIdleConns(db_idleconn)
	defer db.Close()

	// verify DB connection
	// NB as of 05.2019 lib/pq does not seem to implement timeouts on connect
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	err = db.PingContext(ctx)
	die_if_error(err)

	// start up an HTTP server
	mux := http.NewServeMux()

	mux.HandleFunc("/count", handle_count(db))
	mux.HandleFunc("/kv", handle_kv(db))

	err = http.ListenAndServe(bind_address, mux)
	die_if_error(err)

}
