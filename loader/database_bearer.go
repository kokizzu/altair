package loader

import (
	"github.com/codefluence-x/altair/core"
	"github.com/codefluence-x/monorepo/db"
	"github.com/pkg/errors"
)

// ErrDatabasesIsNotExists thrown when database is not exists or not initialized yet in database bearer
var ErrDatabasesIsNotExists = errors.New("Database is not exists")

type databaseBearer struct {
	databases map[string]db.DB
	configs   map[string]core.DatabaseConfig
}

// DatabaseBearer handling on retrieval database instance
func DatabaseBearer(databases map[string]db.DB, configs map[string]core.DatabaseConfig) core.DatabaseBearer {
	return &databaseBearer{databases: databases, configs: configs}
}

func (d *databaseBearer) Database(dbName string) (db.DB, core.DatabaseConfig, error) {
	db, ok := d.databases[dbName]
	if ok == false {
		return nil, nil, ErrDatabasesIsNotExists
	}

	return db, d.configs[dbName], nil
}
