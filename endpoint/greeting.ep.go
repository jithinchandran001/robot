package endpoint

import (
	"context"
	"net/http"
	"robot/pkg/codecs"

	"robot/service"

	"github.com/go-kit/kit/endpoint"
)

func makeGreetingEndpoint(srv service.RobotService) endpoint.Endpoint {
	// Each endpoint can return an error of type EndpointError,
	// and if so, the response writer will handle error writing
	// appropriately. It is also possible for a Service method
	// (srv.Greeting()) in our example, to 	return an error of type
	// endpoint.EndpointError too, which then means, this method
	// simply directly return that error
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r, err := srv.Greeting()
		if codecs.IsEE(err) {
			return nil, err
		}
		return r, nil
	}
}

func DecodeGreetingIncomingRequest(context.Context, *http.Request) (request interface{}, err error) {
	return nil, nil
}

func EncodeGreetingOutgoingResponse(c context.Context, w http.ResponseWriter, response interface{}) error {
	codecs.ResponseSingleJson(response, w)
	return nil
}
