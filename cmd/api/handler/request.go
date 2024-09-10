package handler

import "fmt"

// TODO! check/fix gorm annotation
type CreateOpeningRequest struct {
	ID          uint   `json:"id" gorm:"unique;primaryKey;autoIncrement=true;not null"`
	Role        string `json: "role"`
	Description string `json: "description"`
	Company     string `json: "company"`
	Location    string `json: "location"`
	IsRemote    *bool  `json: "isRemote"`
	Link        string `json: "link"`
	Salary      int64  `json: "salary"`
}

type UpdateOpeningRequest struct {
	Role        string `json: "role"`
	Description string `json: "description"`
	Company     string `json: "company"`
	Location    string `json: "location"`
	IsRemote    *bool  `json: "isRemote"`
	Link        string `json: "link"`
	Salary      int64  `json: "salary"`
}

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

func (c *CreateOpeningRequest) Validate() error {
	minSalary := int64(5000)

	if c.Role == "" && c.Description == "" && c.Company == "" && c.Location == "" && c.IsRemote == nil && c.Link == "" && c.Salary <= 0 {
		return fmt.Errorf("request body is a invalid json or empty")
	}

	if c.Role == "" {
		return errParamIsRequired("role", "string")
	}

	if c.Description == "" {
		return errParamIsRequired("description", "string")
	}

	if c.Company == "" {
		return errParamIsRequired("company", "string")
	}

	if c.Location == "" {
		return errParamIsRequired("location", "string")
	}

	if c.Link == "" {
		return errParamIsRequired("link", "string")
	}

	if c.IsRemote == nil {
		return errParamIsRequired("isRemote", "bool")
	}

	if c.Salary <= 0 {
		return errParamIsRequired("salary", "interger")
	}
	if minSalary > 0 && c.Salary < minSalary {
		return fmt.Errorf("salary must be greater than %d", minSalary)
	}

	return nil
}

func (u *UpdateOpeningRequest) Validate() error {
	//? if any field is provided, validation is truthy
	if u.Company != "" || u.Description != "" || u.IsRemote != nil || u.Location != "" || u.Role != "" || u.Salary > 0 || u.Link != "" {
		return nil
	}

	return fmt.Errorf("at least one valid field must be provided")
}
