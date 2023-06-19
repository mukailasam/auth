package handlers

import (
	"encoding/json"
	"net/http"
)

type responsErrorMessage struct {
	Status   int
	Message  string
	Path     string
	Redirect string
}

type JsonResponse struct {
	ResponseWriter http.ResponseWriter
	Data           interface{}
	StatusCode     int
}

func JsonResponseWriter(response JsonResponse) {
	response.ResponseWriter.Header().Set("Content-Type", "Application/Json")
	if response.Data != nil {
		rsp, err := json.Marshal(response.Data)
		if err != nil {
			// set response header status code
			SetResponseStatusCode(response.ResponseWriter, http.StatusInternalServerError)

			// response message
			response.ResponseWriter.Write([]byte(http.StatusText(http.StatusInternalServerError)))
			return
		}

		// set response header status code giving by application
		SetResponseStatusCode(response.ResponseWriter, response.StatusCode)

		// response message
		response.ResponseWriter.Write(rsp)
		return
	}

	// set response header status code
	SetResponseStatusCode(response.ResponseWriter, http.StatusInternalServerError)

	// response message
	response.ResponseWriter.Write([]byte(http.StatusText(http.StatusInternalServerError)))
	return
}
