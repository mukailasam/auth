package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
)

type JsonRequest struct {
	Data map[string]interface{}
}

func JsonRequestDecoder(r *http.Request) (*JsonRequest, error) {
	jd := &JsonRequest{}
	err := json.NewDecoder(r.Body).Decode(&jd.Data)
	return jd, err
}

func GetValue(jd *JsonRequest, key string) (*string, error) {
	err := errors.New("error")
	value, ok := jd.Data[key]
	if !ok {
		return nil, err
	}

	val, ok := value.(string)
	if !ok {
		return nil, err
	}

	return &val, nil
}
