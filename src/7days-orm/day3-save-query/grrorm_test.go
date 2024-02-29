package day3_save_query

import (
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func OpenDb(t *testing.T) *Engine {
	t.Helper()
	engine, err := NewEngine("sqlite3", "gee.db")
	if err != nil {
		t.Fatal("failed to connect", err)
	}
	return engine
}

func TestNewEngine(t *testing.T) {
	engine := OpenDb(t)
	defer engine.Close()
}
