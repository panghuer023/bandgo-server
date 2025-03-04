package bandgo

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type BandUserRouter struct{}

func (r *BandUserRouter) InitBandUserPublicRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")
	bandUserApi := v1.ApiGroupApp.BandgoApiGroup.BandUserApi
	{
		userRouter.POST("register", bandUserApi.Register)
		userRouter.POST("login", bandUserApi.Login)
	}
}

func (r *BandUserRouter) InitBandUserPrivateRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")
	bandUserApi := v1.ApiGroupApp.BandgoApiGroup.BandUserApi
	{
		// userRouter.POST("register", bandUserApi.Register)
		userRouter.GET("profile", bandUserApi.GetProfile)
		userRouter.PUT("profile", bandUserApi.UpdateProfile)
		userRouter.GET("skills", bandUserApi.GetUserSkills)
		// userRouter.POST("skills", bandUserApi.AddUserSkill)
		// userRouter.PUT("skills/:id", bandUserApi.UpdateUserSkill)
	}
}
