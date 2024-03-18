package schema

import (
	"day7_migrate/dialect"
	"log"
	"testing"
)

type User struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

var TestDial, _ = dialect.GetDialect("sqlite3")

func TestParse(t *testing.T) {
	schema := Parse(&User{}, TestDial)

	if schema.Name != "User" || len(schema.Fields) != 2 {
		t.Fatal("failed to parse User struct")
	}

	if schema.GetField("Name").Tag != "PRIMARY KEY" {
		log.Fatal("failed to parse primary key")
	}
}
