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

var travelService services.TravelService

func InitTravel(service services.TravelService) {
	travelService = service
}

func GetTravel(ctx *fiber.Ctx) error {

	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"GETS_TRAVEL": fiber.Map{"status": 400, "data": "ID tidak ditemukan"}})
	}

	res, err := travelService.GetTravel(context.Background(), &pb.GetTravelRequest{
		GET_TRAVEL: &pb.GetTravelRequestData{
			Id: id,
		},
	})

	if err != nil {
		errStatus, _ := status.FromError(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"GETS_TRAVEL": fiber.Map{"status": 400, "data": errStatus.Message()}})
	}

	return ctx.Status(fiber.StatusOK).JSON(res)
}

func GetsTravel(ctx *fiber.Ctx) error {

	HEADER := new(validation.TravelHeader)

	if err := ctx.QueryParser(HEADER); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"GETS_TRAVEL": fiber.Map{"status": 400, "data": err}})
	}

	res, err := travelService.GetsTravel(context.Background(), &pb.GetsTravelRequest{
		GETS_TRAVEL: &pb.TravelHeader{
			Page: HEADER.Pagination.Page,
			Size: HEADER.Pagination.PerPage,
			Query: &pb.TravelHeader_Query{
				TravelledCountriesCode: HEADER.Filter.TravelledCountriesCode,
				Search:                 HEADER.Filter.Search,
			},
			Sort: &pb.TravelHeader_Sort{
				Field: HEADER.Sort.Field,
				Order: HEADER.Sort.Order,
			},
		},
	})

	if err != nil {
		errStatus, _ := status.FromError(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"GETS_TRAVEL": fiber.Map{"status": 400, "data": errStatus.Message()}})
	}

	ctx.Response().Header.Add("Access-Control-Expose-Headers", "X-Total-Count")
	ctx.Response().Header.Add("X-Total-Count", strconv.FormatInt(res.GETS_TRAVEL.RecordsTotal, 10))

	return ctx.Status(fiber.StatusOK).JSON(res)
}

func CreateTravel(ctx *fiber.Ctx) error {

	// b := validation.CreateTravelRequestBody{}

	// if err := ctx.BodyParser(&b); err != nil {
	// 	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"CREATE_TRAVEL": fiber.Map{"status": 400, "data": "Review your input"}})
	// }

	res, err := travelService.CreateTravel(context.Background(), &pb.CreateTravelRequest{
		CREATE_TRAVEL: &pb.Travel{
			Id:                     0,
			TravelledCountriesCode: "",
			Total:                  0,
			CreatedAt:              "",
			UpdatedAt:              "",
			Deleted:                "",
		},
	})

	if err != nil {
		errStatus, _ := status.FromError(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"CREATE_TRAVEL": fiber.Map{"status": 400, "data": errStatus.Message()}})
	}

	return ctx.Status(fiber.StatusOK).JSON(res)
}

func UpdateTravel(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"UPDATE_TRAVEL": fiber.Map{"status": 400, "data": "ID tidak ditemukan"}})
	}

	b := validation.UpdateTravelRequestBody{}

	if err := ctx.BodyParser(&b); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"UPDATE_TRAVEL": fiber.Map{"status": 400, "data": "Review your input"}})
	}

	res, err := travelService.UpdateTravel(context.Background(), &pb.UpdateTravelRequest{
		UPDATE_TRAVEL: &pb.Travel{Id: id},
	})

	if err != nil {
		errStatus, _ := status.FromError(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"UPDATE_TRAVEL": fiber.Map{"status": 400, "data": errStatus.Message()}})
	}

	return ctx.Status(fiber.StatusOK).JSON(res)
}

func DeleteTravel(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"DELETE_TRAVEL": fiber.Map{"status": 400, "data": "ID tidak ditemukan"}})
	}

	res, err := travelService.DeleteTravel(context.Background(), &pb.DeleteTravelRequest{
		DELETE_TRAVEL: &pb.DeleteTravelRequestData{
			Id: id,
		},
	})

	if err != nil {
		errStatus, _ := status.FromError(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"DELETE_TRAVEL": fiber.Map{"status": 400, "data": errStatus.Message()}})
	}

	return ctx.Status(fiber.StatusOK).JSON(res)
}
