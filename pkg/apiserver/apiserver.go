package apiserver

import (
	"fmt"
	"net/http"

	"github.com/sipcapture/heplify-server/pkg/apiserver/options"
	"github.com/sipcapture/heplify-server/pkg/auth"

	"github.com/emicklei/go-restful"
	"github.com/golang/glog"
)

//APIServer ... http server configure
type APIServer struct {
	SecurePort        int
	InsecurePort      int
	SwaggerPath       string
	TLSCertFile       string
	TLSPrivateKeyFile string
	UIPath            string
	EnableUI          bool

	wsContainer *restful.Container
}

//NewAPIServer ...new a apiserver
func NewAPIServer(config *options.ServerOption) *APIServer {

	auth.SetAdminPasswrod(config.AdminPwd)
	return &APIServer{
		SecurePort:        config.SecurePort,
		InsecurePort:      config.InsecurePort,
		SwaggerPath:       config.SwaggerPath,
		TLSCertFile:       config.TLSCertFile,
		TLSPrivateKeyFile: config.TLSPrivateKeyFile,
		UIPath:            config.UIPath,
		EnableUI:          config.EnableUI,
	}
}

//Run ...start http server run
func (apis *APIServer) Run() error {

	//apis.wsContainer = restful.NewContainer()
	apis.wsContainer = restful.DefaultContainer
	apis.wsContainer.Router(restful.CurlyRouter{})

	apis.install(apis.wsContainer)
	//apis.installUI(apis.wsContainer)

	// Add container filter to enable CORS
	// cors := restful.CrossOriginResourceSharing{
	// 	ExposeHeaders:  []string{"X-My-Header"},
	// 	AllowedHeaders: []string{"Content-Type", "Accept"},
	// 	AllowedMethods: []string{"GET", "POST"},
	// 	CookiesAllowed: false,
	// 	Container:      apis.wsContainer}
	// apis.wsContainer.Filter(cors.Filter)

	port := apis.InsecurePort
	var tls bool
	if apis.SecurePort != 0 {
		tls = true
		port = apis.SecurePort
		if apis.TLSCertFile == "" || apis.TLSPrivateKeyFile == "" {
			return fmt.Errorf("must give cert and private key file")
		}
	}

	addr := ":" + fmt.Sprintf("%d", port)
	glog.V(5).Infof("server on (%v %v)", apis.InsecurePort, apis.SecurePort)
	server := &http.Server{Addr: addr, Handler: apis.wsContainer}

	if tls {
		if len(apis.SwaggerPath) > 0 {
			apis.installSwaggerAPI(apis.wsContainer, true, port)
		}
		return server.ListenAndServeTLS(apis.TLSCertFile, apis.TLSPrivateKeyFile)
	}

	if len(apis.SwaggerPath) > 0 {
		apis.installSwaggerAPI(apis.wsContainer, false, port)
	}
	return server.ListenAndServe()
}
