package validation

import "github.com/go-playground/validator/v10"

type IError struct {
	Field   string
	Tag     string
	Value   string
	Message string
}

type Pagination struct {
	Page    int64 `json:"page"`
	PerPage int64 `json:"perPage"`
}

var Validator = validator.New()

func MsgForTag(tag string, param string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	case "min":
		return "Minimum character is " + param
	}
	return ""
}
