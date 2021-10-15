package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	c "robot/constant"
	e "robot/error"
	"robot/model"
	"robot/repository/pg"
)

type RobotApp struct {
	robotRepo pg.RobotRepo
}

func New(r pg.RobotRepo) RobotService {
	return &RobotApp{robotRepo: r}
}

func (r *RobotApp) Greeting() (string, error) {
	return "Hello from greeting", nil
}

func (r *RobotApp) AddSurvivor(survivor []model.Survivor) (string, error) {
	var resp string
	var err error
	if resp, err = r.robotRepo.AddSurvivor(survivor); err != nil {
		return "", err
	}
	return resp, nil
}

func (r *RobotApp) GetSurvivors() ([]model.Survivor, error) {
	var resp []model.Survivor
	var err error
	if resp, err = r.robotRepo.GetSurvivors(); err != nil {
		return resp, err
	}
	return resp, nil
}

func (r *RobotApp) GetSurvivorByID(id uint64) (model.Survivor, error) {
	var resp model.Survivor
	var err error
	if resp, err = r.robotRepo.GetSurvivorByID(id); err != nil {
		return resp, err
	}
	return resp, nil
}

func (r *RobotApp) UpdateLocation(survivor model.Survivor) (string, error) {
	var resp string
	var err error
	if resp, err = r.robotRepo.UpdateLocation(survivor); err != nil {
		return resp, err
	}
	return resp, nil
}

func (r *RobotApp) UpdateInfected(req model.Infected) (string, error) {
	var resp string
	var err error
	if !req.Infected {
		return c.MsgSurvivorUpdateNotAffected, nil
	}

	if resp, err = r.robotRepo.UpdateInfected(model.Survivor{
		ID: req.ID,
	}); err != nil {
		return resp, err
	}
	return resp, nil
}

func (r *RobotApp) GetRobotList() (*model.RobotData, error) {
	var resp model.RobotData
	var robotResp []model.Robot
	response, err := http.Get(c.RobotUrl)

	if err != nil {
		return nil,e.RobotListError
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil,e.RobotListError
	}
	defer response.Body.Close()
	err = json.Unmarshal(body, &robotResp)
	if err != nil {
		return nil,err
	}
	robotLand := []model.Robot{}
	robotFly := []model.Robot{}
	for i, _ := range robotResp {
		if robotResp[i].Category == "Land" {
			robotLand = append(robotLand,robotResp[i])
		} else {
			robotFly = append(robotFly,robotResp[i])
		}
	}
	resp.LandRobot = robotLand
	resp.AirRobot = robotFly

	return &resp,nil
}



func (r *RobotApp) GetReport() (*model.SurvivorReport, error) {

	report,err := r.robotRepo.GetSurvivorReport()
	if err != nil {
		return nil,err
	}
	robotList,err := r.GetRobotList()
	if err != nil {
		return nil,err
	}
	report.Robot = robotList
	return &report,nil

}

