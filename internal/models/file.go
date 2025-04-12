package models

type File struct {
	Id       uint `gorm:"primarykey"`
	Name     string
	Location string
}

type FileDTO struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

func (f *FileDTO) ToModel() *File {
	return &File{
		Id:       f.Id,
		Name:     f.Name,
		Location: f.Location,
	}
}

func (f *File) ToDTO() *FileDTO {
	return &FileDTO{
		Id:       f.Id,
		Name:     f.Name,
		Location: f.Location,
	}
}

func NewFile(name string, location string) *File {
	return &File{Name: name, Location: location}
}
