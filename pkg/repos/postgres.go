package repos

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PgConfig struct {
	UserName string `long:"db.username" description:"Db connection username" required:"yes"`
	Password string `long:"db.password" description:"Db connection password" required:"yes"`
	Dbname   string `long:"db.dbname" description:"Database name" required:"yes"`
	Port     string `long:"db.port" description:"Db connection port" required:"yes"`
	Host     string `long:"db.host" description:"Db connection host" required:"yes"`
}

// GetDBInstance return connection to Postgres
func GetDBInstance(pgConfig PgConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%s",
		pgConfig.Host,
		pgConfig.UserName,
		pgConfig.Dbname,
		pgConfig.Password,
		pgConfig.Port,
	)

	db, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn, PreferSimpleProtocol: true}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
