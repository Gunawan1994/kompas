package dashboardmain

import (
	artikel "kompas/internal/service/dashboard"
	"kompas/pkg/errors"

	"github.com/jmoiron/sqlx"
)

// Resource class
type Resource struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Resource {
	return &Resource{
		db: db,
	}
}

func (r *Resource) GetAll() ([]artikel.DashboardModel, error) {

	dashb := []artikel.DashboardModel{}
	sql := `SELECT 
				judul, body, author
			FROM 
				public.artikel `
	err := r.db.Select(&dashb, sql)
	return dashb, errors.Wrap(err)
}

func (r *Resource) Get(id uint64) (artikel.DashboardModel, error) {
	sql := `SELECT 
				judul, body, author
			FROM 
				public.artikel 
			WHERE 
				id= $1`
	dashb := artikel.DashboardModel{}
	err := r.db.Get(&dashb, sql, id)
	return dashb, errors.Wrap(err)
}

func (r *Resource) Add(judul string, body string, author string) (artikel.DashboardModel, error) {
	sql := `INSERT INTO 
				public.artikel 
				(judul, body, author) 
			VALUES 
				($1, $2, $3) 
			RETURNING 
				judul, body, author`
	dashb := artikel.DashboardModel{}
	err := r.db.QueryRowx(sql, judul, body, author).StructScan(&dashb)
	return dashb, errors.Wrap(err)
}
