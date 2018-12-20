package store

import (
	"encoding/json"
	restful "github.com/emicklei/go-restful"
	"github.com/golang/glog"
	"github.com/sipcapture/heplify-server/pkg/api"
	"github.com/sipcapture/heplify-server/pkg/api/apierr"
)

func GetTimeRange(request *restful.Request, response *restful.Response) {

	w := response.ResponseWriter
	w.Header().Set("Content-Type", "application/json")
	statusCode := 200
	var output []byte
	var err error

	defer func() {
		w.WriteHeader(statusCode)
		w.Write(output)
	}()

	timerangeres := api.ProfileStoreSearch{
		Status: 200,
		Data: []api.ProfileStoreSearchData{

		},
	}
	
	output, err = json.MarshalIndent(timerangeres, "  ", "  ")

	if err != nil {
		glog.Errorln("search failure  ", err)
		newErr := apierr.NewInternalError(err.Error())
		output = api.EncodeError(newErr)
		statusCode = 500
		return
	}	
	return
}