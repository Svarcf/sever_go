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
