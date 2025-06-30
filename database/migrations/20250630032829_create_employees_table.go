package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250630032829CreateEmployeesTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250630032829CreateEmployeesTable) Signature() string {
	return "20250630032829_create_employees_table"
}

// Up Run the migrations.
func (r *M20250630032829CreateEmployeesTable) Up() error {
	if !facades.Schema().HasTable("employees") {
		return facades.Schema().Create("employees", func(table schema.Blueprint) {
			table.ID()
			table.UnsignedBigInteger("user_id")
			table.Foreign("user_id").References("id").On("users")
			table.String("user_uuid")
			table.String("employee_id")
			table.String("email")
			table.String("first_name")
			table.String("last_name").Nullable()
			table.String("full_name")
			table.Date("date_of_birth").Nullable()
			table.String("place_of_birth").Nullable()
			table.String("gender")
			table.String("marital_status")
			table.String("address").Nullable()
			table.String("postal_code").Nullable()
			table.String("active_status")
			table.TimestampsTz()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250630032829CreateEmployeesTable) Down() error {
	return facades.Schema().DropIfExists("employees")
}
