package log

// Pagination struct for audit_log data
type Pagination struct {
	TotalData    int `json:"total_data"`
	NextPage     int `json:"next_page"`
	PreviousPage int `json:"prev_page"`
}

// Constanta Message
const (
	SucGet = "Data berhasil ditemukan"
	ErrGet = "Data tidak ditemukan"
)
