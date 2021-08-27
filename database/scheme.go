package database

import (
	"github.com/xeviknal/background-commons/models"
)

type tables []table
type table struct {
	Name    string
	Model   interface{}
	Key     string
	Columns columns
}

type columns []column
type column struct {
	Name    string
	NotNull bool
}

func CreateScheme() error {
	for _, table := range getTables() {
		t := dbmap.AddTableWithName(table.Model, table.Name).SetKeys(true, table.Key)
		for _, column := range table.Columns {
			t.ColMap(column.Name).SetNotNull(column.NotNull)
		}
	}
	// TODO: Implement migration tool (rubenv/sql-migrate)
	return dbmap.CreateTablesIfNotExists()
}

func DropTables() error {
	return dbmap.DropTablesIfExists()
}

func getTables() tables {
	return tables{
		{
			"jobs",
			models.Job{},
			"Id",
			columns{
				{"Id", true},
				{"ObjectId", true},
				{"CreatedAt", true},
				{"Sleep", false},
				{"FinishedAt", false},
				{"StartedAt", false},
			},
		},
	}
}
