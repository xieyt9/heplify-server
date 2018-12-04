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
		Path("/api/v1/profile").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON) // you can specify this per route as well
	container.Add(store)

	route := store.GET("store").To(GetStore).
		Doc("session api").
		Returns(http.StatusOK, http.StatusText(http.StatusOK), []api.SessionRespData{}).
		Operation("session")
	store.Route(route)

	route = store.POST("store/timerange").To(GetTimeRange).
		Doc("search api").
		Returns(http.StatusOK, http.StatusText(http.StatusOK), []api.TimerRange{}).
		Operation("session")
	store.Route(route)

	route = store.POST("store/timezone").To(GetTimeZone).
		Doc("search api").
		Returns(http.StatusOK, http.StatusText(http.StatusOK), []api.TimeZone{}).
		Operation("session")
	store.Route(route)

	route = store.POST("store/search").To(Search).
		Doc("search api").
		Returns(http.StatusOK, http.StatusText(http.StatusOK), []api.ProfileStoreSearch{}).
		Operation("session")
	store.Route(route)

	route = store.POST("store/transaction").To(Search).
		Doc("search api").
		Returns(http.StatusOK, http.StatusText(http.StatusOK), []api.ProfileStoreSearch{}).
		Operation("session")
	store.Route(route)

	route = store.POST("store/result").To(Search).
		Doc("search api").
		Returns(http.StatusOK, http.StatusText(http.StatusOK), []api.ProfileStoreSearch{}).
		Operation("session")
	store.Route(route)

	route = store.POST("store/node").To(Search).
		Doc("search api").
		Returns(http.StatusOK, http.StatusText(http.StatusOK), []api.ProfileStoreSearch{}).
		Operation("session")
	store.Route(route)
}
