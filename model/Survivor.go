package model

type Survivor struct {
	tableName struct{}   `pg:"survivor"`
	ID        uint64     `json:"id" pg:"id,pk"`
	Name      string     `json:"name" pg:"name"`
	Gender    string     `json:"gender" pg:"gender"`
	Longitude string     `json:"longitude" pg:"longitude"`
	Latitude  string     `json:"latitude" pg:"latitude"`
	Infected  int        `json:"infected" pg:"infected"`
	Resources []Resource `json:"resources" pg:"rel:has-many"`
}

type Resource struct {
	tableName    struct{}  `pg:"inventory"`
	ID           uint64    `json:"id" pg:"id,pk"`
	SurvivorID   uint64    `json:"survivor_id" pg:"survivor_id"`
	Survivor     *Survivor `json:"-" pg:"rel:has-one"`
	ResourceName string    `json:"resource_name" pg:"inventory_name"`
	Quantity     string    `json:"quantity" pg:"quantity"`
	Unit         string    `json:"unit" pg:"unit"`
}

type Infected struct {
	ID        uint64     `json:"id"`
	Infected  bool        `json:"infected"`
}