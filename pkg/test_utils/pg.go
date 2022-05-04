package testutils

import (
	"context"
	"identityserver/pkg/test_utils/containers"
	"identityserver/pkg/test_utils/fixtures"
	"log"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	"gorm.io/gorm"
)

type DBTest struct {
	t           *testing.T
	Conn        *gorm.DB
	container   testcontainers.Container
	fxtFiles    []string
	schemeFiles []string
}

// Option configures fixtures.
type Option func(*DBTest)

// WithFixtureFiles specify fixtures files.
func WithFixtureFiles(files ...string) Option {
	return func(opts *DBTest) {
		opts.fxtFiles = append(opts.fxtFiles, files...)
	}
}

// WithSchemaFiles specify schema files.
func WithSchemaFiles(files ...string) Option {
	return func(opts *DBTest) {
		opts.schemeFiles = append(opts.schemeFiles, files...)
	}
}

func NewDbTestSetUp(t *testing.T, opts ...Option) *DBTest {
	container, dsn, err := containers.CreatePGTestContainer(context.Background(), "test")
	if err != nil {
		t.Fatal(err)
	}
	log.Println("postgres container ready and running at dsn: ", dsn)
	db := &DBTest{
		t:         t,
		Conn:      dsn.Debug(),
		container: container,
	}

	for _, opt := range opts {
		opt(db)
	}

	sqlDB, _ := db.Conn.DB()

	fixtures.SetUpSchema(db.t, sqlDB, []byte(";"), db.schemeFiles...)
	fixtures.SetUpFixtures(sqlDB, db.fxtFiles...)

	return db
}

func (d DBTest) TearDown() {
	sqlDB, _ := d.Conn.DB()
	_ = sqlDB.Close()
	_ = d.container.Terminate(context.Background())
}
