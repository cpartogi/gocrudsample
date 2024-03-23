package db

import (
	"context"
	"database/sql"
	"fmt"
	"gocrudsample/lib/pkg/utils"
	"log"
	"time"

	_ "github.com/lib/pq" // postgres driver
)

type PgDB struct {
	DB   *sql.DB
	Conn *sql.Conn
}

func CreatePGConnection(opts map[string]string) (*PgDB, error) {
	var err error

	port := utils.GenerateInt(opts["port"])
	if port == 0 {
		log.Println("Invalid port number : ", opts["port"])
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=%s",
		opts["host"], port, opts["user"], opts["password"], opts["dbname"], opts["sslmode"])

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println("DB Error")
	}

	// assuming max open connection pool for 25, and need allocate open + idle
	db.SetMaxOpenConns(25)

	// assuming will set max idle 5 per 10 minutes
	db.SetConnMaxIdleTime(10 * time.Minute)

	// set db connection pooling slot 25 each con
	// expired connection each hour
	db.SetConnMaxLifetime(time.Hour)

	ctx := context.Background()

	// check conn
	conn, err := db.Conn(ctx)
	if err != nil {
		log.Println("Conn Error")
	}

	log.Println("Connected to PG DB Server: ", opts["host"], " at port:", opts["port"], " successfully!")

	return &PgDB{DB: db, Conn: conn}, nil
}
