package seeders

import (
	"goravel/app/models"
	"log"

	"github.com/goravel/framework/facades"
	"golang.org/x/crypto/bcrypt"
)

type UserSeeder struct {
}

// Signature The name and signature of the seeder.
func (s *UserSeeder) Signature() string {
	return "UserSeeder"
}

// Run executes the seeder logic.
func (s *UserSeeder) Run() error {
	// Dapatkan ID role super_admin
	var superAdminRole models.Role
	err := facades.Orm().Query().Where("name", "super_admin").First(&superAdminRole)
	if err != nil {
		log.Fatalf("Super admin role not found, please run RoleSeeder first: %v", err)
	}

	// Buat password hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	// Cek apakah super_admin user sudah ada
	var count int64
	err = facades.Orm().Query().Model(&models.User{}).Where("email", "superadmin@example.com").Count(&count)
	if err != nil {
		log.Fatalf("Failed to check if super admin exists: %v", err)
	}

	if count == 0 {
		// Buat user super_admin
		user := models.User{
			FullName: "Super Admin",
			Email:    "superadmin@example.com",
			Password: string(hashedPassword),
		}
		if err := facades.Orm().Query().Create(&user); err != nil {
			log.Fatalf("Failed to create super admin user: %v", err)
		}
		log.Println("Super admin user created successfully.")

		// Buat data employee untuk super admin
		employee := models.Employee{
			UserID:        user.ID,
			EmployeeID:    "SA-001",
			Email:         user.Email,
			FirstName:     "Super",
			LastName:      "Admin",
			FullName:      "Super Admin",
			PlaceOfBirth:  "Jakarta",
			Gender:        "U",
			MaritalStatus: "0",
			Address:       "Jl. Admin No. 1",
			PostalCode:    "10000",
			ActiveStatus:  "active",
			RoleID:        superAdminRole.ID,
		}
		if err := facades.Orm().Query().Create(&employee); err != nil {
			log.Fatalf("Failed to create employee for super admin: %v", err)
		}
		log.Println("Employee data for super admin created successfully.")
	} else {
		log.Println("Super admin user already exists.")
	}

	return nil
}
