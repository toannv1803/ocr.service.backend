package model
type TDOPutObject struct {
	UserId   string `json:"user_id,omitempty" bson:"user_id,omitempty"`
	RequestId  string `json:"request_id,omitempty" bson:"request_id,omitempty"`
}