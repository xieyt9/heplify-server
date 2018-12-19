package options

import (
	serveroption "github.com/sipcapture/heplify-server/pkg/apiserver/options"
	//	"github.com/spf13/pflag"
)

//SIPCapOptions  app options
type SIPCapOptions struct {
	//mysql dsn string
	HomerDataDSN string

	HomerStatisticDSN string

	HomerConfigurationDSN string

	Server *serveroption.ServerOption
}

//NewSIPCapOptions create  options
func NewSIPCapOptions() *SIPCapOptions {
	return &SIPCapOptions{
		Server: serveroption.NewServerOption(),
	}
}
