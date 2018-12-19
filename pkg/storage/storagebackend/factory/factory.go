package factory

// Create creates a storage backend based on given config.
import (
	"fmt"

	"github.com/sipcapture/heplify-server/pkg/storage"
	"github.com/sipcapture/heplify-server/pkg/storage/storagebackend"
)

// DestroyFunc is to destroy any resources used by the storage returned in Create() together.
type DestroyFunc func()

//Create a storage interface
func Create(c storagebackend.Config) (storage.Interface, DestroyFunc, error) {
	switch c.Type {
	case storagebackend.StorageTypeUnset, storagebackend.StorageTypeMysql:
		return newMysqlStorage(c)
	default:
		return nil, nil, fmt.Errorf("unknown storage type: %s", c.Type)
	}
}
