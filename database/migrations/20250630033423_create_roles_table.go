package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250630033423CreateRolesTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250630033423CreateRolesTable) Signature() string {
	return "20250630033423_create_roles_table"
}

// Up Run the migrations.
func (r *M20250630033423CreateRolesTable) Up() error {
	if !facades.Schema().HasTable("roles") {
		return facades.Schema().Create("roles", func(table schema.Blueprint) {
			table.ID()
			table.String("name")
			table.Unique("name")
			table.TimestampsTz()
			table.SoftDeletes()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250630033423CreateRolesTable) Down() error {
	return facades.Schema().DropIfExists("roles")
}
