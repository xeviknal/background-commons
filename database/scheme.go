package database

import (
	"fmt"
	"log"

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
	if dbmap == nil {
		err := fmt.Errorf("can't create schemes because db is not present")
		log.Fatal(err)
		return err
	}

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
	var err error
	if dbmap != nil {
		err = dbmap.DropTablesIfExists()
	}
	return err
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
			},
		},
	}
}
