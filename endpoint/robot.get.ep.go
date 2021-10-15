package endpoint

import (
	"context"
	"net/http"
	"robot/pkg/codecs"
	"robot/service"

	"github.com/go-kit/kit/endpoint"
)

func makeRobotGetEndpoint(srv service.RobotService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r, err := srv.GetRobotList()
		if err != nil {
			return nil, err
		}
		return r, nil
	}
}

func DecodeRobotGetRequest(c context.Context, req *http.Request) (request interface{}, err error) {
	return nil, nil
}

func EncodeRobotGetResponse(c context.Context, w http.ResponseWriter, response interface{}) error {
	codecs.ResponseSingleJson(response, w)
	return nil
}
