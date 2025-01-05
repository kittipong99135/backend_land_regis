package models

import "time"

type Permission struct {
	PermissionID   string     `gorm:"primaryKey;column:permission_id;size:50" json:"permission_id"` // Unique ID for the permission (Primary Key)
	PermissionName string     `gorm:"column:permission_name;size:100" json:"permission_name"`       // Name of the permission
	Module         string     `gorm:"column:module;size:100" json:"module"`                         // Module or feature associated with the permission
	CreatedAt      time.Time  `gorm:"column:created_at" json:"created_at"`                          // Timestamp when the permission was created
	UpdatedAt      *time.Time `gorm:"column:updated_at" json:"updated_at"`                          // Timestamp when the permission was last updated (Nullable)
}

func (Permission) TableName() string {
	return "permissions"
}
