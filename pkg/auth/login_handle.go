package auth

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/sipcapture/heplify-server/pkg/api"
	"github.com/sipcapture/heplify-server/pkg/api/apierr"
	"github.com/sipcapture/heplify-server/pkg/api/validation"

	restful "github.com/emicklei/go-restful"

	"github.com/golang/glog"
	"github.com/seanchann/goutil/cache"
)

const (
	adminUserName = "admin"
)

var adminPassword = "test123"
var adminTokenCache = cache.NewLRUExpireCache(5)

//SetAdminPasswrod set admin password
func SetAdminPasswrod(pwd string) {
	if len(pwd) > 0 {
		adminPassword = pwd
	}
}

func randBearerToken() (string, error) {
	token := make([]byte, 16)
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", token), err
}

func newToken(login *api.Login) error {
	if len(login.Spec.AuthName) == 0 || len(login.Spec.Auth) == 0 {
		return fmt.Errorf("invalid request")
	}

	if login.Spec.Auth != adminPassword {
		return fmt.Errorf("auth failure")
	}

	token, err := randBearerToken()
	if err != nil {
		glog.Errorf("generate token failure %v\r\n", err)
		return err
	}
	//adminToken[token] = adminUserName
	adminTokenCache.Add(token, adminUserName, time.Duration(time.Minute*60))

	login.Spec.Token = token
	login.Spec.AuthID = "admin"
	login.Spec.Auth = string("")

	return err
}

//CheckToken check token
func CheckToken(input string) (string, error) {
	parts := strings.Split(input, " ")
	if len(parts) < 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", fmt.Errorf("invalid token")
	}

	token := parts[1]
	userInterface, ok := adminTokenCache.Get(token)
	if !ok {
		return "", fmt.Errorf("invalid token")
	}
	user, ok := userInterface.(string)
	return user, nil
}

//PostLogin handle login method
func PostLogin(request *restful.Request, response *restful.Response) {
	w := response.ResponseWriter

	statusCode := 200
	output := apierr.NewSuccess().Encode()

	defer func() {
		w.WriteHeader(statusCode)
		w.Write(output)
	}()

	login := new(api.Login)
	err := request.ReadEntity(login)
	if err != nil {
		glog.Errorf("invalid request body:%v", err)
		newErr := apierr.NewBadRequestError("request body invalid")
		output = api.EncodeError(newErr)
		statusCode = 400
		return
	}

	err = validation.ValidateLogin(*login)
	if err != nil {
		newErr := apierr.NewBadRequestError(err.Error())
		output = api.EncodeError(newErr)
		statusCode = 400
		return
	}

	// glog.Infof("Got Post logins:%+v\n", login)
	err = newToken(login)
	if err != nil {
		newErr := apierr.NewBadRequestError(err.Error())
		output = api.EncodeError(newErr)
		statusCode = 400
	} else {
		glog.V(5).Infof("Got login len %+v", len(login.Spec.Auth))
		output, err = json.Marshal(login)
		statusCode = 200
	}

	return

}
