package controller

import (
	"bluebell_gyf/logic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
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

func CommunityDetailHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	communityDetail, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("CommunityDetailHandler login error", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, communityDetail)
}
