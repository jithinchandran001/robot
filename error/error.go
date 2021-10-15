package error

import (
	"net/http"
	"robot/pkg/codecs"
)

var DBError *codecs.EndpointError = codecs.NewEndpointError(http.StatusInternalServerError, "DB error")
var SurvivorDataNotFound *codecs.EndpointError = codecs.NewEndpointError(http.StatusBadRequest, "survivor data not found")
var RobotListError *codecs.EndpointError = codecs.NewEndpointError(http.StatusInternalServerError, "robot data fetch error")

