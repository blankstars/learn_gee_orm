package session

import (
	"database/sql"
	"github.com/blankstars/gee_orm/clause"
	"github.com/blankstars/gee_orm/dialect"
	"github.com/blankstars/gee_orm/log"
	"github.com/blankstars/gee_orm/schema"
	"strings"
)

type Session struct {
	db       *sql.DB
	dialect  dialect.Dialect
	tx       *sql.Tx
	relTable *schema.Schema
	clause   clause.Clause
	sql      strings.Builder
	sqlVals  []any
}

type CommonDB interface {
	Query(query string, args ...any) (*sql.Rows, error)
	QueryRow(query string, args ...any) *sql.Row
	Exec(query string, args ...any) (sql.Result, error)
}

var _ CommonDB = (*sql.DB)(nil)
var _ CommonDB = (*sql.Tx)(nil)

func New(db *sql.DB, dialect dialect.Dialect) *Session {
	return &Session{
		db:      db,
		dialect: dialect,
	}
}

func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlVals = nil
	s.clause = clause.Clause{}
}

func (s *Session) DB() CommonDB {
	if s.tx != nil {
		return s.tx
	}
	return s.db
}

func (s *Session) Raw(sql string, values ...any) *Session {
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	s.sqlVals = append(s.sqlVals, values...)
	return s
}

func (s *Session) Exec() (result sql.Result, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVals)
	if result, err = s.db.Exec(s.sql.String(), s.sqlVals...); err != nil {
		log.Error(err)
	}
	return
}

func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVals)
	return s.db.QueryRow(s.sql.String(), s.sqlVals...)
}

func (s *Session) QueryRows() (rows *sql.Rows, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVals)
	if rows, err = s.db.Query(s.sql.String(), s.sqlVals...); err != nil {
		log.Error(err)
	}
	return
}
