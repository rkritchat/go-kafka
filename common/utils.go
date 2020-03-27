package common

import (
	"encoding/json"
	"go-kafka/model"
	"go-kafka/str"
	"net/http"
)

func RespErr(w http.ResponseWriter, err error) {
	w.Header().Set(str.ContentType, str.ApplicationJson)
	json.NewEncoder(w).Encode(initCommonRes("", err))
}

func RespSuccess(w http.ResponseWriter, s string) {
	w.Header().Set(str.ContentType, str.ApplicationJson)
	json.NewEncoder(w).Encode(initCommonRes(s, nil))
}

func initCommonRes(s string, err error) model.Common {
	if err != nil {
		return model.Common{Status: str.Failed, Desc: err.Error()}
	}

	return model.Common{Status: str.Success, Desc: s}
}
