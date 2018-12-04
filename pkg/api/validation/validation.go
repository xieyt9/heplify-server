package validation

import (
	"fmt"

	"github.com/sipcapture/heplify-server/pkg/api"
)

//ValidateLogin validate login
func ValidateLogin(login api.Login) error {

	if !(len(login.Spec.AuthName) > 0) {
		return fmt.Errorf("invalid authname")
	}

	if !(len(login.Spec.Auth) > 0) {
		return fmt.Errorf("invalid auth")
	}

	return nil
}
