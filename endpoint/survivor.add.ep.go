package endpoint

import (
	"context"
	"encoding/json"
	"net/http"
	"robot/model"
	"robot/pkg/codecs"

	"robot/service"

	"github.com/go-kit/kit/endpoint"
)

func makeAddSurvivorEndpoint(srv service.RobotService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		u := request.([]model.Survivor)
		r, err := srv.AddSurvivor(u)
		if codecs.IsEE(err) {
			return nil, err
		}
		return r, nil
	}
}

func DecodeAddSurvivorRequest(c context.Context, req *http.Request) (request interface{}, err error) {
	decoder := json.NewDecoder(req.Body)
	var user []model.Survivor
	err = decoder.Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func EncodeAddSurvivorResponse(c context.Context, w http.ResponseWriter, response interface{}) error {
	codecs.ResponseCreatedJson(response, w)
	return nil
}
