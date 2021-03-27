package dashboardmain

import (
	"context"
	"errors"
	"log"
)

type IDashboardService interface {
	Get(id uint64) (DashboardModel, error)
	Add(judul string, body string, author string) (DashboardModel, error)
	GetAll() ([]DashboardModel, error)
}

type ILogResource interface {
	InsertAuditLog(nirp uint64, action int, module string)
}

type Artikel struct {
	artikel IDashboardService
	log     ILogResource
}

func New(t IDashboardService, l ILogResource) *Artikel {
	return &Artikel{
		artikel: t,
		log:     l,
	}
}

func (t *Artikel) AddArtikel(ctx context.Context, judul string, body string, author string) (DashboardModel, error) {
	var (
		loggedUser uint64
		ok         bool
	)

	if v := ctx.Value("nirp"); v != nil {
		loggedUser, ok = v.(uint64)
		if !ok {
			err := errors.New("error converting value to uint64")
			log.Println(err, v)
			return DashboardModel{}, err
		}
	}

	data, err := t.artikel.Add(judul, body, author)
	if err == nil {
		t.log.InsertAuditLog(loggedUser, 0, "Artikel")
	}
	return data, err
}

func (t *Artikel) GetID(id uint64) (DashboardModel, error) {

	return t.artikel.Get(id)
}

func (t *Artikel) GetAllArtikel() ([]DashboardModel, error) {

	return t.artikel.GetAll()
}
