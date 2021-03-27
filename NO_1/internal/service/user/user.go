package user

import (
	"context"
	"errors"
	"log"
)

// IUserResource dependency abstraction for user resource
type IUserResource interface {
	Get(nirp uint64) (UserModel, error)
	Add(nirp uint64, username, email, passsword, role string, idSatwil uint64, createdBy uint64) (UserModel, error)
	Edit(nirp uint64, username string, email string, role string, idSatwil uint64, updatedBy uint64) (UserModel, error)
	Del(nirp uint64, deletedBy uint64) (UserModel, error)
	GetAll(search string, page int, limit int) ([]UserModel, error)
	Operator(nirp uint64) (Operator, error)
	GetAllCount(search string) (int, error)
}

// ILogResource interface
type ILogResource interface {
	InsertAuditLog(nirp uint64, action int, module string)
}

// User class
type User struct {
	user IUserResource
	log  ILogResource
}

// New will create the User object
func New(u IUserResource, l ILogResource) *User {
	return &User{
		user: u,
		log:  l,
	}
}

//AddUser create user
func (u *User) AddUser(ctx context.Context, nirp uint64, username string, email string, password string, role string, idSatwil uint64) (UserModel, error) {

	// to get nirp user yang login
	var (
		loggedUser uint64
		ok         bool
	)

	if v := ctx.Value("nirp"); v != nil {
		loggedUser, ok = v.(uint64)
		if !ok {
			err := errors.New("error converting value to uint64")
			return UserModel{}, err
		}
	}

	data, err := u.user.Add(nirp, username, email, password, role, idSatwil, loggedUser)
	if err == nil {
		u.log.InsertAuditLog(loggedUser, 0, "User")
	}
	return data, err
}

//EditUser edit user
func (u *User) EditUser(ctx context.Context, nirp uint64, username string, email string, role string, idSatwil uint64) (UserModel, error) {
	// to get nirp user yang login
	var (
		loggedUser uint64
		ok         bool
	)

	if v := ctx.Value("nirp"); v != nil {
		loggedUser, ok = v.(uint64)
		if !ok {
			err := errors.New("error converting value to uint64")
			log.Println(err, v)
			return UserModel{}, err
		}
	}

	data, err := u.user.Edit(nirp, username, email, role, idSatwil, loggedUser)
	if err == nil {
		u.log.InsertAuditLog(loggedUser, 1, "User")
	}
	return data, err
}

//GetNirp Get nirp
func (u *User) GetNirp(nirp uint64) (UserModel, error) {
	return u.user.Get(nirp)
}

//DelUser delete user
func (u *User) DelUser(ctx context.Context, nirp uint64) (UserModel, error) {
	var (
		loggedUser uint64
		ok         bool
	)

	if v := ctx.Value("nirp"); v != nil {
		loggedUser, ok = v.(uint64)
		if !ok {
			err := errors.New("error converting value to uint64")
			log.Println(err, v)
			return UserModel{}, err
		}
	}

	data, err := u.user.Del(nirp, loggedUser)
	if err == nil {
		u.log.InsertAuditLog(loggedUser, 2, "User")
	}
	return data, err
}

// GetAllData will return all users
func (u *User) GetAllData(search string, page int, limit int) ([]UserModel, error) {

	return u.user.GetAll(search, page, limit)
}

// GetAllCount will return int
func (u *User) GetAllCount(search string) (int, error) {

	return u.user.GetAllCount(search)
}

//Operator create user
func (u *User) Operator(ctx context.Context) (Operator, error) {

	// to get nirp user yang login
	var (
		loggedUser uint64
		ok         bool
	)

	if v := ctx.Value("nirp"); v != nil {
		loggedUser, ok = v.(uint64)
		if !ok {
			err := errors.New("error converting value to uint64")
			return Operator{}, err
		}
	}

	return u.user.Operator(loggedUser)
}
