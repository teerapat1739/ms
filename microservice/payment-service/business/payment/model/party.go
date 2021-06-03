package model

type Request struct {
	Email string `json:"email"`
	ID    string `json:"id"`
}

type Return struct {
	Data     interface{} `json:"data"`
	ErrorMsg string      `json:"errorMsg"`
}
