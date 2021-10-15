package service

import (
	"robot/model"
)

type RobotService interface {
	Greeting() (string, error)
	AddSurvivor(survivor []model.Survivor) (string, error)
	GetSurvivors() ([]model.Survivor, error)
	GetSurvivorByID(id uint64) (model.Survivor, error)
	UpdateLocation(survivor model.Survivor) (string, error)
	UpdateInfected(req model.Infected) (string, error)
	GetRobotList() (*model.RobotData, error)
	GetReport() (*model.SurvivorReport, error)

}