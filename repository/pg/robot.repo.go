package pg

import (
	"github.com/go-pg/pg/v10"
	c "robot/constant"
	e "robot/error"
	"robot/model"
	"robot/pkg/logger"
)

type Repo struct {
	DBConn *pg.DB
}

func New(conn *pg.DB) RobotRepo {
	return &Repo{DBConn: conn}
}

func (u *Repo) AddSurvivor(survivors []model.Survivor) (string, error) {
	tx, err := u.DBConn.Begin()
	if err != nil {
		return "",e.DBError
	}
	defer tx.Close()
	for _, survivor := range survivors {
		_, err = u.DBConn.Model(&survivor).Returning("id").Insert()
		if err != nil {
			tx.Rollback()
			logger.Get().ErrorWithoutSTT("Exception in adding survivor", "Error", err.Error())
			return "", e.DBError
		}

		for i, _ := range survivor.Resources {
			survivor.Resources[i].SurvivorID = survivor.ID
		}
		_, err := u.DBConn.Model(&survivor.Resources).Insert()
		if err != nil {
			tx.Rollback()
			logger.Get().ErrorWithoutSTT("Exception in adding survivor", "Error", err.Error())
			return "", e.DBError
		}
	}
	tx.Commit()
	return c.MsgAddSurvivorSuccess, nil
}

func (u *Repo) GetSurvivors() ([]model.Survivor, error) {
	var survivors []model.Survivor
	err := u.DBConn.Model(&survivors).Relation("Resources", func(q *pg.Query) (*pg.Query, error) {
		return q, nil
	}).Select()
	if err != nil {
		if err != nil && err == pg.ErrNoRows {
			return survivors, e.SurvivorDataNotFound
		}
		logger.Get().ErrorWithoutSTT("Exception in getting product", "Error", err.Error())
		return survivors, e.DBError
	}
	return survivors, nil
}

func (u *Repo) GetSurvivorByID(id uint64) (model.Survivor, error) {
	var survivor model.Survivor
	err := u.DBConn.Model(&survivor).Where("id = ?",id).Relation("Resources", func(q *pg.Query) (*pg.Query, error) {
		return q, nil
	}).Select()
	if err != nil {
		if err != nil && err == pg.ErrNoRows {
			return survivor, e.SurvivorDataNotFound
		}
		logger.Get().ErrorWithoutSTT("Exception in getting product", "Error", err.Error())
		return survivor, e.DBError
	}
	return survivor, nil
}


func (u *Repo) UpdateLocation(survivor model.Survivor) (string, error) {
	r, err := u.DBConn.Model(&survivor).Column("longitude","latitude").WherePK().Update()
	if err != nil {
		logger.Get().ErrorWithoutSTT("Exception in updating product", "Error", err.Error())
		return "", e.DBError
	}
	if r.RowsAffected() == 0 {
		return c.MsgSurvivorUpdateNotAffected, nil

	}
	return c.MsgsurvivorUpdateSuccess, nil
}

func (u *Repo) UpdateInfected(survivor model.Survivor) (string, error) {

	r, err := u.DBConn.Model(&survivor).Set("infected = infected+1").WherePK().Update()
	if err != nil {
		logger.Get().ErrorWithoutSTT("Exception in updating product", "Error", err.Error())
		return "", e.DBError
	}
	if r.RowsAffected() == 0 {
		return c.MsgSurvivorUpdateNotAffected, nil

	}
	return c.MsgsurvivorUpdateSuccess, nil
}

func (u *Repo) GetSurvivorReport() (model.SurvivorReport, error) {
	var survivors []model.Survivor
	report := model.SurvivorReport{}
	err := u.DBConn.Model(&survivors).Select()
	if err != nil {
		if err != nil && err == pg.ErrNoRows {
			return report, e.SurvivorDataNotFound
		}
		logger.Get().ErrorWithoutSTT("Exception in getting product", "Error", err.Error())
		return report, e.DBError
	}
	infected := model.SurvivorData{}
	nonInfected := model.SurvivorData{}
	infectedList := []model.SurvivorList{}
	nonInfectedList := []model.SurvivorList{}
	for i, _ := range survivors {
		var survivorData model.SurvivorList
		survivorData.ID = survivors[i].ID
		survivorData.Name = survivors[i].Name
		survivorData.Gender =survivors[i].Gender
		survivorData.Latitude = survivors[i].Latitude
		survivorData.Longitude = survivors[i].Longitude
		if survivors[i].Infected >= c.IsInfectedCount {
			infectedList = append(infectedList, survivorData)
		}else {
			nonInfectedList = append(nonInfectedList, survivorData)
		}
	}
	totalCount := len(survivors)
	infectedCount := len(infectedList)
	nonInfectedCount := len(nonInfectedList)
	infected.Survivors = infectedList
	nonInfected.Survivors= nonInfectedList
	infected.Percentage = (float64(infectedCount)/float64(totalCount))*100
	nonInfected.Percentage = (float64(nonInfectedCount)/float64(totalCount))*100
	report.Infected = infected
	report.NonInfected = nonInfected
	return report, nil
}
