package main

import (
	"github.com/palestine-nights/auth/src/api"
	"github.com/palestine-nights/auth/src/tools"
)

func main() {
	databaseUser := tools.GetEnv("DATABASE_USER", "root")
	databasePassword := tools.GetEnv("DATABASE_PASSWORD", "")
	databaseName := tools.GetEnv("DATABASE_NAME", "restaurant")
	databaseHost := tools.GetEnv("DATABASE_HOST", "localhost")
	databasePort := tools.GetEnv("DATABASE_PORT", "3306")

	server := api.GetServer(
		databaseUser,
		databasePassword,
		databaseName,
		databaseHost,
		databasePort,
	)

	server.Router.Run()
}
