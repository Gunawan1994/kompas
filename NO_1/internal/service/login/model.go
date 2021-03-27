package login

// ModelLogin to contain login information
type ModelLogin struct {
	Id        uint64 `db:"id"`
	Username  string `db:"username"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	Role      string `db:"role"`
	TokenInfo TokenDetails
}

// TokenDetails Struct
type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AtExpires    int64
	RtExpires    int64
}
