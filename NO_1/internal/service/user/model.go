package user

import "time"

// UserModel for user data type
type UserModel struct {
	Nirp        uint64     `db:"nirp" json:"nirp"`
	Username    string     `db:"username" json:"username"`
	Email       string     `db:"email" json:"email"`
	Password    string     `db:"password" json:"password"`
	Role        string     `db:"role" json:"role"`
	IDSatwil    *uint64    `db:"id_satwil" json:"id_satwil"`
	NamaSatwil  *string    `db:"nama_satwil" json:"nama_satwil"`
	createdTime *time.Time `db:"created_time"`
	createdBy   uint64     `db:"created_by"`
	updatedTime time.Time  `db:"updated_time"`
	updatedBy   uint64     `db:"updated_by"`
	LastLogin   *time.Time `db:"last_login" json:"last_login"`
	deleted     bool       `db:"deleted"`
	deletedTime time.Time  `db:"deleted_time"`
	deletedBy   uint64     `db:"deleted_by"`
}

// Operator for user data type
type Operator struct {
	Nirp       uint64 `db:"nirp" json:"nirp"`
	IDSatwil   uint64 `db:"id_satwil" json:"id_satwil"`
	NamaSatwil string `db:"nama_satwil" json:"nama_satwil"`
}
