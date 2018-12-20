package store

import (
	"encoding/json"

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

	store := api.DashboardStore{
		Status: 200,
		Data: []api.DashboardStoreDataItem{
			{
				Name:  "Home",
				Href:  "home",
				Class: "fa fa-home",
				ID:    "_1427728371642",
			},
			{
				Name:  "SIP Search",
				Href:  "search",
				Class: "fa fa-search",
				ID:    "_1426001444630",
			},
			{
				Name:  "Alarms",
				Href:  "alarms",
				Class: "fa fa-warning",
				ID:    "_1431721484444",
			},
			{
				Name:     "Panels",
				Href:     "#",
				Class:    "fa fa-dashboard",
				RowClass: "fa fa-angle-left pull-right",
				SubItems: []api.DashboardStoreDataItemSubItem{
					{
						BoardID: "_1450278390101",
						Name:    "Get Started",
						Class:   "fa fa-angle-double-right",
					},
					{
						BoardID: "_1411848944046",
						Name:    "Stats: IP Network",
						Class:   "fa fa-angle-double-right",
					},
					{
						BoardID: "_1430318378410",
						Name:    "System Admin",
						Class:   "fa fa-angle-double-right",
					},
					{
						BoardID: "_1452069187383",
						Name:    "Geo Chart",
						Class:   "fa fa-angle-double-right",
					},
					{
						BoardID: "_1428431423814",
						Name:    "Stats: VoIP Traffic",
						Class:   "fa fa-angle-double-right",
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
