package api

import (
	"github.com/golang/glog"
	"github.com/sipcapture/heplify-server/pkg/api/apierr"
)

//EncodeError encode error
func EncodeError(status interface{}) []byte {
	var output []byte
	internalErr, ok := status.(*apierr.StatusError)
	if ok {
		output = internalErr.ErrStatus.Encode()
	} else {
		glog.Errorln("status type error")
	}

	return output
}
