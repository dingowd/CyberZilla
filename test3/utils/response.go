package utils

import "encoding/json"

type Err struct {
	Err string `json:"error"`
}

func ReturnError(msg string) []byte {
	e := Err{
		Err: msg,
	}
	b, _ := json.Marshal(e)
	return b
}

type Res struct {
	Res string `json:"result"`
}

func ReturnStatus(msg string) []byte {
	r := Res{
		Res: msg,
	}
	b, _ := json.Marshal(r)
	return b
}
