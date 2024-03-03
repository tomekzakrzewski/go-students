package models

var (
	minLen = 2
	maxLen = 50

	invalidName    = "name must not be more than 50 and less than 2 characters long"
	invalidSurname = "surname must not be more than 50 and less than 2 characters long"
	invalidAge     = "age must be greater than 0 and less than 150"
)

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
	if len(p.Name) >= maxLen || len(p.Name) < minLen {
		errors["name"] = invalidName
	}
	if p.Surname == "" {
		errors["surname"] = "surname must be provided"
	}
	if len(p.Surname) >= maxLen || len(p.Surname) < minLen {
		errors["surname"] = invalidSurname
	}
	if p.Age <= 0 {
		errors["age"] = invalidAge
	}
	return errors
}

func (p *CreateStudentParams) ValidateUpdate() map[string]string {
	errors := make(map[string]string)
	if p.Name != "" {
		if len(p.Name) >= maxLen || len(p.Name) < minLen {
			errors["name"] = invalidName
		}
	}
	if p.Surname != "" {
		if len(p.Surname) >= maxLen || len(p.Surname) < minLen {
			errors["surname"] = invalidSurname
		}
	}
	if p.Age < 0 {
		errors["age"] = invalidAge
	}
	return errors
}
