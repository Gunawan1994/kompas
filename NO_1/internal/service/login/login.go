package login

// IUserResource dependency abstraction for user resource
type IUserResource interface {
	Login(email string, password string) (ModelLogin, error)
}

// Login class
type Login struct {
	user IUserResource
}

// New will create the Login object
func New(u IUserResource) *Login {
	return &Login{
		user: u,
	}
}

// UserLogin will return login data by login ID
func (u *Login) UserLogin(email string, pass string) (ModelLogin, error) {
	return u.user.Login(email, pass)
}
