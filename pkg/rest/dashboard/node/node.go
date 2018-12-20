package node

import (
	"encoding/json"

	restful "github.com/emicklei/go-restful"
	"github.com/golang/glog"
	"github.com/sipcapture/heplify-server/pkg/api"
	"github.com/sipcapture/heplify-server/pkg/api/apierr"
)

func GetNode(request *restful.Request, response *restful.Response) {
	
	w := response.ResponseWriter
	w.Header().Set("Content-Type", "application/json")

	statusCode := 200
	var output []byte
	var err error

	defer func() {
		w.WriteHeader(statusCode)
		w.Write(output)
	}()

	node := api.DashboardNode{
		Status: 200,
		Sid: "s0h7q0ccs7gbokrvup56tb7604",
		Auth: "true",
		Message: "ok",
		Data: []api.DashboardNodeData{
			{
				Id: "1",
				Name: "node1",
			},
		},
		Count: 1,
	}
	output, err = json.MarshalIndent(node, "  ", "  ") 
//	output, err = json.Marshal(node)

	if err != nil {
		glog.Errorln("get node failure  ", err)
		newErr := apierr.NewInternalError(err.Error())
		output = api.EncodeError(newErr)
		statusCode = 500
		return
	}

	return
}