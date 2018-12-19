package mysqls

import (
	_ "github.com/go-sql-driver/mysql"
	dbmysql "github.com/jinzhu/gorm"
)

// Store holds db handle.
type Store struct {
	Client *dbmysql.DB
}

//New create a mysql store
func New(client *dbmysql.DB) *Store {
	return &Store{
		Client: client,
	}
}

func (s *Store) Type() string {
	return "mysql"
}
