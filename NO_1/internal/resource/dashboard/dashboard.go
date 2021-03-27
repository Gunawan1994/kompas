package tahanan

import (
	"kompas/internal/service/dashboard"
	"kompas/pkg/errors"

	"github.com/jmoiron/sqlx"
)

// Resource class
type Resource struct {
	db *sqlx.DB
}

// New will return object of Resource class
func New(db *sqlx.DB) *Resource {
	return &Resource{
		db: db,
	}
}

// GetAllHistoryKunjungan exits id_kunjungan in table kunjungan
func (r *Resource) GetAllHistoryKunjungan(search string, page int, limit int) ([]tahanan.TahananModel, error) {
	page--
	offset := page * limit
	sql := `SELECT 
				id_tahanan, id_sel, nama, 
				umur, jenis_kelamin, status, 
				count(id_tahanan) over() as total_data 
			FROM 
				public.tahanan
			WHERE 
				deleted=FALSE 
			AND 
				(EXISTS (
					SELECT 
						NULL
					FROM 
						kunjungan
					WHERE 
						kunjungan.id_tahanan = tahanan.id_tahanan
					AND 
						kunjungan.deleted = FALSE
					AND
						kunjungan.status = 2)
					AND
						(CAST(id_tahanan as TEXT) ILIKE '%' || $3 || '%' 
					OR 
						CAST(id_sel as TEXT) ILIKE '%' || $3 || '%' 
					OR 
						CAST(umur as TEXT) ILIKE '%' || $3 || '%' 
					OR 
						CAST(jenis_kelamin as TEXT) ILIKE '%' || $3 || '%' 
					OR 
						CAST(status as TEXT) ILIKE '%' || $3 || '%' 
					OR 
						nama ILIKE '%' || $3 || '%')) 
			LIMIT $1 
			OFFSET $2`
	tahanan := []tahanan.TahananModel{}
	err := r.db.Select(&tahanan, sql, limit, offset, search)
	return tahanan, errors.Wrap(err)
}

// GetAll data
func (r *Resource) GetAll() ([]tahanan.TahananModel, error) {

	tahanan := []tahanan.TahananModel{}
	sql := `SELECT 
				judul, body, author
			FROM 
				public.artikel `
	err := r.db.Select(&tahanan, sql)
	return tahanan, errors.Wrap(err)
}

// Get id_tahanan
func (r *Resource) Get(id uint64) (tahanan.TahananModel, error) {
	sql := `SELECT 
				judul, body, author
			FROM 
				public.artikel 
			WHERE 
				id= $1`
	tahanan := tahanan.TahananModel{}
	err := r.db.Get(&tahanan, sql, id)
	return tahanan, errors.Wrap(err)
}

// Add data
func (r *Resource) Add(judul string, body string, author string) (tahanan.TahananModel, error) {
	sql := `INSERT INTO 
				public.artikel 
				(judul, body, author) 
			VALUES 
				($1, $2, $3) 
			RETURNING 
				judul, body, author`
	tahanan := tahanan.TahananModel{}
	err := r.db.QueryRowx(sql, judul, body, author).StructScan(&tahanan)
	return tahanan, errors.Wrap(err)
}

// Edit data
func (r *Resource) Edit(id uint64, idSel uint64, nama string, umur uint8, jk string, status string, updatedBy uint64) (tahanan.TahananModel, error) {
	sql := `UPDATE 
				public.tahanan 
			SET 
				id_sel=$1, 
				nama=$2, 
				umur=$3, 
				jenis_kelamin=$4, 
				status=$5, 
				updated_by=$6, 
				updated_time=now() 
			WHERE 
				deleted=FALSE 
			AND 
				id_tahanan=$7 
			RETURNING 
				id_tahanan, id_sel, nama, 
				umur, jenis_kelamin, status`
	tahanan := tahanan.TahananModel{}
	err := r.db.QueryRowx(sql, idSel, nama, umur, jk, status, updatedBy, id).StructScan(&tahanan)
	return tahanan, errors.Wrap(err)
}

// Del data
func (r *Resource) Del(id uint64, deletedBy uint64) (tahanan.TahananModel, error) {
	sql := `UPDATE 
				public.tahanan 
			SET 
				deleted=TRUE, 
				deleted_by=$1, 
				deleted_time=now() 
			WHERE 
				id_tahanan=$2 
			RETURNING 
				id_tahanan, id_sel, nama, 
				umur, jenis_kelamin, status`
	tahanan := tahanan.TahananModel{}
	err := r.db.QueryRowx(sql, deletedBy, id).StructScan(&tahanan)
	return tahanan, errors.Wrap(err)
}

// GetAllCount count all data
func (r *Resource) GetAllCount() (int, error) {
	var count int
	sql := `SELECT 
				COUNT(id_tahanan) AS id_tahanan 
			FROM 
				tahanan 
			WHERE 
				deleted=FALSE`
	err := r.db.Get(&count, sql)
	return count, errors.Wrap(err)
}

// GetGenderCount count data jenis kelamin
func (r *Resource) GetGenderCount(jk string) (int, error) {
	var count int
	sql := `SELECT 
				COUNT(id_tahanan) AS id_tahanan 
			FROM 
				tahanan 
			WHERE 
				deleted=FALSE 
			AND 
				jenis_kelamin=$1`
	err := r.db.Get(&count, sql, jk)
	return count, errors.Wrap(err)
}

// AutoComplete will return ID and Name Tahanan
func (r *Resource) AutoComplete(search string) ([]tahanan.TahananModel, error) {
	sql := `SELECT 
				id_tahanan, nama 
			FROM 
				public.tahanan  
			WHERE 
				deleted=FALSE 
			AND 
				nama ILIKE $1 || '%'
			LIMIT 20` // Limit 20 data
	satwil := []tahanan.TahananModel{}
	err := r.db.Select(&satwil, sql, search)
	return satwil, errors.Wrap(err)
}
