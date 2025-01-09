package routes

import (
	"agent_office/src/database"
	"agent_office/src/models"
)

func getallRoleEndpoint() []models.RoleValidate {
	return []models.RoleValidate{
		{
			RoleId: "ROLE_TEST",
			Permission: []models.PermissionValidate{
				{
					PermissionId: "PERM_AGENT",
					PermissionDetail: database.JsonPermission{
						Create: false,
						View:   true,
						Edit:   true,
						Delete: true,
					},
				},
			},
		},
	}
}
