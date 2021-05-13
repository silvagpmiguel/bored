package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/silvagpmiguel/bored/bored-backend/model"
	"github.com/silvagpmiguel/bored/bored-backend/repository"
)

// Service structure
type Service struct {
	repo *repository.Repository
}

// New returns a new service with an initialized repository
func New() (*Service, error) {
	repo, err := repository.InitRepository("db/oss.db")

	if err != nil {
		return nil, fmt.Errorf("failed to create new service : %v", err)
	}

	return &Service{repo}, nil
}

// GetBoredFromBoredAPI fetch bored result from bored API and convert it into an Activity
func GetBoredFromBoredAPI() (*model.Activity, error) {
	resp, err := http.Get("http://www.boredapi.com/api/activity/")
	bored := model.NewActivity()

	if err != nil {
		return bored, fmt.Errorf("error fetching response : %v", err)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return bored, fmt.Errorf("error reading body : %v", err)
	}

	resp.Body.Close()
	err = json.Unmarshal(body, &bored)

	if err != nil {
		return bored, fmt.Errorf("error building bored struct : %v", err)
	}

	return bored, nil
}

// CreateActivity create and persist new Activity in db
func (serv *Service) CreateActivity(activity *model.Activity) bool {
	return serv.repo.CreateActivity(activity)
}

// CreateActivity get activity with id? and type? from db
func (serv *Service) GetActivity(_id string, _type string) ([]model.Activity, bool) {
	activity := model.NewActivity()
	return serv.repo.GetActivity(activity.WithID(_id).WithType(_type))
}

// UpdateActivity update and persist newActivity with id? and type? from db
func (serv *Service) UpdateActivity(_id string, _type string, newActivity model.Activity) bool {
	oldActivity := model.NewActivity()
	return serv.repo.UpdateActivity(oldActivity.WithID(_id).WithType(_type), newActivity)
}
