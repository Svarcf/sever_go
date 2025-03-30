package models

type User struct {
	Id        uint   `gorm:"primarykey"`
	Username  string `gorm:"unique"`
	FirstName string
	LastName  string
	Password  string
	RoleID    uint
	Role      Role `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func NewUser(username string, firstName string, lastName string, password string, roleID uint) *User {
	return &User{Username: username, FirstName: firstName, LastName: lastName, Password: password, RoleID: roleID}
}
