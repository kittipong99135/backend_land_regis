package database

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at;type:datetime"`
	UpdatedAt *time.Time     `json:"updated_at" gorm:"column:updated_at;type:datetime"`
	DeletedAt gorm.DeletedAt `gorm:"index;column:deleted_at;type:datetime"`
}

type Account struct {
	BaseModel
	AccountId     string     `json:"account_id" gorm:"primaryKey;column:account_id;type:varchar(10)"`
	Username      string     `json:"username" gorm:"uniqueIndex:idx_username;column:username'type:varchar(50)"`
	Email         string     `json:"email" gorm:"uniqueIndex:idx_email;column:email;type:varchar(50)"`
	Password      string     `json:"password" gorm:"column:password"`
	PhoneNumber   string     `json:"phone_number" gorm:"uniqueIndex:idx_phone_number;column:phone_number;type:varchar(10)"`
	OtpCode       string     `json:"otp_code" gorm:"column:otp_code;type:varchar(10)"`
	OtpExpiry     *time.Time `json:"otp_expiry" gorm:"column:otp_expiry;type:datetime"`
	AzureAdId     *string    `json:"azure_ad_id" gorm:"column:azure_ad_id;type:varchar(50)"`
	AuthType      string     `json:"auth_type" gorm:"column:auth_type;type:varchar(50)"`
	FirstName     string     `json:"firstname" gorm:"column:firstname;type:varchar(50)"`
	LastName      string     `json:"lastname" gorm:"column:lastname;type:varchar(50)"`
	Status        string     `json:"status" gorm:"column:status;type:varchar(20)"`
	RoleOfficeId  string     `gorm:"column:role_office_id;type:varchar(20)" json:"role_office_id"`
	RoleWebsiteId string     `gorm:"column:role_website_id;type:varchar(20)" json:"role_website_id"`
	RoleOffice    Role       `gorm:"foreignKey:RoleOfficeId;references:RoleId" json:"role_office"`
	RoleWebsite   Role       `gorm:"foreignKey:RoleWebsiteId;references:RoleId" json:"role_website"`
}

func (Account) TableName() string {
	return "tbl_accounts"
}

type Role struct {
	BaseModel
	RoleId   string `gorm:"primaryKey;column:role_id;type:varchar(20)" json:"role_id"`
	RoleName string `gorm:"uniqueIndex:idx_role_name;column:role_name;type:varchar(20);unique" json:"role_name"`
	RoleRef  string `gorm:"column:role_ref;type:varchar(20)" json:"role_ref"`
}

func (Role) TableName() string {
	return "tbl_roles"
}

type Permission struct {
	BaseModel
	PermissionId   string `gorm:"primaryKey;column:permission_id;type:varchar(20)" json:"permission_id"`
	PermissionName string `gorm:"uniqueIndex:idx_permission_name;column:permission_name;type:varchar(20);unique" json:"permission_name"`
	Module         string `gorm:"column:module;type:varchar(20)" json:"module"`
}

func (Permission) TableName() string {
	return "tbl_permissions"
}

type JsonPermission struct {
	Create bool `json:"create"`
	View   bool `json:"view"`
	Edit   bool `json:"edit"`
	Delete bool `json:"delete"`
}

func (jp *JsonPermission) Scan(value interface{}) error {
	return json.Unmarshal([]byte(value.(string)), jp)
}

func (jp JsonPermission) Value() (driver.Value, error) {
	return json.Marshal(jp)
}

type RolePermission struct {
	BaseModel
	RolePermissionId string         `gorm:"primaryKey;column:role_permission_id;type:varchar(10)" json:"role_permission_id"`
	RoleRefId        string         `gorm:"column:role_id;type:varchar(20);index:idx_role_permission,unique" json:"role_id"`
	PermissionRefId  string         `gorm:"column:permission_id;type:varchar(20);index:idx_role_permission,unique" json:"permission_id"`
	RoleData         Role           `gorm:"foreignKey:RoleRefId;references:RoleId" json:"role_data"`
	PermissionData   Permission     `gorm:"foreignKey:PermissionRefId;references:PermissionId" json:"permission_data"`
	Permissions      JsonPermission `gorm:"column:permissions;type:json" json:"permissions"`
}

func (RolePermission) TableName() string {
	return "tbl_role_permission"
}
