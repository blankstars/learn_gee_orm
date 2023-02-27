package session

import (
	"database/sql"
	"github.com/blankstars/gee_orm/dialect"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"testing"
)

type User struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

func TestSession_CreateTable(t *testing.T) {
	db, err := sql.Open("sqlite3", "gee.db")
	if err != nil {
		log.Panic(err)
	}
	dial, _ := dialect.GetDialect("sqlite3")
	s := New(db, dial).Model(&User{})

	s.DropTable()
	s.CreateTable()
	if !s.HasTable() {
		t.Fatal("Failed to create table User")
	}

}
