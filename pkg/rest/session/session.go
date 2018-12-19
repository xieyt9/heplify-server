package session

import (
	"encoding/json"
	"time"

	"github.com/sipcapture/heplify-server/pkg/api"
	"github.com/sipcapture/heplify-server/pkg/api/apierr"

	restful "github.com/emicklei/go-restful"
	"github.com/golang/glog"
)

func GetSession(request *restful.Request, response *restful.Response) {

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

	session := api.SessionResp{
		Status:  200,
		SID:     "c6c1kuocp3kn2g47aidkh0vob3",
		Auth:    "true",
		Message: "ok",
		Data: api.SessionRespData{
			UID:       "1",
			Username:  "admin",
			GID:       "10",
			GRP:       "users,admins",
			FirstName: "Admin",
			LastName:  "Admin",
			Email:     "admin@test.com",
			LastVisit: time.Now(),
		},
	}

	output, err = json.Marshal(session)

	if err != nil {
		glog.Errorln("session failure  ", err)
		newErr := apierr.NewInternalError(err.Error())
		output = api.EncodeError(newErr)
		statusCode = 500
		return
	}

	return
}

func POSTSession(request *restful.Request, response *restful.Response) {

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

	session := api.SessionResp{
		Status:  200,
		SID:     "c6c1kuocp3kn2g47aidkh0vob3",
		Auth:    "true",
		Message: "ok",
		Data: api.SessionRespData{
			UID:       "1",
			Username:  "admin",
			GID:       "10",
			GRP:       "users,admins",
			FirstName: "Admin",
			LastName:  "Admin",
			Email:     "admin@test.com",
			LastVisit: time.Now(),
		},
	}

	output, err = json.Marshal(session)

	if err != nil {
		glog.Errorln("session failure  ", err)
		newErr := apierr.NewInternalError(err.Error())
		output = api.EncodeError(newErr)
		statusCode = 500
		return
	}

	return
}
