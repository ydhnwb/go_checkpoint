package service

import (
	"github.com/mashingan/smapping"
	"github.com/ydhnwb/go_checkpoint/dto"
	"github.com/ydhnwb/go_checkpoint/entity"
	"github.com/ydhnwb/go_checkpoint/repo"
)

type LogService interface {
	CreateLog(logRequest dto.CreateLogRequest) (*dto.LogResponse, error)
	FindAllLog() (*[]dto.LogResponse, error)
}

type logService struct {
	logRepo repo.LogRepo
}

func NewLogService(logRepo repo.LogRepo) LogService {
	return &logService{
		logRepo: logRepo,
	}
}

func (c *logService) CreateLog(logRequest dto.CreateLogRequest) (*dto.LogResponse, error) {
	log := entity.LogEntity{}
	err := smapping.FillStruct(&log, smapping.MapFields(&logRequest))
	if err != nil {
		println("failed map when insert")
		return nil, err
	}
	log, err = c.logRepo.CreateLog(log)
	if err != nil {
		println("failed when insert to db")
		return nil, err
	}

	logResponse := dto.LogResponse{
		ID:   log.ID,
		Name: log.Name,
		Date: log.Date.String(),
	}
	return &logResponse, nil
}

func (c *logService) FindAllLog() (*[]dto.LogResponse, error) {
	logs, err := c.logRepo.FindAllLog()
	if err != nil {
		return nil, err
	}

	logResponses := []dto.LogResponse{}
	for _, v := range logs {
		temp := dto.LogResponse{
			ID:   v.ID,
			Name: v.Name,
			Date: v.Date.String(),
		}
		logResponses = append(logResponses, temp)
	}
	return &logResponses, nil

}
