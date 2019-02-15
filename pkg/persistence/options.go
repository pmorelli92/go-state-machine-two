package persistence

import (
	"fmt"
	"os"
	"strconv"
)

type PostgresOptions struct {
	host		string
	port		int
	database	string
	user		string
	password	string
}

func NewPostgresOptions() PostgresOptions {

	port, err := strconv.Atoi(os.Getenv("POSTGRES_SERVICE_PORT"))
	panicWhenError(err)

	return PostgresOptions{
		host: os.Getenv("POSTGRES_SERVICE_HOST"),
		port: port,
		database: os.Getenv("POSTGRES_DATABASE"),
		user: os.Getenv("POSTGRES_USER"),
		password: os.Getenv("POSTGRES_PASSWORD"),
	}
}

func (opt *PostgresOptions) getConnection() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		opt.host, opt.port, opt.user, opt.password, opt.database)
}

func panicWhenError(err error) {
	if err != nil {
		panic(err)
	}
}