package rest

import (
	"github.com/golang/glog"

	_ "github.com/sipcapture/heplify-server/pkg/rest/dashboard/store"
	_ "github.com/sipcapture/heplify-server/pkg/rest/profile/store"
	_ "github.com/sipcapture/heplify-server/pkg/rest/session"
	_ "github.com/sipcapture/heplify-server/pkg/rest/dashboard/node"
	_ "github.com/sipcapture/heplify-server/pkg/rest/search"
)

func init() {
	glog.Infof("install rest api")
}
