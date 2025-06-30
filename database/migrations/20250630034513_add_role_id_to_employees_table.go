package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250630034513AddRoleIdToEmployeesTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250630034513AddRoleIdToEmployeesTable) Signature() string {
	return "20250630034513_add_role_id_to_employees_table"
}

// Up Run the migrations.
func (r *M20250630034513AddRoleIdToEmployeesTable) Up() error {
	return facades.Schema().Table("employees", func(table schema.Blueprint) {
		table.UnsignedBigInteger("role_id")
		table.Foreign("role_id").References("id").On("roles")
	})
}

// Down Reverse the migrations.
func (r *M20250630034513AddRoleIdToEmployeesTable) Down() error {
	return facades.Schema().Table("employees", func(table schema.Blueprint) {
		table.DropForeign("role_id")
		table.DropColumn("role_id")
	})
}
