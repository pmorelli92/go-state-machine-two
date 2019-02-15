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
			return c.NoContent(http.StatusConflict)
		}

		return c.JSON(http.StatusCreated, ResourceResponse{ID: vehicle.ID()})
	})

	e.GET("/vehicles/:id", func(c echo.Context) error {

		id, err := uuid.FromString(c.Param("id"))
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		vehicle, err := rp.GetByID(id)
		if err != nil {
			return c.NoContent(http.StatusNotFound)
		}

		return c.JSON(http.StatusOK, ToResponseModel(vehicle))
	})

	e.PUT("/vehicles/:id/startRide", func(c echo.Context) error {
		return getVehicleApplyAndPersist(c, rp, func(vehicle *domain.Vehicle, rq *BaseRequest) error {
			return vehicle.StartRide(rq.UserRole)
		})
	})

	e.PUT("/vehicles/:id/finishRide", func(c echo.Context) error {

		return getVehicleApplyAndPersist(c, rp, func(vehicle *domain.Vehicle, rq *BaseRequest) error {
			return vehicle.FinishRide(rq.BatteryLeft, rq.UserRole)
		})
	})

	e.PUT("/vehicles/:id/collect", func(c echo.Context) error {
		return getVehicleApplyAndPersist(c, rp, func(vehicle *domain.Vehicle, rq *BaseRequest) error {
			return vehicle.Collect(rq.UserRole)
		})
	})

	e.PUT("/vehicles/:id/drop", func(c echo.Context) error {
		return getVehicleApplyAndPersist(c, rp, func(vehicle *domain.Vehicle, rq *BaseRequest) error {
			return vehicle.Drop(rq.UserRole)
		})
	})

	e.PUT("/vehicles/:id/ready", func(c echo.Context) error {
		return getVehicleApplyAndPersist(c, rp, func(vehicle *domain.Vehicle, rq *BaseRequest) error {
			return vehicle.Ready(rq.UserRole)
		})
	})

	e.PUT("/vehicles/:id/bounty", func(c echo.Context) error {
		return getVehicleApplyAndPersist(c, rp, func(vehicle *domain.Vehicle, rq *BaseRequest) error {
			return vehicle.SetBounty(rq.UserRole)
		})
	})

	e.PUT("/vehicles/:id/batteryLow", func(c echo.Context) error {
		return getVehicleApplyAndPersist(c, rp, func(vehicle *domain.Vehicle, rq *BaseRequest) error {
			return vehicle.SetBatteryLow(rq.UserRole)
		})
	})

	e.PUT("/vehicles/:id/unknown", func(c echo.Context) error {
		return getVehicleApplyAndPersist(c, rp, func(vehicle *domain.Vehicle, rq *BaseRequest) error {
			return vehicle.Unknown(rq.UserRole)
		})
	})

	e.PUT("/vehicles/setReadyToBounty", func(c echo.Context) error {

		vehicles, err := rp.GetAllWhereReadyState()
		if err != nil {
			return err
		}

		if len(vehicles) == 0 {
			return c.NoContent(http.StatusAccepted)
		}

		if errs := domain.SetVehiclesFromReadyToBounty(vehicles); err != nil {
			return c.JSON(http.StatusForbidden, ToErrorResponseModel(errs))
		}

		var rsp []ResourceResponse

		for _, v := range vehicles {
			err = rp.AddOrUpdate(v)
			if err != nil {
				return c.NoContent(http.StatusConflict)
			}
			rsp = append(rsp, ResourceResponse{ID: v.ID()})
		}

		return c.JSON(http.StatusAccepted, rsp)
	})

	e.PUT("/vehicles/setOldStateToUnknown", func(c echo.Context) error {

		vehicles, err := rp.GetAllWithLastChangeOfStateOlderThanTwoDays()
		if err != nil {
			return err
		}

		if len(vehicles) == 0 {
			return c.NoContent(http.StatusAccepted)
		}

		if errs := domain.SetVehiclesFromReadyToUnknown(vehicles); err != nil {
			return c.JSON(http.StatusForbidden, ToErrorResponseModel(errs))
		}

		var rsp []ResourceResponse

		for _, v := range vehicles {
			err = rp.AddOrUpdate(v)
			if err != nil {
				return c.NoContent(http.StatusConflict)
			}
			rsp = append(rsp, ResourceResponse{ID: v.ID()})
		}

		return c.JSON(http.StatusAccepted, rsp)
	})

	err := e.Start(":8080")
	return err
}

type applyFn func(vehicle *domain.Vehicle, rq *BaseRequest) error

func getVehicleApplyAndPersist(c echo.Context, rp domain.VehicleRepository, applyFn applyFn) error {
	id, err := uuid.FromString(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	u := new(BaseRequest)
	if err = c.Bind(u); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	vehicle, err := rp.GetByID(id)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	if err = applyFn(vehicle, u); err != nil {
		return c.JSON(http.StatusForbidden, ErrorResponse{Message:err.Error()})
	}

	err = rp.AddOrUpdate(vehicle)
	if err != nil {
		return c.NoContent(http.StatusConflict)
	}

	return c.JSON(http.StatusAccepted, ResourceResponse{ID: vehicle.ID()})
}