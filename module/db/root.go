package db

type FindOption struct {
	Limit int64
	Skip  int64
	Sort  interface{}
}

type IDB interface {
	Find(filter interface{}, res interface{}, option FindOption) (int64, error)
	InsertOne(data interface{}) (string, error)
	Update(filter interface{}, data interface{}) (int64, error)
	Delete(image interface{}) (int64, error)
	Distinct(field string, filter interface{}) ([]interface{}, error)
}

func NewDB(dbName string, collectionName string) (IDB, error) {
	var q IDB
	var err error
	q, err = NewMongoRepository(dbName, collectionName)
	return q, err
}
