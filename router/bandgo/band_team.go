package bandgo

import (
	"fmt"
    "github.com/gin-gonic/gin"
    // v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
)

type BandTeamRouter struct{}

func (r *BandTeamRouter) InitBandTeamRouter(Router *gin.RouterGroup) {
    bandRouter := Router.Group("band")
	fmt.Print(bandRouter)
    // bandTeamApi := v1.ApiGroupApp.BandgoApiGroup.BandTeamApi
    // {
    //     bandRouter.POST("", bandTeamApi.CreateBand)                    // 创建乐队
    //     bandRouter.GET("/:id", bandTeamApi.GetBandInfo)               // 获取乐队信息
    //     bandRouter.PUT("/:id", bandTeamApi.UpdateBand)                // 更新乐队信息
    //     bandRouter.GET("/:id/members", bandTeamApi.GetBandMembers)    // 获取成员列表
    //     bandRouter.POST("/join", bandTeamApi.JoinRequest)             // 申请加入
    //     bandRouter.POST("/join/:id/process", bandTeamApi.ProcessJoinRequest) // 处理申请
    // }
} 