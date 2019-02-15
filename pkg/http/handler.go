package http

import (
	"github.com/labstack/echo"
	"github.com/pmorelli92/go-state-machine-two/pkg/domain"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

func Bootstrap(rp domain.VehicleRepository) error {

	e := echo.New()

	e.POST("/vehicles", func(c echo.Context) error {

		vehicle := domain.NewVehicle()
		if err := rp.AddOrUpdate(vehicle); err != nil {
			return err
		}

		return c.JSON(http.StatusCreated, ResourceResponse{Id:vehicle.Id()})
	})

	e.GET("/vehicles/:id", func(c echo.Context) error {

		id, err := uuid.FromString(c.Param("id"))
		if err != nil {
			return err
		}

		vehicle, err := rp.GetById(id)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, ToResponseModel(vehicle))
	})

	e.PUT("/vehicles/:id/ready", func(c echo.Context) error {

		id, err := uuid.FromString(c.Param("id"))
		if err != nil {
			return err
		}

		u := new(ReadyRequest)
		if err = c.Bind(u); err != nil {
			return err
		}

		vehicle, err := rp.GetById(id)
		if err != nil {
			return err
		}

		if err = vehicle.Ready(u.UserRole); err != nil {
			return c.JSON(http.StatusForbidden, ErrorResponse{Message:err.Error()})
		}

		err = rp.AddOrUpdate(vehicle)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusAccepted, ResourceResponse{Id:vehicle.Id()})
	})


	err := e.Start(":8080")
	return err
}