package types

type OBUData struct {
	OBUID     int     `json:"obu_id"` // Use json tags for proper JSON marshaling
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
