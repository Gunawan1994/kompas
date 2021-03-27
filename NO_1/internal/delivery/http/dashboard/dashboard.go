package dashboardmain

import (
	"context"
	"log"
	"net/http"
	"strconv"

	dashboard "kompas/internal/service/dashboard"
	"kompas/pkg/errors"
	"kompas/pkg/response"

	"github.com/gorilla/mux"
)

// Handler class
type Handler struct {
	thnSvc IDashboardService
}

type IDashboardService interface {
	GetID(id uint64) (dashboard.DashboardModel, error)
	AddArtikel(ctx context.Context, judul string, body string, author string) (dashboard.DashboardModel, error)
	GetAllArtikel() ([]dashboard.DashboardModel, error)
}

func New(t IDashboardService) Handler {
	return Handler{
		thnSvc: t,
	}
}

func (h Handler) Get(w http.ResponseWriter, r *http.Request) {
	resp := &response.Response{}
	defer resp.RenderJSON(w, r)

	vars := mux.Vars(r)
	var id string
	id = vars["id"]

	u64ID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		log.Println(err)
		return
	}

	data, err := h.thnSvc.GetID(uint64(u64ID))
	if err != nil {
		log.Println(err)
		resp.SetError(errors.New(ErrGet), http.StatusBadRequest)
		return
	}

	resp.Data = data
	resp.Error.Msg = SucGet
	resp.Error.Code = http.StatusOK
	return
}

func (h Handler) Post(w http.ResponseWriter, r *http.Request) {
	resp := &response.Response{}
	defer resp.RenderJSON(w, r)

	var (
		judul, body, author string
	)

	judul = r.PostFormValue("judul")
	body = r.PostFormValue("body")
	author = r.PostFormValue("author")

	data, err := h.thnSvc.AddArtikel(r.Context(), judul, body, author)
	if err != nil {
		log.Println(err)
		resp.SetError(errors.New(ErrAdd), http.StatusBadRequest)
		return
	}

	resp.Data = data
	resp.Error.Msg = SucAdd
	resp.Error.Code = http.StatusOK
	return
}

func (h Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	resp := &response.Response{}
	defer resp.RenderJSON(w, r)

	data, err := h.thnSvc.GetAllArtikel()
	if err != nil {
		log.Println(err)
		resp.SetError(errors.New(ErrGet), http.StatusBadRequest)
		return
	}

	resp.Data = data
	resp.Error.Msg = SucGet
	resp.Error.Code = http.StatusOK
	return
}
