package dashboardmain

type DashboardModel struct {
	Judul  string `db:"judul" json:"judul"`
	Body   string `db:"body" json:"body"`
	Author string `db:"author" json:"author"`
}
