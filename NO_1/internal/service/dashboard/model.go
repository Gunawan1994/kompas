package tahanan

// TahananModel for tahanan data type
type TahananModel struct {
	Judul  string `db:"judul" json:"judul"`
	Body   string `db:"body" json:"body"`
	Author string `db:"author" json:"author"`
}
