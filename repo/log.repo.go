package repo

import (
	"github.com/ydhnwb/go_checkpoint/entity"
	"gorm.io/gorm"
)

type LogRepo interface {
	CreateLog(log entity.LogEntity) (entity.LogEntity, error)
	FindAllLog() ([]entity.LogEntity, error)
}

type logRepo struct {
	db *gorm.DB
}

func NewLogRepo(db *gorm.DB) LogRepo {
	return &logRepo{
		db: db,
	}
}

func (c *logRepo) CreateLog(log entity.LogEntity) (entity.LogEntity, error) {
	c.db.Save(&log)
	return log, nil

}

func (c *logRepo) FindAllLog() ([]entity.LogEntity, error) {
	logs := []entity.LogEntity{}
	c.db.Find(&logs)
	return logs, nil
}
