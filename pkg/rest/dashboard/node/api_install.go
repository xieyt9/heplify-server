package node

import (
	"net/http"

	"github.com/sipcapture/heplify-server/pkg/api"

	"github.com/emicklei/go-restful"
)

func init() {
	installnode(restful.DefaultContainer)
}

func installnode(container *restful.Container) {
	node := new(restful.WebService)
	node.
		Path("/api/v1/dashboard").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON) // you can specify this per route as well
	container.Add(node)

	route := node.GET("node").To(GetNode).
		Doc("node api").
		Returns(http.StatusOK, http.StatusText(http.StatusOK), api.DashboardNode{}).
		Operation("node")
	node.Route(route)

}

