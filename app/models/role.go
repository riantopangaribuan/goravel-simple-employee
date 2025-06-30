package models

import (
	"github.com/goravel/framework/database/orm"
)

// Role adalah model untuk tabel 'roles'
type Role struct {
	orm.Model
	Name string `gorm:"unique;not null"` // Nama role: super_admin, admin, employee
}

// TableName mengoverride nama tabel default GORM jika berbeda dari 'roles'
func (r *Role) TableName() string {
	return "roles"
}
