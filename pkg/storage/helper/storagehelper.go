package helper

import (
	"github.com/sipcapture/heplify-server/pkg/storage"
	"github.com/sipcapture/heplify-server/pkg/storage/storagebackend"
	"github.com/sipcapture/heplify-server/pkg/storage/storagebackend/factory"
)

var homerData storage.Interface
var homerStatis storage.Interface
var homerCfg storage.Interface

//GetMysqlSotrage get global mysql store interface.
func GetMysqlSotrage() storage.Interface {
	return homerData
}

//GetHomerStatisSotrage get global mysql store interface.
func GetHomerStatisSotrage() storage.Interface {
	return homerStatis
}

//GetHomerCfgSotrage get global mysql store interface.
func GetHomerCfgSotrage() storage.Interface {
	return homerCfg
}

//CreateGlobalMysqlStorage create global mysql store
func CreateGlobalMysqlStorage(dsn string) (factory.DestroyFunc, error) {
	mysqlConf := storagebackend.Config{
		Type: storagebackend.StorageTypeMysql,
		Mysql: storagebackend.MysqlConfig{
			ServerList: []string{dsn},
		},
	}

	store, destory, err := factory.Create(mysqlConf)

	homerData = store

	return destory, err
}

//CreateGlobalHomeCfgStorage create global mysql store
func CreateGlobalHomeCfgStorage(dsn string) (factory.DestroyFunc, error) {
	mysqlConf := storagebackend.Config{
		Type: storagebackend.StorageTypeMysql,
		Mysql: storagebackend.MysqlConfig{
			ServerList: []string{dsn},
		},
	}

	store, destory, err := factory.Create(mysqlConf)

	homerCfg = store

	return destory, err
}

//CreateGlobalHomerStatislStorage create global mysql store
func CreateGlobalHomerStatislStorage(dsn string) (factory.DestroyFunc, error) {
	mysqlConf := storagebackend.Config{
		Type: storagebackend.StorageTypeMysql,
		Mysql: storagebackend.MysqlConfig{
			ServerList: []string{dsn},
		},
	}

	store, destory, err := factory.Create(mysqlConf)

	homerStatis = store

	return destory, err
}
