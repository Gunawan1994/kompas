package tahanan

import (
	"context"
	"errors"
	"log"
)

// ITahananResource dependency abstraction for tahanan resource
type ITahananResource interface {
	Get(id uint64) (TahananModel, error)
	Add(judul string, body string, author string) (TahananModel, error)
	Edit(id uint64, sel uint64, nama string, umur uint8, jk string, status string, updatedBy uint64) (TahananModel, error)
	Del(id uint64, deletedBy uint64) (TahananModel, error)
	GetAll() ([]TahananModel, error)
	GetAllHistoryKunjungan(search string, page int, limit int) ([]TahananModel, error)
	GetGenderCount(jk string) (int, error)
	AutoComplete(search string) ([]TahananModel, error)
}

// ILogResource interface
type ILogResource interface {
	InsertAuditLog(nirp uint64, action int, module string)
}

// Tahanan class
type Tahanan struct {
	tahanan ITahananResource
	log     ILogResource
}

// New will create the tahanan object
func New(t ITahananResource, l ILogResource) *Tahanan {
	return &Tahanan{
		tahanan: t,
		log:     l,
	}
}

//AddTahanan create tahanan
func (t *Tahanan) AddTahanan(ctx context.Context, judul string, body string, author string) (TahananModel, error) {
	var (
		loggedUser uint64
		ok         bool
	)

	if v := ctx.Value("nirp"); v != nil {
		loggedUser, ok = v.(uint64)
		if !ok {
			err := errors.New("error converting value to uint64")
			log.Println(err, v)
			return TahananModel{}, err
		}
	}

	data, err := t.tahanan.Add(judul, body, author)
	if err == nil {
		t.log.InsertAuditLog(loggedUser, 0, "Artikel")
	}
	return data, err
}

//EditTahanan edit tahanan
func (t *Tahanan) EditTahanan(ctx context.Context, id uint64, sel uint64, nama string, umur uint8, jk string, status string) (TahananModel, error) {
	var (
		loggedUser uint64
		ok         bool
	)

	if v := ctx.Value("nirp"); v != nil {
		loggedUser, ok = v.(uint64)
		if !ok {
			err := errors.New("error converting value to uint64")
			log.Println(err, v)
			return TahananModel{}, err
		}
	}

	data, err := t.tahanan.Edit(id, sel, nama, umur, jk, status, loggedUser)
	if err == nil {
		t.log.InsertAuditLog(loggedUser, 1, "Tahanan")
	}
	return data, err
}

//GetID Get id_tahanan
func (t *Tahanan) GetID(id uint64) (TahananModel, error) {

	return t.tahanan.Get(id)
}

//DelTahanan delete tahanan
func (t *Tahanan) DelTahanan(ctx context.Context, id uint64) (TahananModel, error) {
	var (
		loggedUser uint64
		ok         bool
	)

	if v := ctx.Value("nirp"); v != nil {
		loggedUser, ok = v.(uint64)
		if !ok {
			err := errors.New("error converting value to uint64")
			log.Println(err, v)
			return TahananModel{}, err
		}
	}

	data, err := t.tahanan.Del(id, loggedUser)
	if err == nil {
		t.log.InsertAuditLog(loggedUser, 2, "Tahanan")
	}
	return data, err
}

// GetAllTahanan will return all tahanan
func (t *Tahanan) GetAllTahanan() ([]TahananModel, error) {

	return t.tahanan.GetAll()
}

// GetAllHistory will return all tahanan
func (t *Tahanan) GetAllHistory(search string, page int, limit int) ([]TahananModel, error) {

	return t.tahanan.GetAllHistoryKunjungan(search, page, limit)
}

// GetGenderCount tahanan
func (t *Tahanan) GetGenderCount(jk string) (int, error) {

	return t.tahanan.GetGenderCount(jk)
}

// AutoComplete tahanan
func (t *Tahanan) AutoComplete(search string) ([]TahananModel, error) {

	return t.tahanan.AutoComplete(search)
}
