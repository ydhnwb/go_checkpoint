package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/go_checkpoint/dto"
	"github.com/ydhnwb/go_checkpoint/service"
)

type LogController interface {
	CreateLog(ctx *gin.Context)
	FindAllLog(ctx *gin.Context)
}

type logController struct {
	logService service.LogService
}

func NewLogController(logService service.LogService) LogController {
	return &logController{
		logService: logService,
	}
}

func (c *logController) CreateLog(ctx *gin.Context) {
	var logReq dto.CreateLogRequest
	err := ctx.ShouldBind(&logReq)

	if err != nil {
		data := map[string]string{
			"message": err.Error(),
		}
		ctx.JSON(400, data)
		return
	}

	result, err := c.logService.CreateLog(logReq)
	if err != nil {
		data := map[string]string{
			"message": err.Error(),
		}
		ctx.JSON(400, data)
		return

	}

	ctx.JSON(200, result)
}

func (c *logController) FindAllLog(ctx *gin.Context) {
	result, err := c.logService.FindAllLog()
	if err != nil {
		data := map[string]string{
			"message": err.Error(),
		}
		ctx.JSON(400, data)
		return
	}

	ctx.JSON(200, result)

}
