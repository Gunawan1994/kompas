package log

import (
	"log"
	"net/http"
	"strconv"

	audit "github.com/mentorin/telepenjara/internal/service/log"
	"github.com/mentorin/telepenjara/pkg/errors"
	"github.com/mentorin/telepenjara/pkg/response"
)

// Handler class
type Handler struct {
	logSvc ILogService
}

// ILogService is the abstraction of audit_log service
type ILogService interface {
	GetAllLog(search string, page int, limit int) ([]audit.LogModel, error)
}

// New will create object for class Handler
func New(l ILogService) Handler {
	return Handler{
		logSvc: l,
	}
}

// GetAll audit_log data handler
func (h Handler) GetAll(w http.ResponseWriter, r *http.Request) {
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

	data, err := h.logSvc.GetAllLog(search, pageInt, limitInt)
	if err != nil {
		log.Println(err)
		resp.SetError(errors.New(ErrGet), http.StatusBadRequest)
		return
	}

	resp.Data = data
	resp.Error.Msg = SucGet
	resp.Error.Code = http.StatusOK
	if len(data) != 0 {
		resp.Pagination = Pagination{
			NextPage:     pageInt + 1,
			TotalData:    data[0].TotalData,
			PreviousPage: pageInt - 1,
		}
		return
	}
	resp.Pagination = Pagination{
		TotalData: 0,
	}
	return
}
