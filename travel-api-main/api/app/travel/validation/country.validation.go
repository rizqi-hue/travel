package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type CountryData struct {
	Name                   string `json:"name"`
	Alpha2                 string `json:"alpha_2"`
	Alpha3                 string `json:"alpha_3" `
	CountryCode            string `json:"country_code"`
	Iso31662               string `json:"iso_3166_2"`
	Region                 string `json:"region"`
	SubRegion              string `json:"sub_region"`
	IntermediateRegion     string `json:"intermediate_region"`
	RegionCode             string `json:"region_code" `
	SubRegionCode          string `json:"sub_region_code"`
	IntermediateRegionCode string `json:"intermediate_region_code"`
}

type CountryDataUpdate struct {
	Id                     int64  `json:"id"`
	Name                   string `json:"name"`
	Alpha2                 string `json:"alpha_2"`
	Alpha3                 string `json:"alpha_3" `
	CountryCode            string `json:"country_code"`
	Iso31662               string `json:"iso_3166_2"`
	Region                 string `json:"region"`
	SubRegion              string `json:"sub_region"`
	IntermediateRegion     string `json:"intermediate_region"`
	RegionCode             string `json:"region_code" `
	SubRegionCode          string `json:"sub_region_code"`
	IntermediateRegionCode string `json:"intermediate_region_code"`
}

type CreateCountryRequestBody struct {
	CREATE_COUNTRY CountryData
}

type UpdateCountryRequestBody struct {
	UPDATE_COUNTRY CountryDataUpdate
}

type CountryFilter struct {
	Name        string `json:"name"`
	Alpha2      string `json:"alpha_2"`
	Alpha3      string `json:"alpha_3" `
	CountryCode string `json:"country_code"`
}

type CountryHeader struct {
	Pagination Pagination    `query:"pagination" json:"pagination"`
	Filter     CountryFilter `query:"filter" json:"filter"`
}

func ValidateCreateCountry(ctx *fiber.Ctx) error {
	var errors []*IError
	body := new(CreateCountryRequestBody)
	ctx.BodyParser(&body)

	err := Validator.Struct(body)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el IError
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			el.Message = MsgForTag(err.Tag(), err.Param())
			errors = append(errors, &el)
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"CREATE_COUNTRY": fiber.Map{"status": 400, "data": errors}})
	}
	return ctx.Next()
}

func ValidateUpdateCountry(ctx *fiber.Ctx) error {
	var errors []*IError
	body := new(UpdateCountryRequestBody)
	ctx.BodyParser(&body)

	err := Validator.Struct(body)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el IError
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			el.Message = MsgForTag(err.Tag(), err.Param())
			errors = append(errors, &el)
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"UPDATE_COUNTRY": fiber.Map{"status": 400, "data": errors}})
	}
	return ctx.Next()
}
