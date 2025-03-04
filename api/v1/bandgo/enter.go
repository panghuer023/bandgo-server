package bandgo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service"
)

type ApiGroup struct {
	BaseApi
	BandUserApi
	// BandTeamApi
	// BandEventApi
	// WeiXinApi
}

var (
	bandUserService = service.ServiceGroupApp.BandgoServiceGroup.BandUserService
	bandTeamService = service.ServiceGroupApp.BandgoServiceGroup.BandTeamService
	bandEventService = service.ServiceGroupApp.BandgoServiceGroup.BandEventService

	WeiXinService = service.ServiceGroupApp.BandgoServiceGroup.WeiXinService
) 