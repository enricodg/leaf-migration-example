package main

import (
	"github.com/enricodg/leaf-migration-example/migrations/mongo"
	"github.com/enricodg/leaf-migration-example/migrations/mysql"
	"github.com/enricodg/leaf-migration-example/migrations/postgre"
	migration "github.com/paulusrobin/leaf-utilities/leafMigration"
)

func main() {
	migration.New().
		WithMySql(mysql.InitializeMigrations).
		WithPostgre(postgre.InitializeMigrations).
		WithMongo(mongo.InitializeMigrations).
		Run()
}
