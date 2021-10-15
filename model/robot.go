package model

type Robot struct {
	Model string `json:"model"`
	SerialNumber        string   `json:"serialNumber"`
	ManufacturedDate  string   `json:"manufacturedDate"`
	Category string `json:"category""`
}


type RobotData struct {
	LandRobot []Robot  `json:"land_robot"`
	AirRobot []Robot  `json:"air_robot"`
}
