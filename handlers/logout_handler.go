package handlers

import (
	"fmt"
	"net/http"

	"github.com/ftsog/auth/customerrors"
)

func (app *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		err := app.Model.DeleteSession(w, r)
		if err != nil {
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responsErrorMessage{
					Status:   http.StatusInternalServerError,
					Message:  customerrors.InternalServerError.Error(),
					Path:     r.URL.Path,
					Redirect: fmt.Sprintf("/api/errors/%s", http.StatusInternalServerError),
				},
				StatusCode: http.StatusInternalServerError,
			})

			return
		}

		JsonResponseWriter(JsonResponse{
			ResponseWriter: w,
			Data: responsErrorMessage{
				Status:   http.StatusOK,
				Message:  "Succesfully Logout",
				Path:     r.URL.Path,
				Redirect: "/api/auth/login",
			},
			StatusCode: http.StatusOK,
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
