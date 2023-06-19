package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/ftsog/auth/customerrors"
	"github.com/ftsog/auth/utils"
)

func (app *Handler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		checker, _ := app.Model.CheckSession(w, r)
		if *checker {
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responsErrorMessage{
					Status:   http.StatusBadRequest,
					Message:  "Already Login",
					Path:     r.URL.Path,
					Redirect: "api/home",
				},
				StatusCode: http.StatusBadRequest,
			})

			return
		}

		jr, err := JsonRequestDecoder(r)
		if err != nil {
			app.Log.Error.Println(err)
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responsErrorMessage{
					Status:   http.StatusBadRequest,
					Message:  customerrors.BadRequest.Error(),
					Path:     r.URL.Path,
					Redirect: r.URL.Path,
				},
				StatusCode: http.StatusBadRequest,
			})

			return
		}

		if len(jr.Data) < 0 || len(jr.Data) < 2 || len(jr.Data) > 2 {
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responsErrorMessage{
					Status:   http.StatusBadRequest,
					Message:  customerrors.BadRequest.Error(),
					Path:     r.URL.Path,
					Redirect: r.URL.Path,
				},
				StatusCode: http.StatusBadRequest,
			})

			return
		}

		usr, err1 := GetValue(jr, "username")
		pwd, err2 := GetValue(jr, "password")

		if err1 != nil || err2 != nil {
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responsErrorMessage{
					Status:   http.StatusBadRequest,
					Message:  customerrors.BadRequest.Error(),
					Path:     r.URL.Path,
					Redirect: r.URL.Path,
				},
				StatusCode: http.StatusBadRequest,
			})

			return
		}

		username := strings.TrimSpace(*usr)
		password := strings.TrimSpace(*pwd)

		if username == "" || password == "" {
			app.Log.Error.Println("Emty fields")
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responsErrorMessage{
					Status:   http.StatusBadRequest,
					Message:  customerrors.AllFieldRequired.Error(),
					Path:     r.URL.Path,
					Redirect: r.URL.Path,
				},
				StatusCode: http.StatusBadRequest,
			})

			return
		}

		ud, err := app.Model.GetUserByUsername(username)
		if err != nil {
			app.Log.Error.Println(err)
			if err == sql.ErrNoRows {
				JsonResponseWriter(JsonResponse{
					ResponseWriter: w,
					Data: responsErrorMessage{
						Status:   http.StatusBadRequest,
						Message:  "Incorrect email or password",
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
					Redirect: fmt.Sprintf("/api/errors/%s", http.StatusInternalServerError),
				},
				StatusCode: http.StatusInternalServerError,
			})

			return
		}

		isVerify, err := app.Model.IsVerified(strings.TrimSpace(ud.Username))
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

		if !(*isVerify) {
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responsErrorMessage{
					Status:   http.StatusBadRequest,
					Message:  "Unverified email, verify your email",
					Path:     r.URL.Path,
					Redirect: r.URL.Path,
				},
				StatusCode: http.StatusBadRequest,
			})

			return
		}

		hashedPassword := utils.HashLoginPassword(password, strings.TrimSpace(ud.Salt))
		if hashedPassword != strings.TrimSpace(ud.Password) {
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responsErrorMessage{
					Status:   http.StatusBadRequest,
					Message:  "Incorrect email or password",
					Path:     r.URL.Path,
					Redirect: r.URL.Path,
				},
				StatusCode: http.StatusBadRequest,
			})

			return
		}

		err = app.Model.CreateSession(w, r, ud.UserID)
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
				Message:  "Succesfully Login",
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
