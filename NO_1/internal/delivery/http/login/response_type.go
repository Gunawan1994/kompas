package login

// GetResponse data
type GetResponse struct {
	Id           uint64 `json:"id"`
	Username     string `json:"username"`
	Role         string `json:"role"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// Constanta Message
const (
	SuccessLogin = "Login Successfully"
	FailedLogin  = "Email atau password salah !"
)
