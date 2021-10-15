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

func makeInfectedUpdateEndpoint(srv service.RobotService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		u := request.(model.Infected)
		r, err := srv.UpdateInfected(u)
		if codecs.IsEE(err) {
			return nil, err
		}
		return r, nil
	}
}

func DecodeInfectedRequest(c context.Context, req *http.Request) (request interface{}, err error) {
	decoder := json.NewDecoder(req.Body)
	var survivor model.Infected
	err = decoder.Decode(&survivor)
	if err != nil {
		return nil, err
	}
	return survivor, nil
}

func EncodeInfectedResponse(c context.Context, w http.ResponseWriter, response interface{}) error {
	codecs.ResponseSingleJson(response, w)
	return nil
}
