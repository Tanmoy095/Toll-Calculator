package types

type OBUData struct {
	OBUID     int     `json:"obuID"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"long"`
}

type Distance struct {
	Value float64 `json:"value"`
	OBUID int     `json:"obuID"`
	Unix  int64   `json:"unix"` //time.Unix
}
