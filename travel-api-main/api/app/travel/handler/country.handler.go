package handler

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/travel-api-main/api/app/travel/services"
	"github.com/travel-api-main/api/app/travel/validation"
	"github.com/travel-api-main/pkg/pb"
	"google.golang.org/grpc/status"
)

var countryService services.CountryService

func InitCountry(service services.CountryService) {
	countryService = service
}

func GetCountry(ctx *fiber.Ctx) error {

	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"GETS_COUNTRY": fiber.Map{"status": 400, "data": "ID tidak ditemukan"}})
	}

	res, err := countryService.GetCountry(context.Background(), &pb.GetCountryRequest{
		GET_COUNTRY: &pb.GetCountryRequestData{
			Id: id,
		},
	})

	if err != nil {
		errStatus, _ := status.FromError(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"GETS_COUNTRY": fiber.Map{"status": 400, "data": errStatus.Message()}})
	}

	return ctx.Status(fiber.StatusOK).JSON(res)
}

func GetsCountry(ctx *fiber.Ctx) error {

	HEADER := new(validation.CountryHeader)

	if err := ctx.QueryParser(HEADER); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"GETS_COUNTRY": fiber.Map{"status": 400, "data": err}})
	}

	res, err := countryService.GetsCountry(context.Background(), &pb.GetsCountryRequest{
		GETS_COUNTRY: &pb.CountryHeader{
			Page: HEADER.Pagination.Page,
			Size: HEADER.Pagination.PerPage,
			Query: &pb.CountryHeader_Query{
				Id:     0,
				Name:   "",
				Alpha2: "",
				Alpha3: "",
			},
			Sort: &pb.CountryHeader_Sort{
				Id:     0,
				Name:   "",
				Alpha2: "",
				Alpha3: "",
			},
		},
	})

	if err != nil {
		errStatus, _ := status.FromError(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"GETS_COUNTRY": fiber.Map{"status": 400, "data": errStatus.Message()}})
	}

	ctx.Response().Header.Add("Access-Control-Expose-Headers", "X-Total-Count")
	ctx.Response().Header.Add("X-Total-Count", strconv.FormatInt(res.GETS_COUNTRY.TotalData, 10))

	return ctx.Status(fiber.StatusOK).JSON(res)
}

func CreateCountry(ctx *fiber.Ctx) error {

	b := validation.CreateCountryRequestBody{}

	if err := ctx.BodyParser(&b); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"CREATE_COUNTRY": fiber.Map{"status": 400, "data": "Review your input"}})
	}

	res, err := countryService.CreateCountry(context.Background(), &pb.CreateCountryRequest{
		CREATE_COUNTRY: &pb.Country{
			Name:                   b.CREATE_COUNTRY.Name,
			Alpha2:                 b.CREATE_COUNTRY.Alpha2,
			Alpha3:                 b.CREATE_COUNTRY.Alpha3,
			CountryCode:            b.CREATE_COUNTRY.CountryCode,
			Iso31662:               b.CREATE_COUNTRY.Iso31662,
			Region:                 b.CREATE_COUNTRY.Region,
			SubRegion:              b.CREATE_COUNTRY.SubRegion,
			IntermediateRegion:     b.CREATE_COUNTRY.IntermediateRegion,
			RegionCode:             b.CREATE_COUNTRY.RegionCode,
			SubRegionCode:          b.CREATE_COUNTRY.SubRegionCode,
			IntermediateRegionCode: b.CREATE_COUNTRY.IntermediateRegionCode,
		},
	})

	if err != nil {
		errStatus, _ := status.FromError(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"CREATE_COUNTRY": fiber.Map{"status": 400, "data": errStatus.Message()}})
	}

	return ctx.Status(fiber.StatusOK).JSON(res)
}

func UpdateCountry(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"UPDATE_COUNTRY": fiber.Map{"status": 400, "data": "ID tidak ditemukan"}})
	}

	b := validation.UpdateCountryRequestBody{}

	if err := ctx.BodyParser(&b); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"UPDATE_COUNTRY": fiber.Map{"status": 400, "data": "Review your input"}})
	}

	res, err := countryService.UpdateCountry(context.Background(), &pb.UpdateCountryRequest{
		UPDATE_COUNTRY: &pb.Country{
			Id:                     id,
			Name:                   b.UPDATE_COUNTRY.Name,
			Alpha2:                 b.UPDATE_COUNTRY.Alpha2,
			Alpha3:                 b.UPDATE_COUNTRY.Alpha3,
			CountryCode:            b.UPDATE_COUNTRY.CountryCode,
			Iso31662:               b.UPDATE_COUNTRY.Iso31662,
			Region:                 b.UPDATE_COUNTRY.Region,
			SubRegion:              b.UPDATE_COUNTRY.SubRegion,
			IntermediateRegion:     b.UPDATE_COUNTRY.IntermediateRegion,
			RegionCode:             b.UPDATE_COUNTRY.RegionCode,
			SubRegionCode:          b.UPDATE_COUNTRY.SubRegionCode,
			IntermediateRegionCode: b.UPDATE_COUNTRY.IntermediateRegionCode,
		},
	})

	if err != nil {
		errStatus, _ := status.FromError(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"UPDATE_COUNTRY": fiber.Map{"status": 400, "data": errStatus.Message()}})
	}

	return ctx.Status(fiber.StatusOK).JSON(res)
}

func DeleteCountry(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"DELETE_COUNTRY": fiber.Map{"status": 400, "data": "ID tidak ditemukan"}})
	}

	res, err := countryService.DeleteCountry(context.Background(), &pb.DeleteCountryRequest{
		DELETE_COUNTRY: &pb.DeleteCountryRequestData{
			Id: id,
		},
	})

	if err != nil {
		errStatus, _ := status.FromError(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"DELETE_COUNTRY": fiber.Map{"status": 400, "data": errStatus.Message()}})
	}

	return ctx.Status(fiber.StatusOK).JSON(res)
}
