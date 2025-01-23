package models

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	RoleID   uint   `json:"role_id"` // Role: 1 = Admin, 2 = User
}
