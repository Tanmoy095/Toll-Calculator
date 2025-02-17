package types

type ObuData struct {

	//we will send websocket as json
	OBUID int `json:"obuID"`
	//latitude
	Lat float64 `json:"lat"`
	//longitude
	Long float64 `json:"long"`
}
