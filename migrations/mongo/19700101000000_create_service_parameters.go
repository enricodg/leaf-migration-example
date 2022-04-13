package mongo

import (
	"context"
	nosqlConnection "github.com/paulusrobin/leaf-utilities/database/nosql/nosql"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type migration_19700101000000 struct {
	Log  leafLogger.Logger
	Conn nosqlConnection.Mongo
}

// NOTE: DO NOT CHANGE MIGRATION Version
func (m *migration_19700101000000) Version() uint64 {
	return uint64(19700101000000)
}

// NOTE: DO NOT CHANGE MIGRATION Name
func (m *migration_19700101000000) Name() string {
	return "create_service_parameters"
}

func (m *migration_19700101000000) Migrate() error {
	if err := m.Conn.DB().CreateCollection(context.Background(), "service_parameters"); err != nil {
		return err
	}

	if _, err := m.Conn.Indexes("service_parameters").CreateOne(context.Background(),
		mongo.IndexModel{
			Keys: bsonx.Doc{
				{
					Key:   "variable",
					Value: bsonx.Int32(1),
				},
				{
					Key:   "is_deleted",
					Value: bsonx.Int32(1),
				},
			},
			Options: options.Index().
				SetName("variable_1_is_deleted_1").
				SetBackground(true),
		}); err != nil {
		return err
	}

	return nil
}

func (m *migration_19700101000000) Rollback() error {
	if err := m.Conn.DB().Collection("service_parameters").Drop(context.Background()); err != nil {
		return err
	}

	return nil
}
