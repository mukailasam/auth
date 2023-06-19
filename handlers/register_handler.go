package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"database/sql"

	"github.com/ftsog/auth/customerrors"
	"github.com/ftsog/auth/utils"
	"github.com/ftsog/auth/validators"
)

func (app *Handler) Register(w http.ResponseWriter, r *http.Request) {
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

		if len(jr.Data) < 0 || len(jr.Data) < 3 || len(jr.Data) > 3 {
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
		eml, err2 := GetValue(jr, "email")
		pwd, err3 := GetValue(jr, "password")

		if err1 != nil || err2 != nil || err3 != nil {
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

		username := strings.TrimSpace(*usr)
		email := strings.TrimSpace(*eml)
		password := strings.TrimSpace(*pwd)

		isEmpty := validators.IsEmpty(username, email, password)
		if isEmpty {
			app.Log.Error.Println("Emty Fields")
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

		userID := utils.GenerateUUID()

		usr, err = validators.ValidateUsername(username)
		if err != nil {
			app.Log.Error.Println(err)
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responsErrorMessage{
					Status:   http.StatusBadRequest,
					Message:  customerrors.InvalidUsername.Error(),
					Path:     r.URL.Path,
					Redirect: r.URL.Path,
				},
				StatusCode: http.StatusBadRequest,
			})

			return
		}

		err = app.Model.UsernameExists(*usr)
		if err == nil {
			app.Log.Error.Println("User Exist")
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responsErrorMessage{
					Status:   http.StatusBadRequest,
					Message:  customerrors.UserExist.Error(),
					Path:     r.URL.Path,
					Redirect: r.URL.Path,
				},
				StatusCode: http.StatusBadRequest,
			})

			return

		} else if err != sql.ErrNoRows {
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

		eml, err = validators.ValidateEmail(email)
		if err != nil {
			app.Log.Error.Println(err)
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responsErrorMessage{
					Status:   http.StatusBadRequest,
					Message:  customerrors.InvalidEmail.Error(),
					Path:     r.URL.Path,
					Redirect: r.URL.Path,
				},
				StatusCode: http.StatusBadRequest,
			})

			return
		}

		err = app.Model.EmailExists(*eml)
		if err == nil {
			app.Log.Error.Println(err)
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responsErrorMessage{
					Status:   http.StatusBadRequest,
					Message:  customerrors.EmailExist.Error(),
					Path:     r.URL.Path,
					Redirect: r.URL.Path,
				},
				StatusCode: http.StatusBadRequest,
			})

			return

		} else if err != sql.ErrNoRows {
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

		pwd, err = validators.ValidatePassword(password)
		if err != nil {
			app.Log.Error.Println(err)
			JsonResponseWriter(JsonResponse{
				ResponseWriter: w,
				Data: responsErrorMessage{
					Status:   http.StatusBadRequest,
					Message:  customerrors.InvalidPassword.Error(),
					Path:     r.URL.Path,
					Redirect: r.URL.Path,
				},
				StatusCode: http.StatusBadRequest,
			})

			return
		}

		verificationID := utils.GenerateUUID()

		emailVerficationCode := utils.Token()

		salt, hashedPassword := utils.HashRegisterPassword(*pwd)
		err = app.Model.CreateUser(userID, verificationID, *usr, *eml, hashedPassword, salt, emailVerficationCode)
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

		subject := "Verify Your Email"
		err = utils.NewMail(*eml, subject, utils.VerifyMessage, *usr, emailVerficationCode)
		if err != nil {
			_ = app.Model.DeleteUser(*eml)
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
				Message:  "Accounted Successfully Created, Account verification link sent to your email",
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
