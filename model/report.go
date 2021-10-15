package model

type Report struct {
	Total       int    `json:"total"`
	ProductName string `json:"product_name"`
	Balance     int    `json:"balance"`
}

type SurvivorReport struct {
	Infected SurvivorData `json:"infected"`
	NonInfected SurvivorData `json:"non_infected"`
	Robot *RobotData `json:"robots"`
}

type SurvivorData struct {
	Percentage float64 `json:"percentage"`
	Survivors []SurvivorList `json:"survivors"`
}

type SurvivorList struct {
	ID        uint64     `json:"id" `
	Name      string     `json:"name"`
	Gender    string     `json:"gender"`
	Longitude string     `json:"longitude"`
	Latitude  string     `json:"latitude"`
	Infected  int        `json:"infected"`
}
