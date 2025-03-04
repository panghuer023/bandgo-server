package bandgo

import (
	"github.com/gin-gonic/gin"
)

type BandBaseRouter struct{}

func (s *BandBaseRouter) InitBandBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("base")
	{
		// baseRouter.POST("login", baseApi.Login)
		baseRouter.POST("captcha", baseApi.Captcha)
		baseRouter.POST("wx/config", baseApi.WXConfig)
		baseRouter.POST("wx/login", baseApi.WXLogin)
	}
	return baseRouter
}
