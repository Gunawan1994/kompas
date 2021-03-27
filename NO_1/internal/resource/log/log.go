package log

import (
	"kompas/internal/service/log"
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

// InsertAuditLog insert data to table audit_log
func (r *Resource) InsertAuditLog(nirp uint64, action int, module string) {
	var arr [3]string
	arr[0] = "insert"
	arr[1] = "update"
	arr[2] = "delete"

	sql := `INSERT INTO 
				public.audit_log 
				(nirp, action, module_name) 
			VALUES 
				($1, $2, $3)`
	r.db.Exec(sql, nirp, arr[action], module)
}

// GetAll data
func (r *Resource) GetAll(search string, page int, limit int) ([]log.LogModel, error) {
	page--
	offset := page * limit
	log := []log.LogModel{}
	sql := `SELECT 
				ROW_NUMBER() over(ORDER BY action_date DESC) as nomor,
				id_log, nirp, action, module_name, action_date, 
				count(id_log) over() as total_data
			FROM 
				public.audit_log 
			WHERE 
				(CAST(nirp as TEXT) ILIKE '%' || $1 || '%' 
			OR 
				action ILIKE '%' || $1 || '%' 
			OR 
				module_name ILIKE '%' || $1 || '%')
			LIMIT $2 
			OFFSET $3`
	err := r.db.Select(&log, sql, search, limit, offset)
	return log, errors.Wrap(err)
}
