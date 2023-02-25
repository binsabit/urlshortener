package api

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/binsabit/urlshortener/internal/repository/postgres"
	_ "github.com/lib/pq"
)

type config struct {
	port int
	db   struct {
		port         int
		host         string
		name         string
		password     string
		user         string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
}

const version = "1.0.0"

//we need this to maintain dependencies
type application struct {
	config
	models postgres.Models
	//logger
}

func StartAndConfigure() {
	cfg := configure()

	db, err := openDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	app := application{
		config: cfg,
		models: postgres.NewModels(db),
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port),
		Handler:      app.routes(),
		IdleTimeout:  1 * time.Minute,
		WriteTimeout: 10 * time.Minute,
		ReadTimeout:  10 * time.Minute,
	}
	log.Println("server has started")

	log.Fatal(srv.ListenAndServe())

}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.db.host, cfg.db.port, cfg.db.user, cfg.db.password, cfg.db.name))
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.db.maxOpenConns)

	db.SetMaxIdleConns(cfg.db.maxIdleConns)

	duration, err := time.ParseDuration(cfg.db.maxIdleTime)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func configure() config {
	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "indicate the port that server will be running on")
	flag.Parse()
	cfg.db = struct {
		port         int
		host         string
		name         string
		password     string
		user         string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}{
		port:         5432,
		host:         "localhost",
		name:         "urlshortener",
		password:     "admin",
		user:         "admin",
		maxOpenConns: 25,
		maxIdleConns: 25,
		maxIdleTime:  "15m",
	}
	return cfg
}
