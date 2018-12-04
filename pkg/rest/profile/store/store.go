package store

import (
	"encoding/json"
	"time"

	restful "github.com/emicklei/go-restful"
	"github.com/golang/glog"
	"github.com/sipcapture/heplify-server/pkg/api"
	"github.com/sipcapture/heplify-server/pkg/api/apierr"
)

func GetStore(request *restful.Request, response *restful.Response) {

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

	// _, err := auth.CheckToken(encoded)
	// if err != nil {
	// 	glog.Errorln("Unauth request ", err)
	// 	newErr := apierr.NewUnauthorized("invalid token")
	// 	output = api.EncodeError(newErr)
	// 	statusCode = 401
	// 	return
	// }

	store := api.ProfileStoreGetResp{
		Status: 200,
		Data: api.ProfileStoreGetRespData{
			Range: api.TimerRange{
				From:       time.Now(),
				To:         time.Now(),
				Custom:     time.Now(),
				CustomFrom: time.Now(),
				CustomTo:   time.Now(),
			},
			TMZone: api.TimeZone{
				Value:  "-480",
				Name:   "GMT+8 CCT",
				Offset: "+0800",
			},
			Search: api.SearchConf{
				Limit:  "2000",
				Callid: "",
				ToUser: "",
			},
			Transaction: []api.Transaction{},
			Result: api.StoreResult{
				ResultType: api.StoreResultType{
					Name:  "table",
					Value: "TABLE",
				},
			},
			Node: api.NodeConf{
				Node: api.NodeItem{
					Name: "localhost",
					ID:   "localhost",
				},
				DBNodes: []api.DBNode{
					{
						ID:   "1",
						Name: "noded1",
					},
				},
			},
		},
	}

	output, err = json.Marshal(store)

	if err != nil {
		glog.Errorln("store failure  ", err)
		newErr := apierr.NewInternalError(err.Error())
		output = api.EncodeError(newErr)
		statusCode = 500
		return
	}

	return
}
