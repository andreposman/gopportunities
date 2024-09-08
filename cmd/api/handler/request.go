package handler

import "fmt"

type CreateOpeningRequest struct {
	Role        string `json: "role"`
	Description string `json: "description"`
	Company     string `json: "company"`
	IsRemote    *bool  `json: "isRemote"`
	Link        string `json: "link"`
	Salary      int64  `json: "salary"`
}

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

func (c *CreateOpeningRequest) Validate() error {
	if c.Role == "" {
		return errParamIsRequired("role", "string")
	}

	if c.Description == "" {
		return errParamIsRequired("role", "string")
	}

	if c.Company == "" {
		return errParamIsRequired("role", "string")
	}

	if c.Link == "" {
		return errParamIsRequired("role", "string")
	}

	if c.IsRemote == nil {
		return errParamIsRequired("isRemote", "bool")
	}

	if c.Salary <= 0 {
		return errParamIsRequired("salary", "interger")
	}

	return nil
}
