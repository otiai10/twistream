package twistream

type geo struct {
	Type        string     `json:"type"`
	Coordinates [2]float32 `json:"coordinates"`
}
