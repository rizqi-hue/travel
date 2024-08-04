package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type TravelData struct {
	TravelledCountriesCode string `json:"travelled_countries_code"`
	Total                  int64  `json:"total"`
}

type TravelDataUpdate struct {
	Id                     int64  `json:"id"`
	TravelledCountriesCode string `json:"travelled_countries_code"`
	Total                  int64  `json:"total"`
}

type CreateTravelRequestBody struct {
	CREATE_TRAVEL TravelData
}

type UpdateTravelRequestBody struct {
	UPDATE_TRAVEL TravelDataUpdate
}

type TravelFilter struct {
	TravelledCountriesCode string `json:"travelled_countries_code"`
	Search                 string `json:"search"`
}

type TravelSort struct {
	Field string `json:"field"`
	Order string  `json:"order"`
}

type TravelHeader struct {
	Pagination Pagination   `query:"pagination" json:"pagination"`
	Filter     TravelFilter `query:"filter" json:"filter"`
	Sort       TravelSort   `query:"sort" json:"sort"`
}

func ValidateCreateTravel(ctx *fiber.Ctx) error {
	var errors []*IError
	body := new(CreateTravelRequestBody)
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
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"CREATE_TRAVEL": fiber.Map{"status": 400, "data": errors}})
	}
	return ctx.Next()
}

func ValidateUpdateTravel(ctx *fiber.Ctx) error {
	var errors []*IError
	body := new(UpdateTravelRequestBody)
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
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"UPDATE_TRAVEL": fiber.Map{"status": 400, "data": errors}})
	}
	return ctx.Next()
}
