package utils

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

// WriteError will write error on HTTP ResponseWriter with status Code
func WriteError(w http.ResponseWriter, status int, err error) {
	data := struct {
		Msg string `json:"msg"`
	}{
		Msg: err.Error(),
	}

	w.WriteHeader(status)
	encoder := json.NewEncoder(w)
	err = encoder.Encode(data)
	if err != nil {
		logrus.Errorf("Error in Json Encoder, %s", err)
	}
}

// WriteSuccessResponse will write data on HTTP ResponseWriter with StatusOK(200)
func WriteSuccessResponse(w http.ResponseWriter, data interface{}) {
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	err := encoder.Encode(data)
	if err != nil {
		logrus.Errorf("Error in Json Encoder, %s", err)
	}
}
