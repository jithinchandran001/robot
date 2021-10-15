package pg

import "robot/model"

type RobotRepo interface {

	//Survivor
	AddSurvivor(survivor []model.Survivor) (string, error)
	GetSurvivors() ([]model.Survivor, error)
	GetSurvivorByID(id uint64) (model.Survivor, error)
	UpdateLocation(survivor model.Survivor) (string, error)
	UpdateInfected(survivor model.Survivor) (string, error)
	GetSurvivorReport() (model.SurvivorReport, error)
}
