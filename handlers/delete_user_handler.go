package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/ftsog/auth/customerrors"
)

func (app *Handler) DeleteUserAccount(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
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

		uID, err := app.Model.GetUserFromSession(w, r)
		if err != nil {
			app.Log.Error.Println(err)
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responsErrorMessage{
					Status:   http.StatusInternalServerError,
					Message:  customerrors.InternalServerError.Error(),
					Path:     r.URL.Path,
					Redirect: fmt.Sprintf("/api/errors/%v", http.StatusInternalServerError),
				},
				StatusCode: http.StatusInternalServerError,
			})

			return
		}

		email, err := app.Model.GetUserEmailByUserID(strings.TrimSpace(*uID))
		if err != nil {
			app.Log.Error.Println(err)
			if err == sql.ErrNoRows {
				JsonResponseWriter(JsonResponse{
					ResponseWriter: w,
					Data: responsErrorMessage{
						Status:  http.StatusBadRequest,
						Message: "Invalid Account",
						Path:    r.URL.Path,
					},
					StatusCode: http.StatusBadRequest,
				})

				return
			}

			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responsErrorMessage{
					Status:   http.StatusInternalServerError,
					Message:  customerrors.InternalServerError.Error(),
					Path:     r.URL.Path,
					Redirect: fmt.Sprintf("/api/errors/%v", http.StatusInternalServerError),
				},
				StatusCode: http.StatusInternalServerError,
			})

			return
		}

		err = app.Model.DeleteSession(w, r)
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

		err = app.Model.DeleteUser(*email)
		if err != nil {
			app.Log.Error.Println(err)
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responsErrorMessage{
					Status:   http.StatusInternalServerError,
					Message:  customerrors.InternalServerError.Error(),
					Path:     r.URL.Path,
					Redirect: fmt.Sprintf("/api/errors/%v", http.StatusInternalServerError),
				},
				StatusCode: http.StatusInternalServerError,
			})

			return
		}

		JsonResponseWriter(JsonResponse{
			ResponseWriter: w,
			Data: responsErrorMessage{
				Status:   http.StatusOK,
				Message:  "Succesfully Deleted Account",
				Path:     r.URL.Path,
				Redirect: "/api/home",
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
