package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/mjfagundez/login/login"
	"os"
)

func handleError(err error) {
	fmt.Fprintf(os.Stderr, "%v\n", err)
	os.Exit(1)
}

func main() {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		handleError(fmt.Errorf("unable to connect to database: %w", err))
	}
	defer conn.Close(context.Background())

	var name string
	var weight int64
	err = conn.QueryRow(context.Background(), "select name, weight from widgets where id=$1", 42).Scan(&name, &weight)
	if err != nil {
		handleError(fmt.Errorf("query row failed: %w", err))
	}

	fmt.Println(name, weight)
	login.LoginInterface()
}
