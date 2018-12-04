package factory

import (
	"github.com/sipcapture/heplify-server/pkg/storage"
	"github.com/sipcapture/heplify-server/pkg/storage/mysqls"
	"github.com/sipcapture/heplify-server/pkg/storage/storagebackend"

	_ "github.com/go-sql-driver/mysql"
	dbmysql "github.com/jinzhu/gorm"
)

//connectionStr: user:password@tcp(host:port)/dbname
func newMysqlClient(connectionStr string) (*dbmysql.DB, error) {
	var err error
	connStr := string(connectionStr) + string("?parseTime=True&loc=Local")
	//connStr := string(connectionStr)
	db, err := dbmysql.Open(string("mysql"), connStr)
	if err != nil {
		return nil, err
	}
	//db = db.Debug()

	return db, db.DB().Ping()
}

func newMysqlStorage(c storagebackend.Config) (storage.Interface, DestroyFunc, error) {
	endpoints := c.Mysql.ServerList

	client, err := newMysqlClient(endpoints[0])
	if err != nil {
		return nil, nil, err
	}

	destroyFunc := func() {
		client.Close()
	}

	return mysqls.New(client), destroyFunc, nil
}
