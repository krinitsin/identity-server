package fixtures

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/sirupsen/logrus"
)

func SetUpSchema(t *testing.T, db *sql.DB, separator []byte, schemaFiles ...string) {
	for _, f := range schemaFiles {
		t.Log("applying file:", f)

		data, err := ioutil.ReadFile(f) //nolint
		if err != nil {
			t.Fatalf("read file error (%q): %s", f, err)
		}

		objects := bytes.Split(data, separator)

		for _, o := range objects {
			stmt := fmt.Sprintf("%s", o)
			_, err = db.ExecContext(context.Background(), stmt)
			if err != nil {
				t.Log("statement:", stmt)
				t.Fatal("object create error: ", err)
			}
		}
	}
}

func SetUpFixtures(db *sql.DB, fixturesFiles ...string) {

	fixtures, err := testfixtures.New(
		testfixtures.Database(db),            // You database connection
		testfixtures.Dialect("postgres"),     // Available: "postgresql", "timescaledb", "mysql", "mariadb", "sqlite" and "sqlserver"
		testfixtures.Files(fixturesFiles...), // the directory containing the YAML files
	)
	if err != nil {
		logrus.Fatal("creating fixtures ", err)
	}
	if err = fixtures.Load(); err != nil {
		logrus.Fatal("loading fixtures ", err)
	}
}
