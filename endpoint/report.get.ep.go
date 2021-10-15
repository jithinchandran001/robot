package endpoint

import (
	"context"
	"net/http"
	"robot/pkg/codecs"
	"robot/service"

	"github.com/go-kit/kit/endpoint"
)

func makeReportEndpoint(srv service.RobotService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r, err := srv.GetReport()
		if err != nil {
			return nil, err
		}
		return r, nil
	}
}

func DecodeReportRequest(c context.Context, req *http.Request) (request interface{}, err error) {
	return nil, nil
}

func EncodeReportResponse(c context.Context, w http.ResponseWriter, response interface{}) error {
	codecs.ResponseSingleJson(response, w)
	return nil
}
