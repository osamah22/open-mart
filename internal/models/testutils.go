package models

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func NewDatabase() (*sql.DB, error) {
	dsn := viper.GetString("DB_URL")
	return sql.Open("postgres", dsn)
}
