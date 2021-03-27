package login

import (
	"errors"
	"log"
	"net/http"

	"kompas/internal/service/login"
	"kompas/pkg/response"
)

// Handler class
type Handler struct {
	loginSvc ILoginService
}

// ILoginService is the abstraction of payroll service
type ILoginService interface {
	UserLogin(email string, password string) (login.ModelLogin, error)
}

// New will create object for class Handler
func New(p ILoginService) Handler {
	return Handler{
		loginSvc: p,
	}
}

//Login handler
func (h Handler) Login(w http.ResponseWriter, r *http.Request) {
	resp := &response.Response{}
	defer resp.RenderJSON(w, r)
	var (
		email, password string
	)

	email = r.PostFormValue("email")
	password = r.PostFormValue("password")

	data, err := h.loginSvc.UserLogin(email, password)
	if err != nil {
		log.Println(err)
		resp.SetError(errors.New(FailedLogin), http.StatusBadRequest)
		return
	}

	resp.Data = GetResponse{
		Id:           data.Id,
		Role:         data.Role,
		Username:     data.Username,
		AccessToken:  data.TokenInfo.AccessToken,
		RefreshToken: data.TokenInfo.RefreshToken,
	}
	resp.Error.Msg = SuccessLogin
	resp.Error.Code = 200

	return
}
