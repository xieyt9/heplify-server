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

func GetSIPData(request *restful.Request, response *restful.Response) {

	w := response.ResponseWriter
	w.Header().Set("Content-Type", "application/json")

	statusCode := 200
	var output []byte
	var err error

	defer func() {
		w.WriteHeader(statusCode)
		w.Write(output)
	}()

	searchdatareq := api.SearchDataRequest{}
	err = request.ReadEntity(&searchdatareq)

	sipdata, err := query.QuerySIPCaptureCall(searchdatareq)
	if err != nil {
		return
	}

	for idx := range sipdata {
		tmp, err := strconv.Atoi(sipdata[idx].MicroTs)
		if err != nil {
			fmt.Println(tmp)
		}
		sipdata[idx].MicroTs = strconv.Itoa(tmp)
		sipdata[idx].MilliTs = strconv.Itoa(tmp / 1000)
		sipdata[idx].Date = utils.JsonTime(time.Unix(int64(tmp/1000000), 0))
		sipdata[idx].SourceAlias = sipdata[idx].SourceIp
		sipdata[idx].DestinationAlias = sipdata[idx].DestinationIp
	}

	searchdata := api.SearchData{
		Status:  200,
		Sid:     "rdilbfpmbr7p95lsf6c0pv0634",
		Auth:    "true",
		Message: "ok",
		Data:    sipdata,
	}

	output, err = json.MarshalIndent(searchdata, "  ", "  ")

	if err != nil {
		glog.Errorln("store failure  ", err)
		newErr := apierr.NewInternalError(err.Error())
		output = api.EncodeError(newErr)
		statusCode = 500
		return
	}

	return
}
