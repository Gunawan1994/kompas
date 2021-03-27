package log

import "time"

// LogModel for audit_log data type
type LogModel struct {
	IDLog      uint64    `db:"id_log"`
	Nomor      uint64    `db:"nomor" json:"nomor"`
	Nirp       uint64    `db:"nirp" json:"nirp"`
	Action     string    `db:"action" json:"action"`
	Module     string    `db:"module_name" json:"module_name"`
	ActionDate time.Time `db:"action_date" json:"action_date"`
	TotalData  int       `db:"total_data" json:"total_data"`
}
