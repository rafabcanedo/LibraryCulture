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

// Update Book
type UpdateBookRequest struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Reserved    *bool  `json:"reserved"`
	Price       int64  `json:"price"`
}

func (b *UpdateBookRequest) Validate() error {
	// If any field is provided, cvalidation is Truthy
	if b.Title != "" || b.Author != "" || b.Description != "" || b.Reserved != nil || b.Price > 0 {
		return nil
	}
	// If none of the fields were provided, return false
	return fmt.Errorf("at least one valid field must be provided")
}

// -- User ---

// User Request
type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Idade int64  `json:"idade"`
	Cargo string `json:"cargo"`
}

func (u *CreateUserRequest) Validate() error {
	if u.Name == "" && u.Email == "" && u.Idade <= 0 && u.Cargo == "" {
		return fmt.Errorf("Request body is empty or malformed")
	}

	if u.Name == "" {
		return errParamIsRequired("name", "string")
	}

	if u.Email == "" {
		return errParamIsRequired("email", "string")
	}

	if u.Idade <= 0 {
		return errParamIsRequired("idade", "int64")
	}

	if u.Cargo == "" {
		return errParamIsRequired("cargo", "string")
	}

	return nil
}

// Update User
type UpdateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Idade int64  `json:"idade"`
	Cargo string `json:"cargo"`
}

func (u *UpdateUserRequest) Validate() error {
	if u.Name != "" || u.Email != "" || u.Idade > 0 || u.Cargo != "" {
		return nil
	}

	return fmt.Errorf("at least one valid field must be provided")
}
