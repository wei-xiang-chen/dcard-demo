package rest

type RestResult struct {
	Data  interface{} `json:"data"`
	Error *RestError  `json:"error"`
}
