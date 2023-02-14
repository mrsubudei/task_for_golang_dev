package v1

import "go.mongodb.org/mongo-driver/bson/primitive"

// Answer -.
type Answer interface {
	getCode() int
}

// ErrMessage -.
type ErrMessage struct {
	Error  string `json:"error,omitempty"`
	Detail string `json:"detail,omitempty"`
	Code   int    `json:"-"`
}

// getCode -.
func (e ErrMessage) getCode() int {
	return e.Code

}

// Response -.
type Response struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Email    string             `json:"email,omitempty"`
	Salt     string             `json:"salt,omitempty"`
	Password string             `json:"password,omitempty"`
	Code     int                `json:"-"`
}

// getCode -.
func (r Response) getCode() int {
	return r.Code
}

type Key string

const userKey = "user"

const (
	JsonNotCorrect      = "json format is not correct"
	EmptyFiledRequest   = "json body has empty fields"
	EmailParamEmpty     = "email parameter is empty"
	WrongDataFormat     = "wrong data format"
	WrongEmailFormat    = "email has wrong format"
	AdvertCreated       = "advert created"
	EmailFieldEmpty     = "'email:' field is empty"
	PasswordFieldEmpty  = "'password:' field is empty"
	InternalServerError = "Internal server error"
)
