package model

type Image struct {
	Id       string `json:"id,omitempty" bson:"id,omitempty" form:"id,omitempty"`
	UserId   string `json:"user_id,omitempty" bson:"user_id,omitempty" form:"user_id,omitempty"`
	BlockId  string `json:"block_id,omitempty" bson:"block_id,omitempty" form:"block_id,omitempty"`
	Path     string `json:"path,omitempty" bson:"path,omitempty" form:"path,omitempty"`
	Data     string `json:"data,omitempty" bson:"data,omitempty" form:"data,omitempty"`
	Status   string `json:"status,omitempty" bson:"status,omitempty" form:"status,omitempty"`
	Error    string `json:"error,omitempty" bson:"error,omitempty" form:"error,omitempty"`
	CreateAt string `json:"create_at,omitempty" bson:"create_at,omitempty" form:"create_at,omitempty"`
}

type ImageResponse struct {
	Id       string `json:"id,omitempty" bson:"id,omitempty" form:"id,omitempty"`
	UserId   string `json:"user_id,omitempty" bson:"user_id,omitempty" form:"user_id,omitempty"`
	BlockId  string `json:"block_id,omitempty" bson:"block_id,omitempty" form:"block_id,omitempty"`
	Data     string `json:"data,omitempty" bson:"data,omitempty" form:"data,omitempty"`
	Status   string `json:"status,omitempty" bson:"status,omitempty" form:"status,omitempty"`
	Error    string `json:"error,omitempty" bson:"error,omitempty" form:"error,omitempty"`
	CreateAt string `json:"create_at,omitempty" bson:"create_at,omitempty" form:"create_at,omitempty"`
}

type ImageTask struct {
	Id     string `json:"id,omitempty" bson:"id,omitempty" form:"id,omitempty"`
	UserId string `json:"user_id,omitempty" bson:"user_id,omitempty" form:"user_id,omitempty"`
	Path   string `json:"path,omitempty" bson:"path,omitempty" form:"path,omitempty"`
	Status string `json:"status,omitempty" bson:"status,omitempty" form:"status,omitempty"`
}

type ImageFilter struct {
	Id      string `json:"id,omitempty" bson:"id,omitempty" form:"id,omitempty"`
	UserId  string `json:"user_id,omitempty" bson:"user_id,omitempty" form:"user_id,omitempty"`
	BlockId string `json:"block_id,omitempty" bson:"block_id,omitempty" form:"block_id,omitempty"`
	Status  string `json:"status,omitempty" bson:"status,omitempty" form:"status,omitempty"`
}

type ImageUpdate struct {
	Data string `json:"data,omitempty" bson:"data,omitempty" form:"data,omitempty"`
}

type ImageLimitResponse struct {
	Data  []ImageResponse `json:"data,omitempty" bson:"data,omitempty" form:"data,omitempty"`
	Total int64           `json:"total,omitempty" bson:"total,omitempty" form:"total,omitempty"`
}
