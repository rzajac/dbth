package dbth

import (
    "database/sql"
    "fmt"

    _ "github.com/go-sql-driver/mysql"
)

// A mysql represents MySQL database tester.
type mysql struct {
    schema string
    *sql.DB
}

// newMySQL returns new MySQL tester.
func newMySQL(dialect, user, pass, host, schema string) Tester {
    dsn := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8", user, pass, host, schema)
    db, err := sql.Open(dialect, dsn)
    if err != nil {
        panic(err)
    }
    return &mysql{schema: schema, DB: db}
}

func (m *mysql) Tables() ([]string, error) {
    var ts []string
    rows, err := m.Query("SELECT `table_name` FROM `INFORMATION_SCHEMA`.`TABLES` WHERE `table_schema` = ?", m.schema)
    if err != nil {
        return ts, err
    }
    var tn string
    for rows.Next() {
        err := rows.Scan(&tn)
        if err != nil {
            return ts, err
        }
        ts = append(ts, tn)
    }
    return ts, nil
}

func (m *mysql) TablesMust() []string {
    ts, err := m.Tables()
    checkErrOrPanic(err)
    return ts
}

func (m *mysql) TableExists(name string) (bool, error) {
    ts, err := m.Tables()
    if err != nil {
        return false, err
    }
    for _, tn := range ts {
        if tn == name {
            return true, nil
        }
    }
    return false, nil
}

func (m *mysql) TableExistsMust(name string) bool {
    e, err := m.TableExists(name)
    checkErrOrPanic(err)
    return e
}

func (m *mysql) TableTruncate(name string) error {
    query := fmt.Sprintf("TRUNCATE TABLE `%s`", name)
    _, err := m.Exec(query)
    return err
}

func (m *mysql) TableDrop(name string) error {
    query := fmt.Sprintf("DROP TABLE `%s`", name)
    _, err := m.Exec(query)
    return err
}

func (m *mysql) TableDropAll() error {
    ts, err := m.Tables()
    if err != nil {
        return err
    }
    for _, tn := range ts {
        err := m.TableDrop(tn)
        if err != nil {
            return err
        }
    }
    return nil
}

func (m *mysql) RowExists(tableName string, pkName string, pkValue interface{}) (bool, error) {
    var count int
    query := fmt.Sprintf("SELECT COUNT(*) FROM `%s` WHERE `%s`=?", tableName, pkName)
    row := m.QueryRow(query, pkValue)
    err := row.Scan(&count)
    return count == 1, err
}

func (m *mysql) RowExistsMust(tableName, pkName string, pkValue interface{}) bool {
    e, err := m.RowExists(tableName, pkName, pkValue)
    checkErrOrPanic(err)
    return e
}

func (m *mysql) RowCount(tableName string) (int, error) {
    var count int
    query := fmt.Sprintf("SELECT COUNT(*) FROM `%s`", tableName)
    row := m.QueryRow(query)
    err := row.Scan(&count)
    return count, err
}

func (m *mysql) RowCountMust(tableName string) int {
    c, err := m.RowCount(tableName)
    checkErrOrPanic(err)
    return c
}

// checkErrOrPanic panics if error is not nil.
func checkErrOrPanic(err error) {
    if err != nil {
        panic(err)
    }
}
