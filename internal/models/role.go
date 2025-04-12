package models

type Role struct {
	Id   uint   `gorm:"primarykey"`
	Name string `gorm:"unique"`
}

type RoleDTO struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

func (r *Role) ToDTO() *RoleDTO {
	return &RoleDTO{
		Id:   r.Id,
		Name: r.Name,
	}
}

func (r *RoleDTO) ToModel() *Role {
	return &Role{
		Id:   r.Id,
		Name: r.Name,
	}
}

func NewRole(name string) *Role {
	return &Role{Name: name}
}
