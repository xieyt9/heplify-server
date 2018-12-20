package app

import (
	"github.com/sipcapture/heplify-server/pkg/master"

	"github.com/sipcapture/heplify-server/cmd/heplify-server/app/options"
)

// Run runs the specified APIServer.  This should never exit.
func Run(opt *options.SIPCapOptions) error {

	cfg := &master.Config{
		APIServerOpt: opt.Server,
		HomerDataDSN: opt.HomerDataDSN,
	}

	master.InitDB(cfg)
	return master.Run(cfg)
}
