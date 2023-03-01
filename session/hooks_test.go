package session

import (
	"fmt"
	"github.com/blankstars/gee_orm/log"
	"testing"
)

func (u *User) BeforeQuery(s *Session) error {
	log.Info("before query")
	return nil
}

func (u *User) AfterQuery(s *Session) error {
	log.Info("after query")
	u.Name = fmt.Sprintf("%c*%c", u.Name[0], u.Name[2])
	return nil
}

func TestSession_CallMethod(t *testing.T) {
	s := testRecordInit(t)
	u := &User{}
	err := s.First(u)
	if err != nil || u.Name[1] != '*' {
		t.Fatal("Failed to call method")
	}
}
