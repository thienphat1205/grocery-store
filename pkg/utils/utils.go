package utils

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
)

func ObjectToByte(object interface{}) []byte {
	data, _ := json.Marshal(object)
	return data
}

func ConvertToBson(object interface{}) (bson.M, error) {
	if object == nil {
		return bson.M{}, nil
	}

	sel, err := bson.Marshal(object)
	if err != nil {
		return nil, err
	}

	obj := bson.M{}
	bson.Unmarshal(sel, &obj)

	return obj, nil
}
