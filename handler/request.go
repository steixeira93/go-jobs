package handler

import "fmt"

type CreateOpeningRequest struct {
	Role string `json:"role"`
	Company string `json:"company"`
	Location string `json:"location"`
	Link string `json:"link"`
	Remote *bool `json:"remote"`
	Salary int64 `json:"salary"`
}

func errParamsisRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" && r.Company == "" && r.Location == "" && r.Link == "" && r.Remote == nil && r.Salary <= 0 {
		return fmt.Errorf("params: request body is empty or malformed")
	}
	if r.Role == "" {
		return errParamsisRequired("role", "string")
	}
	if r.Company == "" {
		return errParamsisRequired("company", "string")
	}
	if r.Location == "" {
		return errParamsisRequired("location", "string")
	}
	if r.Link == "" {
		return errParamsisRequired("link", "string")
	}
	if r.Remote == nil {
		return errParamsisRequired("remote", "bool")
	}
	if r.Salary <= 0 {
		return errParamsisRequired("salary", "int64")
	}
	return nil
}