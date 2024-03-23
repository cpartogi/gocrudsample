package init

import (
	"fmt"
	"gocrudsample/lib/pkg/db"
	"os"

	"github.com/go-pg/pg"
	"github.com/spf13/viper"
)

func ConnectToPGServerRead() (*db.PgDB, error) {

	dbpg, err := db.CreatePGConnection(map[string]string{
		"host":     viper.GetString(`database.postgres.host`),
		"port":     viper.GetString(`database.postgres.port`),
		"user":     viper.GetString(`database.postgres.user`),
		"password": viper.GetString(`database.postgres.password`),
		"dbname":   viper.GetString(`database.postgres.dbname`),
		"sslmode":  viper.GetString(`database.postgres.sslmode`),
	})

	if err != nil {
		os.Exit(1)
	}

	return dbpg, err
}

func ConnectToPGServerWrite() (*pg.DB, error) {

	db := pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%s", viper.GetString(`database.postgres.host`), viper.GetString(`database.postgres.port`)),
		User:     viper.GetString(`database.postgres.user`),
		Password: viper.GetString(`database.postgres.password`),
		Database: viper.GetString(`database.postgres.dbname`),
	})

	_, err := db.Exec("SELECT 1")
	if err != nil {
		return nil, err
	}
	return db, err
}
