package seeders

import (
	"goravel/app/models"
	"log"

	"github.com/goravel/framework/facades"
)

type RoleSeeder struct {
}

// Signature The name and signature of the seeder.
func (s *RoleSeeder) Signature() string {
	return "RoleSeeder"
}

// Run executes the seeder logic.
func (s *RoleSeeder) Run() error {
	roles := []string{"super_admin", "admin", "employee"}

	for _, roleName := range roles {
		var count int64
		err := facades.Orm().Query().Model(&models.Role{}).Where("name", roleName).Count(&count)
		if err != nil {
			log.Printf("Error checking role '%s': %v", roleName, err)
			return err
		}

		if count == 0 {
			role := models.Role{Name: roleName}
			if err := facades.Orm().Query().Create(&role); err != nil {
				log.Printf("Failed to create role '%s': %v", roleName, err)
				return err
			}
			log.Printf("Role '%s' created successfully.", roleName)
		} else {
			log.Printf("Role '%s' already exists.", roleName)
		}
	}

	return nil
}
