package handlers

import (
	"fmt"
	"net/http"
)

func (app *Handler) Index(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		checker, err := app.Model.CheckSession(w, r)
		if !(*checker) {
			app.Log.Error.Println(err)
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responsErrorMessage{
					Status:   http.StatusBadRequest,
					Message:  "Unauthorize",
					Path:     r.URL.Path,
					Redirect: "api/auth/login",
				},
				StatusCode: http.StatusBadRequest,
			})

			return

		}

		JsonResponseWriter(JsonResponse{
			ResponseWriter: w,
			Data:           "Working...",
			StatusCode:     http.StatusOK,
		})

		return
	}

	JsonResponseWriter(JsonResponse{
		ResponseWriter: w,
		Data: responsErrorMessage{
			Status:   http.StatusMethodNotAllowed,
			Message:  fmt.Sprintf("%s Method not accepted", r.Method),
			Path:     r.URL.Path,
			Redirect: r.URL.Path,
		},
		StatusCode: http.StatusMethodNotAllowed,
	})

	return
}
