package apiserver

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/sipcapture/heplify-server/pkg/api"
	"github.com/sipcapture/heplify-server/pkg/auth"

	_ "github.com/sipcapture/heplify-server/pkg/rest"

	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful/swagger"
	"github.com/golang/glog"
)

func installLoginSrv(container *restful.Container) {
	ws := new(restful.WebService)
	ws.
		Path("/api/v1/logins").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON) // you can specify this per route as well
	container.Add(ws)
	route := ws.POST("").To(auth.PostLogin).
		Doc("get token").
		Returns(http.StatusOK, http.StatusText(http.StatusOK), api.Login{}).
		Param(ws.BodyParameter("body", "identifier of the login").DataType("api.Login")).
		Operation("PostLogin")
	ws.Route(route)
}

var uiRootPath string

func staticFromPathParam(req *restful.Request, resp *restful.Response) {
	actual := fmt.Sprintf("%s/%s", uiRootPath, req.PathParameter("subpath"))
	glog.V(3).Infof("serving %s ... (from %s)\n", actual, req.PathParameter("file"))
	http.ServeFile(
		resp.ResponseWriter,
		req.Request,
		actual)

}

func (apis *APIServer) installUI(container *restful.Container) {
	ws := new(restful.WebService)

	ws.Path("/")
	container.Add(ws)

	route := ws.GET("").To(staticFromPathParam).Produces("*/*")
	ws.Route(route)
	route = ws.GET("{subpath:*}").To(staticFromPathParam).Produces("*/*")
	ws.Route(route)

	uiRootPath = apis.UIPath
	glog.Infof("install ui %s", uiRootPath)
}

func (apis *APIServer) installSwaggerAPI(container *restful.Container, secure bool, port int) {
	//hostAndPort := apis.Host + string(":") + strconv.Itoa(apis.Port)
	hostAndPort := string(":") + strconv.Itoa(port)
	//protocol := "https://"
	protocol := "http://"
	if secure {
		protocol = "https://"
	}

	webServicesURL := protocol + hostAndPort

	// Enable swagger UI and discovery API
	swaggerConfig := swagger.Config{
		WebServicesUrl:  webServicesURL,
		WebServices:     container.RegisteredWebServices(),
		ApiPath:         "/swaggerapi/",
		SwaggerPath:     "/swaggerui/",
		SwaggerFilePath: apis.SwaggerPath,
	}
	swagger.RegisterSwaggerService(swaggerConfig, container)
}

func (apis *APIServer) install(container *restful.Container) error {
	installLoginSrv(container)
	if apis.EnableUI {
		apis.installUI(container)
	}
	return nil
}
