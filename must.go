package dbth

type must struct {
    db DB
}

func (m *must) Tables() []string {
    ret, err := m.db.Tables()
    checkErrOrPanic(err)
    return ret
}

func (m *must) TableExists(name string) bool {
    ret, err := m.db.TableExists(name)
    checkErrOrPanic(err)
    return ret
}

func (m *must) TableTruncate(name string) {
    err := m.db.TableTruncate(name)
    checkErrOrPanic(err)
}

func (m *must) TableDrop(name string) {
    err := m.db.TableDrop(name)
    checkErrOrPanic(err)
}

func (m *must) TableDropAll() {
    err := m.db.TableDropAll()
    checkErrOrPanic(err)
}

func (m *must) RowExists(tableName, pkName string, pkValue interface{}) bool {
    ret, err := m.db.RowExists(tableName, pkName, pkValue)
    checkErrOrPanic(err)
    return ret
}

func (m *must) RowCount(tableName string) int {
    ret, err := m.db.RowCount(tableName)
    checkErrOrPanic(err)
    return ret
}

func (m *must) Close() error {
    return m.db.Close()
}

// checkErrOrPanic panics if error is not nil.
func checkErrOrPanic(err error) {
    if err != nil {
        panic(err)
    }
}
