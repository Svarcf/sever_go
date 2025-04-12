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

type UserDTO struct {
	Id        uint    `json:"id"`
	Username  string  `json:"username"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Password  string  `json:"password"`
	Role      RoleDTO `json:"role"`
}

func (u *User) ToDTO() *UserDTO {
	return &UserDTO{
		Id:        u.Id,
		Username:  u.Username,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Password:  u.Password,
		Role:      *u.Role.ToDTO(),
	}
}

func NewUser(username string, firstName string, lastName string, password string, roleID uint) *User {
	return &User{Username: username, FirstName: firstName, LastName: lastName, Password: password, RoleID: roleID}
}
