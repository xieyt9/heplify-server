package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	//"net"
	//_ "net/http/pprof"

	"github.com/koding/multiconfig"
	"github.com/negbie/logp"
	"github.com/sipcapture/heplify-server/config"
	"github.com/sipcapture/heplify-server/server"

	"github.com/sipcapture/heplify-server/cmd/heplify-server/app"
	"github.com/sipcapture/heplify-server/cmd/heplify-server/app/options"

	//"github.com/seanchann/goutil/flag"
//	"github.com/seanchann/goutil/logs"
	//"github.com/spf13/pflag"
)

type server interface {
	Run()
	End()
}

func init() {
	var err error
	var logging logp.Logging
	var fileRotator logp.FileRotator

	//c := multiconfig.New()
	cfg := config.Get()
	//new(config.HeplifyServer)
	//c.MustLoad(cfg)
	config.Setting = *cfg

	if tomlExists(config.Setting.Config) {
		cf := multiconfig.NewWithPath(config.Setting.Config)
		err := cf.Load(cfg)
		if err == nil {
			config.Setting = *cfg
		} else {
			fmt.Println("Syntax error in toml config file, use flag defaults.", err)
		}
	} else {
		fmt.Println("Could not find toml config file, use flag defaults.", err)
	}

	logp.DebugSelectorsStr = &config.Setting.LogDbg
	logging.Level = config.Setting.LogLvl
	logp.ToStderr = &config.Setting.LogStd
	fileRotator.Path = "./"
	fileRotator.Name = "heplify-server.log"
	logging.Files = &fileRotator

	err = logp.Init("heplify-server", &logging)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func tomlExists(f string) bool {
	_, err := os.Stat(f)
	if os.IsNotExist(err) {
		return false
	} else if !strings.Contains(f, ".toml") {
		return false
	}
	return err == nil
}
func homer_main() {

	opt := options.NewSIPCapOptions()
//	logs.InitLogs()
//	defer logs.FlushLogs()
	opt.HomerDataDSN = fmt.Sprintf("%s:%s@tcp(%s)/homer_data",config.Setting.DBUser,
									config.Setting.DBPass,
									config.Setting.DBAddr)
	opt.Server.UIPath = config.Setting.UIPath
	opt.Server.SwaggerPath = config.Setting.SwaggerPath
	opt.Server.AdminPwd = config.Setting.AdminPwd
	opt.Server.InsecurePort = config.Setting.InsecurePort
	if err := app.Run(opt); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func main() {
	go homer_main()
	if config.Setting.Version {
		fmt.Println(config.Version)
		os.Exit(0)
	}
	var wg sync.WaitGroup
	var sigCh = make(chan os.Signal, 1)

	//go http.ListenAndServe(":8181", http.DefaultServeMux)

	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	hep := input.NewHEPInput()
	servers := []server{hep}

	for _, srv := range servers {
		wg.Add(1)
		go func(s server) {
			defer wg.Done()
			s.Run()
		}(srv)
	}

	<-sigCh

	for _, srv := range servers {
		wg.Add(1)
		go func(s server) {
			defer wg.Done()
			s.End()
		}(srv)
	}
	wg.Wait()
}
