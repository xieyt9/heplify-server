package options

import "github.com/spf13/pflag"

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

//AddFlags add api server options
func (s *ServerOption) AddFlags(fs *pflag.FlagSet) {

	fs.IntVar(&s.InsecurePort, "insecure-port", s.InsecurePort, ""+
		"The port on which to serve unsecured, unauthenticated access. Default 80.")

	fs.IntVar(&s.SecurePort, "secure-port", s.SecurePort, ""+
		"The port on which to serve HTTPS with authentication and authorization. If 0, "+
		"don't serve HTTPS at all.")

	fs.StringVar(&s.SwaggerPath, "swagger-path", s.SwaggerPath, ""+
		"specific a path where found swagger index.html, if not will be disable swagger ui")

	// fs.StringVar(&s.LicenseFile, "license-file", s.LicenseFile, ""+
	// 	"specific a file that contains license")

	fs.StringVar(&s.TLSCertFile, "tls-cert-file", s.TLSCertFile, ""+
		"File containing x509 Certificate for HTTPS. (CA cert, if any, concatenated "+
		"after server cert). If HTTPS serving is enabled, must configure this")

	fs.StringVar(&s.TLSPrivateKeyFile, "tls-private-key-file", s.TLSPrivateKeyFile,
		"File containing x509 private key matching --tls-cert-file.")

	fs.StringVar(&s.UIPath, "ui-path", s.UIPath, ""+
		"specific a path where found  index.html of ui module, if not will be disable ui")
	fs.StringVar(&s.AdminPwd, "admin-pwd", s.AdminPwd, ""+
		"admin password")
}
