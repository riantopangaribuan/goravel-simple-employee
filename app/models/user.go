package models

import (
	"time"

	"github.com/goravel/framework/database/orm"
)

// User adalah model untuk tabel 'users'
type User struct {
	orm.Model
	FullName        string `gorm:"not null"`
	Email           string `gorm:"unique;not null"`
	EmailVerifiedAt *time.Time
	Phone           *string
	Password        string `gorm:"not null"`
	RememberToken   *string
	Employee        *Employee `gorm:"foreignKey:UserID"` // Relasi HasOne ke Employee
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// TableName mengoverride nama tabel default GORM jika berbeda dari 'users'
func (u *User) TableName() string {
	return "users"
}

// GetID mengimplementasikan kontrak User untuk otentikasi Goravel
func (u *User) GetID() uint {
	return u.ID
}

// GetEmail mengimplementasikan kontrak User untuk otentikasi Goravel
func (u *User) GetEmail() string {
	return u.Email
}

// GetPassword mengimplementasikan kontrak User untuk otentikasi Goravel
func (u *User) GetPassword() string {
	return u.Password
}

// GetRememberToken mengimplementasikan kontrak User untuk otentikasi Goravel
func (u *User) GetRememberToken() string {
	if u.RememberToken == nil {
		return ""
	}
	return *u.RememberToken
}

// GetCreatedAt mengimplementasikan kontrak User untuk otentikasi Goravel
func (u *User) GetCreatedAt() time.Time {
	return u.CreatedAt.Local()
}

// GetUpdatedAt mengimplementasikan kontrak User untuk otentikasi Goravel
func (u *User) GetUpdatedAt() time.Time {
	return u.UpdatedAt.Local()
}

// SetRememberToken mengimplementasikan kontrak User untuk otentikasi Goravel
func (u *User) SetRememberToken(token string) {
	u.RememberToken = &token
}
