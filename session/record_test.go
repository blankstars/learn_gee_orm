package session

import (
	"database/sql"
	"github.com/blankstars/gee_orm/dialect"
	"go.uber.org/multierr"
	"log"
	"testing"
)

var (
	user1 = &User{"Tom", 18}
	user2 = &User{"Sam", 25}
	user3 = &User{"Jack", 25}
)

func testRecordInit(t *testing.T) *Session {
	t.Helper()
	db, err := sql.Open("sqlite3", "gee.db")
	if err != nil {
		log.Panic(err)
	}
	dial, _ := dialect.GetDialect("sqlite3")
	s := New(db, dial).Model(&User{})

	err = multierr.Combine(
		s.DropTable(),
		s.CreateTable(),
		func() error {
			_, err := s.Insert(user1, user2)
			return err
		}(),
	)
	if err != nil {
		t.Fatal("failed init test records")
	}

	return s
}

func TestSession_Insert(t *testing.T) {
	s := testRecordInit(t)
	affected, err := s.Insert(user3)
	if err != nil || affected != 1 {
		t.Fatal("failed to create record")
	}
}

func TestSession_Find(t *testing.T) {
	s := testRecordInit(t)
	var users []User
	err := s.Find(&users)
	if err != nil || len(users) != 2 {
		t.Fatal("failed to query all")
	}
}
