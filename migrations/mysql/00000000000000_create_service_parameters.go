package mysql

import (
	"context"
	sqlConnection "github.com/enricodg/leaf-utilities/database/sql/sql"
	"github.com/paulusrobin/leaf-utilities/leafMigration/helper/file"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
)

type migration_00000000000000 struct {
	Log  leafLogger.Logger
	Conn sqlConnection.ORM
}

// NOTE: DO NOT CHANGE MIGRATION Version
func (m *migration_00000000000000) Version() uint64 {
	return uint64(00000000000000)
}

// NOTE: DO NOT CHANGE MIGRATION Name
func (m *migration_00000000000000) Name() string {
	return "00000000000000_create_service_parameters"
}

func (m *migration_00000000000000) Migrate() error {

	script, err := file.ReadToString("./scripts/mysql/00000000000000_create_service_parameters_migrate.sql")
	if err != nil {
		return err
	}

	if err := m.Conn.Exec(context.Background(), script); err != nil {
		return err.Error()
	}

	return nil

}

func (m *migration_00000000000000) Rollback() error {
	script, err := file.ReadToString("./scripts/mysql/00000000000000_create_service_parameters_rollback.sql")
	if err != nil {
		return err
	}

	if err := m.Conn.Exec(context.Background(), script); err != nil {
		return err.Error()
	}

	return nil
}
