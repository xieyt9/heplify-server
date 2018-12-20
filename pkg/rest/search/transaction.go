package search

import (
	"encoding/json"
	restful "github.com/emicklei/go-restful"
	"github.com/golang/glog"
	"github.com/sipcapture/heplify-server/pkg/api"
	"github.com/sipcapture/heplify-server/pkg/api/apierr"
	query "github.com/sipcapture/heplify-server/pkg/storage/query"
)

func GetTransactionData(request *restful.Request, response *restful.Response) {

	w := response.ResponseWriter
	w.Header().Set("Content-Type", "application/json")

	statusCode := 200
	var output []byte
	var err error

	defer func() {
		w.WriteHeader(statusCode)
		w.Write(output)
	}()

	searchtransreq := api.SearchTransRequest{}
	err = request.ReadEntity(&searchtransreq)

	transdata, err := query.QueryTransactionData(searchtransreq)
	if err != nil {
		return
	}
	transaction := api.SearchTransaction{
		Status:  200,
		Sid:     "rdilbfpmbr7p95lsf6c0pv0634",
		Auth:    "true",
		Message: "ok",
		Data:    transdata,
		Count:   6,
	}

	output, err = json.MarshalIndent(transaction, "  ", "  ")
	//	output, err = json.Marshal(store)

	if err != nil {
		glog.Errorln("store failure  ", err)
		newErr := apierr.NewInternalError(err.Error())
		output = api.EncodeError(newErr)
		statusCode = 500
		return
	}

	return
}
