package tahanan

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"kompas/internal/service/dashboard"
	"kompas/pkg/errors"
	"kompas/pkg/response"
)

// Handler class
type Handler struct {
	thnSvc ITahananService
}

// ITahananService is the abstraction of tahanan service
type ITahananService interface {
	GetID(id uint64) (tahanan.TahananModel, error)
	AddTahanan(ctx context.Context, judul string, body string, author string) (tahanan.TahananModel, error)
	EditTahanan(ctx context.Context, id uint64, sel uint64, nama string, umur uint8, jk string, status string) (tahanan.TahananModel, error)
	DelTahanan(ctx context.Context, id uint64) (tahanan.TahananModel, error)
	GetAllTahanan() ([]tahanan.TahananModel, error)
	GetAllHistory(search string, page int, limit int) ([]tahanan.TahananModel, error)
	GetGenderCount(jk string) (int, error)
	AutoComplete(search string) ([]tahanan.TahananModel, error)
}

// New will create object for class Handler
func New(t ITahananService) Handler {
	return Handler{
		thnSvc: t,
	}
}

//Get tahanan data handler
func (h Handler) Get(w http.ResponseWriter, r *http.Request) {
	resp := &response.Response{}
	defer resp.RenderJSON(w, r)

	vars := mux.Vars(r)
	var idTahanan string
	idTahanan = vars["id"]

	u64ID, err := strconv.ParseUint(idTahanan, 10, 64)
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

// Post tahanan data handler
func (h Handler) Post(w http.ResponseWriter, r *http.Request) {
	resp := &response.Response{}
	defer resp.RenderJSON(w, r)

	var (
		judul, body, author string
	)

	// idTahanan = r.PostFormValue("id_tahanan")
	judul = r.PostFormValue("judul")
	body = r.PostFormValue("body")
	author = r.PostFormValue("author")

	// u64IDTahanan, err := strconv.ParseUint(idTahanan, 10, 64)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// u64IDSel, err := strconv.ParseUint(idSel, 10, 64)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// u8Umur, err := strconv.ParseInt(umur, 10, 8)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	data, err := h.thnSvc.AddTahanan(r.Context(), judul, body, author)
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

// Put tahanan data handler
func (h Handler) Put(w http.ResponseWriter, r *http.Request) {
	resp := &response.Response{}
	defer resp.RenderJSON(w, r)

	vars := mux.Vars(r)
	var (
		idTahanan, idSel, nama, umur, jk, status string
	)

	idTahanan = vars["id_tahanan"]
	idSel = r.PostFormValue("id_sel")
	nama = r.PostFormValue("nama")
	umur = r.PostFormValue("umur")
	jk = r.PostFormValue("jenis_kelamin")
	status = r.PostFormValue("status")

	u64IDTahanan, err := strconv.ParseUint(idTahanan, 10, 64)
	if err != nil {
		log.Println(err)
		return
	}
	u64IDSel, err := strconv.ParseUint(idSel, 10, 64)
	if err != nil {
		log.Println(err)
		return
	}
	u8Umur, err := strconv.ParseInt(umur, 10, 8)
	if err != nil {
		log.Println(err)
		return
	}

	data, err := h.thnSvc.EditTahanan(r.Context(), uint64(u64IDTahanan), uint64(u64IDSel), nama, uint8(u8Umur), jk, status)
	if err != nil {
		log.Println(err)
		resp.SetError(errors.New(ErrEdit), http.StatusBadRequest)
		return
	}

	resp.Data = data
	resp.Error.Msg = SucEdit
	resp.Error.Code = http.StatusOK
	return
}

// Del tahanan data handler
func (h Handler) Del(w http.ResponseWriter, r *http.Request) {
	resp := &response.Response{}
	defer resp.RenderJSON(w, r)

	vars := mux.Vars(r)
	var idTahanan string
	idTahanan = vars["id_tahanan"]

	u64ID, err := strconv.ParseUint(idTahanan, 10, 64)
	if err != nil {
		log.Println(err)
		return
	}

	data, err := h.thnSvc.DelTahanan(r.Context(), uint64(u64ID))
	if err != nil {
		log.Println(err)
		resp.SetError(errors.New(ErrDel), http.StatusBadRequest)
		return
	}

	resp.Data = data
	resp.Error.Msg = SucDel
	resp.Error.Code = http.StatusOK
	return
}

// GetAll tahanan data handler
func (h Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	resp := &response.Response{}
	defer resp.RenderJSON(w, r)

	data, err := h.thnSvc.GetAllTahanan()
	if err != nil {
		log.Println(err)
		resp.SetError(errors.New(ErrGet), http.StatusBadRequest)
		return
	}

	resp.Data = data
	resp.Error.Msg = SucGet
	resp.Error.Code = http.StatusOK
	// if len(data) != 0 {
	// 	resp.Pagination = Pagination{
	// 		NextPage:     pageInt + 1,
	// 		TotalData:    data[0].TotalData,
	// 		PreviousPage: pageInt - 1,
	// 	}
	// 	return
	// }
	// resp.Pagination = Pagination{
	// 	TotalData: 0,
	// }
	return
}

// GetHistoryKunjungan tahanan data handler
func (h Handler) GetHistoryKunjungan(w http.ResponseWriter, r *http.Request) {
	resp := &response.Response{}
	defer resp.RenderJSON(w, r)

	var (
		search, limit, page string
	)

	search = r.FormValue("search")
	limit = r.FormValue("limit")
	page = r.FormValue("page")
	if limit == "" {
		limit = "10"
	}
	if page == "" {
		page = "1"
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		log.Println(err)
		return
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		log.Println(err)
		return
	}

	data, err := h.thnSvc.GetAllHistory(search, pageInt, limitInt)
	if err != nil {
		log.Println(err)
		resp.SetError(errors.New(ErrGet), http.StatusBadRequest)
		return
	}

	resp.Data = data
	resp.Error.Msg = SucGet
	resp.Error.Code = http.StatusOK
	// if len(data) != 0 {
	// 	resp.Pagination = Pagination{
	// 		NextPage:     pageInt + 1,
	// 		TotalData:    data[0].TotalData,
	// 		PreviousPage: pageInt - 1,
	// 	}
	// 	return
	// }
	// resp.Pagination = Pagination{
	// 	TotalData: 0,
	// }
	return
}

// GetChart tahanan data handler
func (h Handler) GetChart(w http.ResponseWriter, r *http.Request) {
	resp := &response.Response{}
	defer resp.RenderJSON(w, r)

	male, err := h.thnSvc.GetGenderCount("Laki-laki")
	if err != nil {
		log.Println(err)
		resp.SetError(errors.New(ErrGet), 400)
		return
	}
	female, err := h.thnSvc.GetGenderCount("Perempuan")
	if err != nil {
		log.Println(err)
		resp.SetError(errors.New(ErrGet), 400)
		return
	}

	resp.Data = Chart{
		Male:   male,
		Female: female,
	}
	resp.Error.Msg = SucGet
	return
}

// AutoComplete Tahanan
func (h Handler) AutoComplete(w http.ResponseWriter, r *http.Request) {
	resp := &response.Response{}
	defer resp.RenderJSON(w, r)

	search := r.FormValue("search")

	data, err := h.thnSvc.AutoComplete(search)
	if err != nil {
		log.Println(errors.Wrap(err))
		resp.SetError(errors.New(ErrGet), http.StatusBadRequest)
		return
	}

	resp.Data = data
	resp.Error.Msg = SucGet
	resp.Error.Code = http.StatusOK
	if len(data) == 0 {
		resp.Error.Msg = ErrGet
		return
	}
	return
}
