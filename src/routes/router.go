package routes

import (
	c "agent_office/src/controllers"
	"fmt"

	swagger "github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// @title						Land Registeation API Documentation
// @version					0.0.9
// @description				เอกสารการใช้งาน API Land Registeation สำหรับ Developer
// @BasePath					/api/v1
// @schemes					http https
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization

func Start(host string, port string) error {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5000/",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowCredentials: true,
	}))
	app.Use(logger.New())

	swagg := swagger.Config{
		BasePath: "/api/v1",
		FilePath: "./docs/swagger.json",
		Path:     "docs",
		Title:    "Swagger API Docs",
	}
	app.Use(swagger.New(swagg))

	api := app.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			auth := v1.Group("/auth")
			{
				auth.Post("/sign_in", c.SignInEndpoint)
			}
			role := v1.Group("/role")
			{
				role.Post("/", c.CreateRoleEndpoint)        //todo :[swagger]
				role.Get("/", c.GetallRoleEndpoint)         //todo :[swagger]
				role.Get("/:role_id", c.GetoneRoleEndpoint) //todo :[swagger]
				role.Put("/:role_id", c.EditRolePermissionEndpoint)
				role.Put("/:role_id/permission_detail", c.UpdateRolePermissionEndpoint) //todo :[swagger]

			}
			permission := v1.Group("/permission")
			{
				permission.Post("/", c.CreatePermissionEndpoint)           //todo :[swagger]
				permission.Get("/", c.GetallPermissionEndpoint)            //todo :[swagger]
				permission.Get("/:permission_id", c.GetPermissionEndpoint) //todo :[swagger]
				permission.Put("/:permission_id", c.EditPermissionEndpoint)
				permission.Delete("/:permission_id", c.DeletePermissionEndpoint)
			}
			rolePermission := v1.Group("/role_permission")
			{
				rolePermission.Post("/", c.CreateRolePermissionEndpoint)     //todo :[swagger]
				rolePermission.Get("/", c.GetallRolePermissionEndPoint)      //todo :[swagger]
				rolePermission.Get("/:role_id", c.GetRolePermissionEndPoint) //todo :[swagger]
			}
			agent := v1.Group("/agent")
			{
				agent.Post("/", c.CreateAgentEndpoinrt)                 //* :[success]
				agent.Get("/", c.GetallAgentEndpoint)                   //* :[success]
				agent.Get("/:agent_id", c.GetOneAgentEndpoint)          //* :[success]
				agent.Put("/:agent_id", c.EditAgentEndpoint)            //todo :[swagger]
				agent.Put("/:agent_id/role", c.UpdateRoleAgentEndpoint) //* [success]
				agent.Delete("/:agent_id", c.DeleteAgentEndpoint)
				agent.Get("/export", c.DownloadAgentEndpoint)
			}
			layer := v1.Group("/layer")
			{
				layer.Post("/", c.CreateLayerEndpoint)
			}
		}
	}

	return app.Listen(fmt.Sprintf("%s:%s", host, port))
}
