package twistream

type Geo struct {
	Type        string     `json:"type"`
	Coordinates [2]float32 `json:"coordinates"`
}
