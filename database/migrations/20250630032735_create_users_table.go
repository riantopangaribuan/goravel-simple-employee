package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250630032735CreateUsersTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250630032735CreateUsersTable) Signature() string {
	return "20250630032735_create_users_table"
}

// Up Run the migrations.
func (r *M20250630032735CreateUsersTable) Up() error {
	if !facades.Schema().HasTable("users") {
		return facades.Schema().Create("users", func(table schema.Blueprint) {
			table.ID()
			table.String("full_name")
			table.String("email")
			table.Unique("email")
			table.Timestamp("email_verified_at")
			table.String("phone").Nullable()
			table.String("password")
			table.String("remember_token").Nullable()
			table.TimestampsTz()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250630032735CreateUsersTable) Down() error {
	return facades.Schema().DropIfExists("users")
}
