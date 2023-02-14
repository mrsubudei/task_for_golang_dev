package v1

import "go.mongodb.org/mongo-driver/bson/primitive"

type Answer interface {
	getCode() int
}

type ErrMessage struct {
	Error  string `json:"error,omitempty"`
	Detail string `json:"detail,omitempty"`
	code   int
}

func (e ErrMessage) getCode() int {
	return e.code

}

type Respone struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Email    string             `json:"email,omitempty"`
	Salt     string             `json:"salt,omitempty"`
	Password string             `json:"password,omitempty"`
	code     int
}

func (r Respone) getCode() int {
	return r.code
}

type Key string

const userKey = "user"

const (
	JsonNotCorrect     = "json format is not correct"
	EmptyFiledRequest  = "json body has empty fields"
	EmailParamEmpty    = "email parameter is empty"
	WrongDataFormat    = "wrong data format"
	WrongEmailFormat   = "email has wrong format"
	AdvertCreated      = "advert created"
	EmailFieldEmpty    = "'email:' field is empty"
	PasswordFieldEmpty = "'password:' field is empty"
)
