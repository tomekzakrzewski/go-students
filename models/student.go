package models

type Student struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int    `json:"age"`
}

type CreateStudentParams struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int    `json:"age"`
}

func NewStudentFromParams(params *CreateStudentParams) *Student {
	return &Student{
		Name:    params.Name,
		Surname: params.Surname,
		Age:     params.Age,
	}
}

func (p *CreateStudentParams) Validate() map[string]string {
	errors := make(map[string]string)
	if p.Name == "" {
		errors["name"] = "name must be provided"
	}
	if len(p.Name) >= 50 {
		errors["name"] = "name must not be more than 50 characters long"
	}
	if p.Surname == "" {
		errors["surname"] = "surname must be provided"
	}
	if len(p.Surname) >= 50 {
		errors["surname"] = "surname must not be more than 50 characters long"
	}
	if p.Age <= 0 {
		errors["age"] = "age must be at least 1"
	}
	return errors
}
