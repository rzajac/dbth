package dbth

import (
    "fmt"
)

// Tester is database testing interface.
type Tester interface {
    // Tables returns a list of database table names in the schema.
    Tables() ([]string, error)

    // Tables returns a list of database table names in the schema.
    // Panics on error.
    TablesMust() []string

    // TableExists returns true if database table exists.
    TableExists(name string) (bool, error)

    // TableExists returns true if database table exists.
    // Panics on error.
    TableExistsMust(name string) bool

    // TableTruncate truncates the database table.
    TableTruncate(name string) error

    // TableDrop drops database table.
    TableDrop(name string) error

    // TableDropAll drop all database tables.
    TableDropAll() error

    // RowExists returns true if row with given primary key exists.
    RowExists(tableName, pkName string, pkValue interface{}) (bool, error)

    // RowExists returns true if row with given primary key exists.
    // Panics on error.
    RowExistsMust(tableName, pkName string, pkValue interface{}) bool

    // RowCount returns number of rows in the database table.
    RowCount(tableName string) (int, error)

    // RowCount returns number of rows in the database table.
    // Panics on error.
    RowCountMust(tableName string) int

    // Close closes database connection.
    Close() error
}

// NewDbTester returns new database tester for given dialect.
// Panics if the dialect is not supported.
func NewDbTester(dialect, user, pass, host, name string) Tester {
    switch dialect {
    case "mysql":
        return newMySQL(dialect, user, pass, host, name)
    default:
        panic(fmt.Sprintf("unknown database tester dialect: %s", dialect))
    }
}
