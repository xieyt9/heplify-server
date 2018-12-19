package search

import (
	"encoding/json"
	"fmt"
	restful "github.com/emicklei/go-restful"
	"github.com/golang/glog"
	"github.com/sipcapture/heplify-server/pkg/api"
	"github.com/sipcapture/heplify-server/pkg/api/apierr"
	query "github.com/sipcapture/heplify-server/pkg/storage/query"
	utils "github.com/sipcapture/heplify-server/pkg/utils"
	"strconv"
	"time"
)

func GetMethodData(request *restful.Request, response *restful.Response) {

	w := response.ResponseWriter
	w.Header().Set("Content-Type", "application/json")
	// encoded := request.Request.Header.Get("Authorization")
	statusCode := 200
	var output []byte
	var err error

	defer func() {
		w.WriteHeader(statusCode)
		w.Write(output)
	}()

	searchmethodreq := api.SearchMethodReq{}
	err = request.ReadEntity(&searchmethodreq)

	methoddata, err := query.QueryMethod(searchmethodreq)
	if err != nil {
		return
	}
	for idx := range methoddata {
		tmp, err := strconv.Atoi(methoddata[idx].MicroTs)
		if err != nil {
			fmt.Println(tmp)
		}
		methoddata[idx].MicroTs = strconv.Itoa(tmp)
		methoddata[idx].MilliTs = strconv.Itoa(tmp / 1000)
		methoddata[idx].Date = utils.JsonTime(time.Unix(int64(tmp/1000000), 0))
	}

	searchmethod := api.SearchMethod{
		Status:  200,
		Sid:     "7tlk6nsdf5br4e74mj9o7uf514",
		Auth:    "true",
		Message: "ok",
		Data:    methoddata,
		Count:   1,
	}
	output, err = json.MarshalIndent(searchmethod, "  ", "  ")

	if err != nil {
		glog.Errorln("store failure  ", err)
		newErr := apierr.NewInternalError(err.Error())
		output = api.EncodeError(newErr)
		statusCode = 500
		return
	}

	return
}
