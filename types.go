package godeluge

import(
	"encoding/json"
)

type Deluge struct{
	Session string
	Password string
	Id int32
}

type Request struct{
	Method string `json:"method"`
	Params interface{} `json:"params"`
	Id int32 `json:"id"`
}

type Response struct{
	Id int32
	Result json.RawMessage
	Error Error
}

type Error struct{
	Message string
	Code int32
}