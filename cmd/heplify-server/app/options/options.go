package options

import (
	serveroption "github.com/sipcapture/heplify-server/pkg/apiserver/options"

	"github.com/spf13/pflag"
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

//AddFlags add flags
func (o *SIPCapOptions) AddFlags(fs *pflag.FlagSet) {

	fs.StringVar(&o.HomerDataDSN, "homer-data-dsn", "dbuser:dbpwd@tcp(127.0.0.1:3306)/homer_data?charset=utf8&parseTime=True&loc=Local", ""+
		"mysql connect string to homer_data db")
	fs.StringVar(&o.HomerStatisticDSN, "homer-statis-dsn", "dbuser:dbpwd@tcp(127.0.0.1:3306)/homer_statistic?charset=utf8&parseTime=True&loc=Local", ""+
		"mysql connect string to homer_statistic db")
	fs.StringVar(&o.HomerConfigurationDSN, "homer-cfg-dsn", "dbuser:dbpwd@tcp(127.0.0.1:3306)/homer_configuration?charset=utf8&parseTime=True&loc=Local", ""+
		"mysql connect string to homer_configuration db")

	o.Server.AddFlags(fs)
}
