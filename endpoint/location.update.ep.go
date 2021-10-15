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

func makeSurvivorUpdateEndpoint(srv service.RobotService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		u := request.(model.Survivor)
		r, err := srv.UpdateLocation(u)
		if codecs.IsEE(err) {
			return nil, err
		}
		return r, nil
	}
}

func DecodeLocationRequest(c context.Context, req *http.Request) (request interface{}, err error) {
	decoder := json.NewDecoder(req.Body)
	var survivor model.Survivor
	err = decoder.Decode(&survivor)
	if err != nil {
		return nil, err
	}
	return survivor, nil
}

func EncodeLocationResponse(c context.Context, w http.ResponseWriter, response interface{}) error {
	codecs.ResponseSingleJson(response, w)
	return nil
}
