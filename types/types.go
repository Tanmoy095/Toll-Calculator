package types

type OBUData struct {
	OBUID     int     `json:"obu_id"` // Use json tags for proper JSON marshaling
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Distance struct {
	value float64 `json:"value"`
	OBUID int     `json:"obuID"`
	Unix  int64   `json:"unix"` //time.Unix
}
