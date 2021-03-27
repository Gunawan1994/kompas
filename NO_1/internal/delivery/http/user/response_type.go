package user

// Pagination struct
type Pagination struct {
	TotalData    int `json:"total_data"`
	NextPage     int `json:"next_page"`
	PreviousPage int `json:"prev_page"`
}

// Profil struct
type Profil struct {
	Id       uint64 `json:"id"`
	Role     string `json:"role"`
	Username string `json:"username"`
}

// Constanta Message
const (
	SucAdd    = "Data berhasil ditambahkan"
	ErrAdd    = "Data gagal ditambahkan"
	SucEdit   = "Data berhasil diubah"
	ErrEdit   = "Data gagal diubah"
	SucDel    = "Data berhasil dihapus"
	ErrDel    = "Data gagal dihapus"
	SucGet    = "Data berhasil ditemukan"
	ErrGet    = "Data tidak ditemukan"
	ErrExt    = "File harus ber-ekstensi JPG"
	SucUpload = "Successfully Uploaded File"
)
