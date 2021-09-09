package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/go_checkpoint/config"
	"github.com/ydhnwb/go_checkpoint/controller"
	"github.com/ydhnwb/go_checkpoint/repo"
	"github.com/ydhnwb/go_checkpoint/service"
	"gorm.io/gorm"
)

var (
	db            *gorm.DB                 = config.SetupDBConnection()
	logRepo       repo.LogRepo             = repo.NewLogRepo(db)
	logService    service.LogService       = service.NewLogService(logRepo)
	logController controller.LogController = controller.NewLogController(logService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	server := gin.Default()

	logRoutes := server.Group("api/log")
	{
		logRoutes.POST("/", logController.CreateLog)
		logRoutes.GET("/", logController.FindAllLog)
	}

	server.Run()

}
