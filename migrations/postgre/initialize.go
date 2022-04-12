// GENERATED CODE, DO NOT EDIT THIS FILE
package postgre

import (
	sqlConnection "github.com/enricodg/leaf-utilities/database/sql/sql"
	"github.com/paulusrobin/leaf-utilities/leafMigration/helper/migration"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
)

func InitializeMigrations(conn sqlConnection.ORM, log leafLogger.Logger) []migration.Migration {
	var migrations = make([]migration.Migration, 0)
	migrations = append(migrations, Migration_00000000000000(conn, log))
	return migrations
}

func Migration_00000000000000(conn sqlConnection.ORM, log leafLogger.Logger) migration.Migration {
	return &migration_00000000000000{
		Log:  log,
		Conn: conn,
	}
}