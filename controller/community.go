package controller

import (
	"bluebell_gyf/logic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CommunityHandler(c *gin.Context) {
	communityList, err := logic.GetAllCommunityList()
	if err != nil {
		zap.L().Error("CommunityHandler login error", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, communityList)
}
