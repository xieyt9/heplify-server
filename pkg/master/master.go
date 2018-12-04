package master

import (
	"fmt"
	"sync"

	"github.com/sipcapture/heplify-server/pkg/apiserver"
	serveroption "github.com/sipcapture/heplify-server/pkg/apiserver/options"
	storagehelper "github.com/sipcapture/heplify-server/pkg/storage/helper"
)

//Config master config
type Config struct {
	APIServerOpt *serveroption.ServerOption

	HomerDataDSN string

	HomerStatisticDSN string

	HomerConfigurationDSN string
}

//Run start main loop
func Run(cfg *Config) error {

	destory, err := storagehelper.CreateGlobalMysqlStorage(cfg.HomerDataDSN)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	quit := make(chan struct{})
	DatabaseMaintainWorker(&wg, quit)

	serverHandler := apiserver.NewAPIServer(cfg.APIServerOpt)
	if serverHandler == nil {
		return fmt.Errorf("api server init failure")
	}
	if err := serverHandler.Run(); err != nil {
		return fmt.Errorf("api server run failure")
	}

	close(quit)
	wg.Wait()
	destory()
	return nil
}

func InitDB(cfg *Config) error {
	_, err := storagehelper.CreateGlobalMysqlStorage(cfg.HomerDataDSN)
	if err != nil {
		return err
	}

	return dbMaintain()
}
