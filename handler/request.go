package handler

import (
	"fmt"
)

func errParamRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateOpportunityRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *CreateOpportunityRequest) Validate() error {
	if r.Role == "" && r.Company == "" && r.Location == "" && r.Link == "" && r.Remote == nil && r.Salary <= 0 {
		return fmt.Errorf("request body is empty")
	}

	if r.Role == "" {
		return errParamRequired("role", "string")
	}

	if r.Company == "" {
		return errParamRequired("company", "string")
	}

	if r.Location == "" {
		return errParamRequired("location", "string")
	}

	if r.Link == "" {
		return errParamRequired("link", "string")
	}

	if r.Remote == nil {
		return errParamRequired("remote", "bool")
	}

	if r.Salary <= 0 {
		return errParamRequired("salary", "int64")
	}

	return nil
}

type UpdateOpportunityRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *UpdateOpportunityRequest) Validate() error {
	// If any field is provider, validation is truthy
	if r.Role != "" || r.Company != "" || r.Location != "" || r.Link != "" || r.Remote != nil || r.Salary > 0 {
		return nil
	}

	// If none of the fields are provided, return error
	return fmt.Errorf("at least one valid field is required")
}
