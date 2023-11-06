package handler

import "fmt"

// Function Required Struct
func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateBookRequest struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Reserved    *bool  `json:"reserved"`
	Price       int64  `json:"price"`
}

func (b *CreateBookRequest) Validate() error {
	if b.Title == "" && b.Author == "" && b.Description == "" && b.Reserved == nil && b.Price <= 0 {
		return fmt.Errorf("Request body is empty or malformed")
	}
	if b.Title == "" {
		return errParamIsRequired("title", "string")
	}
	if b.Author == "" {
		return errParamIsRequired("author", "string")
	}
	if b.Description == "" {
		return errParamIsRequired("description", "string")
	}
	if b.Reserved == nil {
		return errParamIsRequired("reserved", "bool")
	}
	if b.Price <= 0 {
		return errParamIsRequired("price", "int64")
	}
	return nil
}
