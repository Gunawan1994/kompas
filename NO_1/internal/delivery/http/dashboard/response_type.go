package tahanan

// Pagination struct for tahanan data
type Pagination struct {
	TotalData    int `json:"total_data"`
	NextPage     int `json:"next_page"`
	PreviousPage int `json:"prev_page"`
}

type Chart struct {
	Male   int `json:"Laki-laki"`
	Female int `json:"Perempuan"`
}

// Constanta Message
const (
	SucAdd  = "Data berhasil ditambahkan"
	ErrAdd  = "Data gagal ditambahkan"
	SucEdit = "Data berhasil diubah"
	ErrEdit = "Data gagal diubah"
	SucDel  = "Data berhasil dihapus"
	ErrDel  = "Data gagal dihapus"
	SucGet  = "Data berhasil ditemukan"
	ErrGet  = "Data tidak ditemukan"
)
