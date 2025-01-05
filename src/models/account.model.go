package models

import "time"

type Account struct {
	AccountID       string     ` json:"account_id" gorm:"column:account_id;primaryKey"`
	FirstName       string     `json:"firstname" gorm:"column:firstname"`
	LastName        string     `json:"lastname" gorm:"column:lastname"`
	Username        string     `json:"username" gorm:"column:username;uniqueIndex"`
	Email           string     `json:"email" gorm:"column:email;uniqueIndex"`
	PhoneNumber     string     `json:"phone_number" gorm:"column:phone_number"`
	OTPCode         string     `json:"otp_code" gorm:"column:otp_code"`
	OTPExpiry       *time.Time `json:"otp_expiry" gorm:"column:otp_expiry"`
	AzureADID       *string    `json:"azure_ad_id" gorm:"column:azure_ad_id"`
	AuthType        string     `json:"auth_type" gorm:"column:auth_type"`
	RoleID          string     `json:"role_id" gorm:"column:role_id;foreignKey:RoleID;references:RoleID"`
	PasswordHash    *string    `json:"password" gorm:"column:password_hash"`
	PasswordConfirm *string    `json:"password_confirm"`
	Status          string     `json:"status" gorm:"column:status"`
	CreatedAt       time.Time  `json:"create_at" gorm:"column:created_at"`
	UpdatedAt       *time.Time `json:"updete_at" gorm:"column:updated_at"`
}

func (Account) TableName() string {
	return "tbl_accounts"
}

type GetAccount struct {
	AccountID   string ` json:"account_id" gorm:"column:account_id;primaryKey"`
	FirstName   string `json:"firstname" gorm:"column:firstname"`
	LastName    string `json:"lastname" gorm:"column:lastname"`
	Username    string `json:"username" gorm:"column:username;uniqueIndex"`
	Email       string `json:"email" gorm:"column:email;uniqueIndex"`
	PhoneNumber string `json:"phone_number" gorm:"column:phone_number"`
	RoleID      string `json:"role_id" gorm:"column:role_id;foreignKey:RoleID;references:RoleID"`
	Status      string `json:"status" gorm:"column:status"`
}

func (GetAccount) TableName() string {
	return "accounts"
}
