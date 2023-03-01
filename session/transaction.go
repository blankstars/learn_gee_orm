package session

import "github.com/blankstars/gee_orm/log"

func (s *Session) Begin() error {
	var err error
	log.Info("transaction begin")
	if s.tx, err = s.db.Begin(); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (s *Session) Commit() error {
	var err error
	log.Info("transaction commit")
	if err = s.tx.Commit(); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (s *Session) Rollback() error {
	var err error
	log.Info("transaction rollback")
	if err = s.tx.Rollback(); err != nil {
		log.Error(err)
		return err
	}
	return nil
}
