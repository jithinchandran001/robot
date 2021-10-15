package endpoint

import (
	"context"
	"github.com/gorilla/mux"
	"net/http"
	"robot/pkg/codecs"
	"robot/pkg/logger"
	"strconv"

	"robot/service"

	"github.com/go-kit/kit/endpoint"
)

func makeSurvivorGetEndpoint(srv service.RobotService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		u := request.(uint64)
		r, err := srv.GetSurvivorByID(u)
		if err != nil {
			return nil, err
		}
		return r, nil
	}
}

func DecodeSurvivorGetRequest(c context.Context, req *http.Request) (request interface{}, err error) {
	vars := mux.Vars(req)
	userId := vars["id"]
	id, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		logger.Get().ErrorWithoutSTT("Error in converting userID ", "id", userId)
		return nil, err
	}
	return uint64(id), nil
}

func EncodeSurvivorGetResponse(c context.Context, w http.ResponseWriter, response interface{}) error {
	codecs.ResponseSingleJson(response, w)
	return nil
}
