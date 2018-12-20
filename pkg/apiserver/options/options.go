package options

// import "github.com/spf13/pflag"

//ServerOption contains apiserver options
type ServerOption struct {
	SecurePort        int
	InsecurePort      int
	SwaggerPath       string
	LicenseFile       string
	TLSCertFile       string
	TLSPrivateKeyFile string
	UIPath            string
	AdminPwd          string
}

//NewServerOption new apiserver options
func NewServerOption() *ServerOption {
	return &ServerOption{
		InsecurePort: 80,
		SecurePort:   0,
	}
}
