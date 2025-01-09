package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func ValidateHmac() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		accCookie := c.Cookies("access_token", "")

		if accCookie == "" {
			return c.Status(503).JSON(fiber.Map{
				"status":  "fail",
				"message": "unauthorization",
			})
		}

		accString := c.Get("Authorization")
		accParse := strings.TrimPrefix(accString, "Bearer ")
		if accCookie != accParse {
			return c.Status(401).JSON(fiber.Map{
				"status":  "fail",
				"message": "unauthorization",
			})
		}

		key := []byte(os.Getenv("SECRET_KEY"))
		result, err := jwt.Parse(accParse, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return key, nil
		})

		if !result.Valid {
			switch {
			case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
				return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
			default:
				return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
			}
		}

		if claims, ok := result.Claims.(jwt.MapClaims); ok && result.Valid {
			userClaims, err := parseClaims(claims)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString(errors.New("[Error] : fail to claim tokan").Error())
			}

			// Serialize userClaims to JSON string
			claimsJSON, err := json.Marshal(userClaims)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString("Failed to serialize claims")
			}

			c.Cookie(&fiber.Cookie{
				Name:     "user_data",
				Value:    string(claimsJSON),
				Expires:  time.Unix(userClaims.Exp, 0),
				HTTPOnly: true,
			})
		} else {
			return c.Status(fiber.StatusInternalServerError).SendString(errors.New("[Error] : fail to claim tokan").Error())
		}

		return c.Next()

	}
}

type UserClaims struct {
	AccountID   string `json:"account_id"`
	Email       string `json:"email"`
	Exp         int64  `json:"exp"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	RoleOffice  string `json:"role_office"`
	RoleWebsite string `json:"role_website"`
}

func parseClaims(claims jwt.MapClaims) (UserClaims, error) {
	userClaims := UserClaims{}
	var ok bool

	// Map values to struct
	if userClaims.AccountID, ok = claims["account_id"].(string); !ok {
		return userClaims, errors.New("invalid account_id")
	}
	if userClaims.Email, ok = claims["email"].(string); !ok {
		return userClaims, errors.New("invalid email")
	}
	if expFloat, ok := claims["exp"].(float64); ok {
		userClaims.Exp = int64(expFloat)
	} else {
		return userClaims, errors.New("invalid exp")
	}
	if userClaims.FirstName, ok = claims["first_name"].(string); !ok {
		return userClaims, errors.New("invalid first_name")
	}
	if userClaims.LastName, ok = claims["last_name"].(string); !ok {
		return userClaims, errors.New("invalid last_name")
	}
	if userClaims.RoleOffice, ok = claims["role_office"].(string); !ok {
		return userClaims, errors.New("invalid role_office")
	}
	if userClaims.RoleWebsite, ok = claims["role_website"].(string); !ok {
		return userClaims, errors.New("invalid role_website")
	}

	return userClaims, nil
}

// func ValidateHmac(roles []models.RoleValidate) func(*fiber.Ctx) error {
// 	return func(c *fiber.Ctx) error {
// 		accCookie := c.Cookies("access_token", "")

// 		if accCookie == "" {
// 			return c.Status(503).JSON(fiber.Map{
// 				"status":  "fail",
// 				"message": "unauthorization",
// 			})
// 		}

// 		accString := c.Get("Authorization")
// 		accParse := strings.TrimPrefix(accString, "Bearer ")
// 		if accCookie != accParse {
// 			return c.Status(401).JSON(fiber.Map{
// 				"status":  "fail",
// 				"message": "unauthorization",
// 			})
// 		}

// 		key := []byte(os.Getenv("SECRET_KEY"))
// 		result, err := jwt.Parse(accParse, func(token *jwt.Token) (interface{}, error) {
// 			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 			}
// 			return key, nil
// 		})

// 		if !result.Valid {
// 			switch {
// 			case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
// 				return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
// 			default:
// 				return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
// 			}
// 		}

// 		roleOfficeId := ""
// 		if claims, ok := result.Claims.(jwt.MapClaims); ok && result.Valid {
// 			roleOfficeId = claims["role_office"].(string)
// 		} else {
// 			return c.Status(fiber.StatusInternalServerError).SendString(errors.New("[Error] : fail to claim tokan").Error())
// 		}

// 		type queryRes struct {
// 			RolePermissionId string                `gorm:"primaryKey;column:role_permission_id;unique;size:50" json:"role_permission_id"`
// 			RoleId           string                `gorm:"column:role_id;size:100" json:"role_id"`
// 			PermissionId     string                `gorm:"column:permission_id" json:"permission_id"`
// 			Permissions      models.PermissionJson `gorm:"column:permissions;type:json" json:"permissions"`
// 			RoleName         string                `gorm:"uniqueIndex;column:role_name;unique;size:100" json:"role_name"`
// 			PermissionName   string                `gorm:"column:permission_name;size:100" json:"permission_name"`
// 			Module           string                `gorm:"column:module;size:100" json:"module"`
// 		}
// 		rolePermission := []queryRes{}

// 		sql := `select rp.*, r.role_name, p.permission_name, p.module from tbl_role_permission as rp
// 			left join tbl_roles  as r on rp.role_id = r.role_id and r.role_id = ?
// 			left join tbl_permissions  as p on rp.permission_id = p.permission_id
// 			where rp.role_id =?`

// 		if err := database.DB.Raw(sql, roleOfficeId, roleOfficeId).Scan(&rolePermission).Error; err != nil {
// 			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 				"status":  "fail",
// 				"message": "[Error] : internal server error",
// 			})
// 		}

// 		respPermission := make([]models.GroupRole, 0)
// 		for _, permission := range rolePermission {
// 			respPermission = append(respPermission, models.GroupRole{
// 				PermissionId:   permission.PermissionId,
// 				PermissionName: permission.PermissionName,
// 				Module:         permission.Module,
// 				Permission:     permission.Permissions,
// 			})
// 		}

// 		role := models.RolePermissionBody{
// 			RoleId:          rolePermission[0].RoleId,
// 			PermissionGroup: respPermission,
// 		}

// 		spew.Dump(roles)

// 		log.Println("--")
// 		roleValid := true
// 		for _, val := range roles {
// 			if val.RoleId == role.RoleId {
// 				for _, perm := range val.Permission {
// 					for _, ele := range role.PermissionGroup {
// 						if perm.PermissionId == ele.PermissionId {
// 							switch perm.Action {
// 							case "c":
// 								if ele.Permission.Create {
// 									return c.Next()
// 								}
// 							case "r":
// 								if ele.Permission.View {
// 									return c.Next()
// 								}
// 							case "u":
// 								if ele.Permission.Edit {
// 									return c.Next()
// 								}
// 							case "d":
// 								if ele.Permission.Delete {
// 									return c.Next()
// 								}
// 							default:
// 								return c.Status(fiber.StatusInternalServerError).SendString(errors.New("[Error] : invalid role").Error())
// 							}
// 						} else {
// 							roleValid = false
// 						}
// 					}
// 				}
// 			} else {
// 				roleValid = false
// 			}
// 		}
// 		if !roleValid {
// 			return c.Status(fiber.StatusInternalServerError).SendString(errors.New("[Error] : invalid role").Error())
// 		} else {
// 			return c.Next()
// 		}
// 	}
// }
