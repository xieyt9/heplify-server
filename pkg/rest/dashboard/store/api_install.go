package store

import (
	"net/http"

	"github.com/sipcapture/heplify-server/pkg/api"

	"github.com/emicklei/go-restful"
)

func init() {
	install(restful.DefaultContainer)
}

func install(container *restful.Container) {
	store := new(restful.WebService)
	store.
		Path("/api/v1/dashboard/store").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON) // you can specify this per route as well
	container.Add(store)

	route := store.GET("").To(GetStore).
		Doc("session api").
		Returns(http.StatusOK, http.StatusText(http.StatusOK), []api.SessionRespData{}).
		Operation("session")
	store.Route(route)

	route = store.GET("home").To(GetStoreHome).
		Doc("session api").
		Returns(http.StatusOK, http.StatusText(http.StatusOK), api.DashboardStoreHome{}).
		Operation("storehome")
	store.Route(route)
}
