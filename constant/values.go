package c

//route URL
const (
	RouteGreeting = "/greetings"
	RouteSurvivor     = "/survivor"
	RouteInfected = "/infected"
	RouteRobot     = "/robot"
	RouteReset    = "/reset"
	RouteIDParam    = "/{id}"
	RouteReport = "/report"

)

const (
	RobotUrl ="https://robotstakeover20210903110417.azurewebsites.net/robotcpu"
)

const (
	CacheHeader = "X-VD-C"
)

const (
	MsgAddSurvivorSuccess      = "survivor added successfully"
	MsgsurvivorUpdateSuccess = "survivor updated successfully"
	MsgSurvivorUpdateNotAffected     = "survivor update not affected"
)
// A survivor is infected when count is updated 3 times
const (
	IsInfectedCount = 3
	)