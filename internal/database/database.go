package database

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type DB struct {
	dbConf string
	DB     *pgxpool.Pool
}

func NewDatabase(cfg string, db *pgxpool.Pool) *DB {
	return &DB{
		dbConf: cfg,
		DB:     db,
	}
}

func Init(dbConf string) *DB {

	db, err := pgxpool.New(context.Background(), dbConf)
	if err != nil {
		log.Fatal(err)

	}
	database := NewDatabase(dbConf, db)

	return database
}

// func (d *DB) Open() error {
// 	db, err := pgxpool.New(context.Background(), d.dbConf)
// 	if err != nil {
// 		log.Fatal(err)
//
// 	}
// 	if err = d.Ping(); err != nil {
// 		return err
// 	}
// 	d.DB = db
// 	return nil
// }

func (d *DB) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := d.DB.Ping(ctx); err != nil {
		return err
	}
	return nil
}

//
// func (d *DB) Close() {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	if d.DB != nil {
// 		if err := d.DB.Close(); err != nil {
// 			log.Println("Error to close database")
// 		}
// 	}
// }
