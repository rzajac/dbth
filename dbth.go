package dbth

import "io"

// DB is database testing interface.
type DB interface {
    io.Closer
    // Tables returns a list of database table names in the schema.
    Tables() ([]string, error)
    // TableExists returns true if database table exists.
    TableExists(name string) (bool, error)
    // TableTruncate truncates the database table.
    TableTruncate(name string) error
    // TableDrop drops database table.
    TableDrop(name string) error
    // TableDropAll drop all database tables.
    TableDropAll() error
    // RowExists returns true if row with given primary key exists.
    RowExists(tableName, pkName string, pkValue interface{}) (bool, error)
    // RowCount returns number of rows in the database table.
    RowCount(tableName string) (int, error)
    // Must returns interface where methods panic if there is an error.
    Must() DBMust
}

// DBMust is database testing interface where methods panic on error.
type DBMust interface {
    io.Closer
    // Tables returns a list of database table names in the schema.
    // Panics on error.
    Tables() []string
    // TableExists returns true if database table exists.
    // Panics on error.
    TableExists(name string) bool
    // TableTruncate truncates the database table.
    TableTruncate(name string)
    // TableDrop drops database table.
    TableDrop(name string)
    // TableDropAll drop all database tables.
    TableDropAll()
    // RowExists returns true if row with given primary key exists.
    // Panics on error.
    RowExists(tableName, pkName string, pkValue interface{}) bool
    // RowCount returns number of rows in the database table.
    // Panics on error.
    RowCount(tableName string) int
}

// NewDbTester returns new database test helper for given dialect.
// Panics if the dialect is not supported.
func NewDbTester(dialect, user, pass, host, name string) DB {
    switch dialect {
    case "mysql":
        return newMySQL(dialect, user, pass, host, name)
    default:
        panic("unknown database tester dialect: " + dialect)
    }
}
