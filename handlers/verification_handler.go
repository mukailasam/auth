package handlers

import (
	"fmt"
	"net/http"

	"database/sql"

	"github.com/ftsog/auth/customerrors"
	"github.com/go-chi/chi"
)

func (app *Handler) VerifyUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		username := chi.URLParam(r, "username")
		token := chi.URLParam(r, "token")

		isVerified, err := app.Model.IsVerified(username)
		if err != nil {
			app.Log.Error.Println(err)
			if err == sql.ErrNoRows {
				JsonResponseWriter(JsonResponse{
					ResponseWriter: w,
					Data: responsErrorMessage{
						Status:   http.StatusBadRequest,
						Message:  "Invalid Link",
						Path:     r.URL.Path,
						Redirect: r.URL.Path,
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

		if *isVerified {
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responsErrorMessage{
					Status:   http.StatusOK,
					Message:  "Already Verified, goto login",
					Path:     r.URL.Path,
					Redirect: "/api/login",
				},
				StatusCode: http.StatusBadRequest,
			})

			return
		}

		verification, err := app.Model.ReadVerifications(username)
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

		if verification.Expired == true {
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responsErrorMessage{
					Status:   http.StatusBadRequest,
					Message:  "Invalid Link",
					Path:     r.URL.Path,
					Redirect: r.URL.Path,
				},
				StatusCode: http.StatusBadRequest,
			})

			return
		}

		if token != verification.Code {
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responsErrorMessage{
					Status:   http.StatusBadRequest,
					Message:  "Invalid Link",
					Path:     r.URL.Path,
					Redirect: r.URL.Path,
				},
				StatusCode: http.StatusBadRequest,
			})

			return
		}

		err = app.Model.VerifyUser(username)
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
				Message:  "Email Successfully Verified",
				Path:     r.URL.Path,
				Redirect: "/api/login",
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
