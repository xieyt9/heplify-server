package search

import (
	"github.com/emicklei/go-restful"
	"github.com/sipcapture/heplify-server/pkg/api"
	"net/http"
)

func init() {
	installSearchData(restful.DefaultContainer)
}

func installSearchData(container *restful.Container) {
	searchdata := new(restful.WebService)
	searchdata.
		Path("/api/v1/search").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON) // you can specify this per route as well
	container.Add(searchdata)

	route := searchdata.POST("data").To(GetSIPData).
		Doc("session api").
		Returns(http.StatusOK, http.StatusText(http.StatusOK), []api.SessionRespData{}).
		Operation("session")
	searchdata.Route(route)

	route = searchdata.POST("method").To(GetMethodData).
		Doc("session api").
		Returns(http.StatusOK, http.StatusText(http.StatusOK), []api.SessionRespData{}).
		Operation("session")
	searchdata.Route(route)

	route = searchdata.POST("transaction").To(GetTransactionData).
		Doc("session api").
		Returns(http.StatusOK, http.StatusText(http.StatusOK), []api.SessionRespData{}).
		Operation("session")
	searchdata.Route(route)

	route = searchdata.POST("export/text").To(GetTextData).
		Doc("session api").
		Returns(http.StatusOK, http.StatusText(http.StatusOK), []api.SessionRespData{}).
		Operation("session")
	searchdata.Route(route)

}
