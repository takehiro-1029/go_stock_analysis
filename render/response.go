package render

import (
	"encoding/json"
	"net/http"
)

const (
	contentTypeHeaderKey = "Content-Type"
	contentType          = "application/json; charset=utf-8"
)

type errorResponse struct {
	Message string `json:"message"`
}

// ContractErrorResponse 呼び出し規約エラーレスポンスを描画する
func ContractErrorResponse(w http.ResponseWriter, err error) error {
	obj := errorResponse{err.Error()}
	return JSONResponse(w, http.StatusBadRequest, obj)
}

// JSONResponse JSONレスポンスを描画する
func JSONResponse(w http.ResponseWriter, statusCode int, responseObj interface{}) error {
	w.Header().Set(contentTypeHeaderKey, contentType)
	res, err := json.Marshal(responseObj)
	if err != nil {
		return err
	}
	return renderJSONResponse(w, statusCode, res)
}

func renderJSONResponse(w http.ResponseWriter, statusCode int, msg []byte) error {
	w.WriteHeader(statusCode)
	_, err := w.Write(msg)
	return err
}
