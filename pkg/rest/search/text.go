package search

import (
	//	"encoding/json"
	restful "github.com/emicklei/go-restful"
	"github.com/golang/glog"
	"github.com/sipcapture/heplify-server/pkg/api"
	"github.com/sipcapture/heplify-server/pkg/api/apierr"
	query "github.com/sipcapture/heplify-server/pkg/storage/query"
)

func GetTextData(request *restful.Request, response *restful.Response) {

	w := response.ResponseWriter
	w.Header().Set("Content-Type", "application/octet-stream")

	statusCode := 200
	var output []byte
	var err error

	defer func() {
		w.WriteHeader(statusCode)
		w.Write(output)
	}()

	searchtransreq := api.SearchTransRequest{}
	err = request.ReadEntity(&searchtransreq)

	textdata, err := query.ExportText(searchtransreq)
	var msgtext string
	for idx := range textdata {
		msgtext = msgtext + textdata[idx].TextHead + "\r\n" + textdata[idx].TextMsg + "\r\n\r\n"
	}
	output = []byte(msgtext)
	if err != nil {
		return
	}

	if err != nil {
		glog.Errorln("export text failure  ", err)
		newErr := apierr.NewInternalError(err.Error())
		output = api.EncodeError(newErr)
		statusCode = 500
		return
	}

	return
}
