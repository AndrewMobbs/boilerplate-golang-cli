package app

import (
	"database/sql"

	"github.com/AndrewMobbs/appdb"
)

type appDB struct {
	db      *sql.DB
	Path    string
	AppName string
}

const schemaVersion uint8 = 1

func schema() []string {
	return []string{
		`CREATE TABLE id(
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		`,
	}
}

func NewAppDB(path string, appName string) *appDB {
	return &appDB{
		Path:    path,
		AppName: appName,
	}
}

func (s *appDB) Open() error {
	var err error
	if s.db == nil {
		s.db, err = appdb.OpenAppDB(s.Path, s.AppName, schemaVersion)
	}
	return err
}

func (s *appDB) Initialize() error {
	db, err := appdb.InitAppDB(s.Path, s.AppName, schemaVersion, schema())
	if err == nil {
		db.Close()
	}
	return err
}

func (s *appDB) Close() error {
	if s.db != nil {
		return s.db.Close()
	}
	return nil
}