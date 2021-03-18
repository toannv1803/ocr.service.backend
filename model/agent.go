package model

type Agent struct {
	UserId string `json:"user_id,omitempty" bson:"user_id,omitempty" form:"user_id,omitempty"`
	Role   string `json:"roles,omitempty" bson:"roles,omitempty" form:"roles,omitempty"`
}
