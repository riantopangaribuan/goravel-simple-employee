package models

import (
	"time"

	"github.com/goravel/framework/database/orm"
)

type Employee struct {
	orm.Model
	UserID        uint   `gorm:"unique;not null"`   // Foreign key ke tabel users, pastikan unik karena 1 user = 1 employee
	User          *User  `gorm:"foreignKey:UserID"` // Relasi BelongsTo ke User
	UserUUID      string `gorm:"unique;not null"`
	EmployeeID    string `gorm:"unique;not null"` // ID Karyawan unik
	Email         string // Redundan jika user_id adalah primary link, tapi tetap ada sesuai permintaan
	FirstName     string `gorm:"not null"`
	LastName      string
	FullName      string     `gorm:"not null"`
	DateOfBirth   *time.Time // Nullable
	PlaceOfBirth  string
	Gender        string `gorm:"type:enum('M','F','U')"` // Male, Female, Unknown
	MaritalStatus string `gorm:"type:enum('0','1','2')"` // 0: Single, 1: Married, 2: Divorce
	Address       string
	PostalCode    string
	ActiveStatus  string `gorm:"type:enum('active','resign');default:'active'"` // Status karyawan
	RoleID        uint   `gorm:"unique;not null"`                               // Foreign key ke tabel roles
	Role          *Role  `gorm:"foreignKey:RoleID"`                             // Relasi BelongsTo ke Role
}

// TableName mengoverride nama tabel default GORM jika berbeda dari 'employees'
func (e *Employee) TableName() string {
	return "employees"
}
