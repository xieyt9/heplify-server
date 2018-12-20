package session

import (
	"net/http"

	"github.com/sipcapture/heplify-server/pkg/api"

	"github.com/emicklei/go-restful"
)

func init() {
	installsession(restful.DefaultContainer)
}

func installsession(container *restful.Container) {
	session := new(restful.WebService)
	session.
		Path("/api/v1/session").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON) // you can specify this per route as well
	container.Add(session)

	route := session.POST("").To(POSTSession).
		Doc("session api").
		Returns(http.StatusOK, http.StatusText(http.StatusOK), []api.SessionRespData{}).
		Operation("session")
	session.Route(route)

	route = session.GET("").To(GetSession).
		Doc("session api").
		Returns(http.StatusOK, http.StatusText(http.StatusOK), []api.SessionRespData{}).
		Operation("session")
	session.Route(route)

}
