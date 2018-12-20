package storagebackend

const (
	//StorageTypeUnset not used any storage backend
	StorageTypeUnset = ""

	//StorageTypeMysql the mysql storage backend
	StorageTypeMysql = "mysql"
)

// Config is configuration for creating a storage backend.
type Config struct {
	// Type defines the type of storage backend, e.g. mysql. Default ("") is "mysql".
	Type string

	//mysql config
	Mysql MysqlConfig
}

//MysqlConfig the configure of mysql driver
type MysqlConfig struct {
	// ServerList is the list of storage servers to connect with.
	ServerList []string
}
