package endpoint

import (
	"robot/service"

	"github.com/go-kit/kit/endpoint"
)

type ServiceEndpoints struct {
	Greeting     endpoint.Endpoint
	AddSurvivor      endpoint.Endpoint
	UpdateLocation   endpoint.Endpoint
	GetSurvivors      endpoint.Endpoint
	GetSurvivor      endpoint.Endpoint
	UpdateInfected   endpoint.Endpoint
	GetRobot        endpoint.Endpoint
    GetReport       endpoint.Endpoint
}

// NewServiceEndpoints initiates new Endpoint
func NewServiceEndpoints(srv service.RobotService) *ServiceEndpoints {
	ep := &ServiceEndpoints{}
	ep.Greeting = makeGreetingEndpoint(srv)
	ep.AddSurvivor = makeAddSurvivorEndpoint(srv)
	ep.UpdateLocation = makeSurvivorUpdateEndpoint(srv)
	ep.UpdateInfected = makeInfectedUpdateEndpoint(srv)
	ep.GetSurvivors = makeGetSurvivorsEndpoint(srv)
	ep.GetSurvivor = makeSurvivorGetEndpoint(srv)
	ep.GetRobot = makeRobotGetEndpoint(srv)
	ep.GetReport = makeReportEndpoint(srv)
	return ep
}
