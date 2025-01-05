package models

import (
	"agent_office/src/database"
	"database/sql/driver"
	"encoding/json"
	"time"
)

type PermissionJson struct {
	Create bool `json:"create"`
	View   bool `json:"view"`
	Edit   bool `json:"edit"`
	Delete bool `json:"delete"`
}

func (j *PermissionJson) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, &j)
}

func (j PermissionJson) Value() (driver.Value, error) {
	return json.Marshal(j)
}

type Role struct {
	RoleID      string         `gorm:"primaryKey;column:role_id;unique;size:50" json:"role_id"`
	RoleName    string         `gorm:"uniqueIndex;column:role_name;unique;size:100" json:"role_name"`
	Permissions PermissionJson `gorm:"column:permissions;type:json" json:"permissions"`
	CreatedAt   time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   *time.Time     `gorm:"column:updated_at" json:"updated_at"`
}

func (Role) TableName() string {
	return "roles"
}

type RolePermission struct {
	RolePermissionId string         `gorm:"primaryKey;column:role_permission_id;unique;size:50" json:"role_permission_id"`
	RoleId           string         `gorm:"column:role_id;size:100;uniqueIndex:idx_role_permission" json:"role_id"`
	PermissionId     string         `gorm:"column:permission_id;uniqueIndex:idx_role_permission" json:"permission_id"`
	Permissions      PermissionJson `gorm:"column:permissions;type:json" json:"permissions"`
	CreatedAt        time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt        *time.Time     `gorm:"column:updated_at" json:"updated_at"`
}

func (RolePermission) TableName() string {
	return "tbl_role_permission"
}

type GroupRolePermission struct {
	PermissionId   string                  `json:"permission_id"`
	PermissionName string                  `json:"permission_name"`
	Module         string                  `json:"module"`
	Permission     database.JsonPermission `json:"permissions"`
}

type GroupRole struct {
	PermissionId   string         `json:"permission_id"`
	PermissionName string         `json:"permission_name"`
	Module         string         `json:"module"`
	Permission     PermissionJson `json:"permissions"`
}

type RolePermissionBody struct {
	RolePermissionId string      `json:"role_permission_id"`
	RoleId           string      `json:"role_id"`
	RoleName         string      `json:"role_name"`
	PermissionGroup  []GroupRole `json:"group_permission"`
}
