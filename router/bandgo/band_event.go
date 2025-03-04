package bandgo

import (
	"fmt"

	// v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type BandEventRouter struct{}

func (r *BandEventRouter) InitBandEventRouter(Router *gin.RouterGroup) {
	eventRouter := Router.Group("event")
	fmt.Print(eventRouter)
	// bandEventApi := v1.ApiGroupApp.BandgoApiGroup.BandEventApi
	// {
	// 	eventRouter.POST("", bandEventApi.CreateEvent)              // 创建活动
	// 	eventRouter.GET("/:id", bandEventApi.GetEventInfo)         // 获取活动信息
	// 	eventRouter.PUT("/:id", bandEventApi.UpdateEvent)          // 更新活动
	// 	eventRouter.GET("", bandEventApi.ListEvents)               // 活动列表
	// 	eventRouter.POST("/:id/join", bandEventApi.JoinEvent)      // 参加活动
	// }
} 