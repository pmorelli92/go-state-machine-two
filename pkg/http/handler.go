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

	e.PUT("/vehicles/:id/startRide", func(c echo.Context) error {

		id, err := uuid.FromString(c.Param("id"))
		if err != nil {
			return err
		}

		u := new(BaseRequest)
		if err = c.Bind(u); err != nil {
			return err
		}

		vehicle, err := rp.GetById(id)
		if err != nil {
			return err
		}

		if err = vehicle.StartRide(u.UserRole); err != nil {
			return c.JSON(http.StatusForbidden, ErrorResponse{Message:err.Error()})
		}

		err = rp.AddOrUpdate(vehicle)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusAccepted, ResourceResponse{Id:vehicle.Id()})
	})

	e.PUT("/vehicles/:id/finishRide", func(c echo.Context) error {

		id, err := uuid.FromString(c.Param("id"))
		if err != nil {
			return err
		}

		u := new(FinishRideRequest)
		if err = c.Bind(u); err != nil {
			return err
		}

		vehicle, err := rp.GetById(id)
		if err != nil {
			return err
		}

		if err = vehicle.FinishRide(u.BatteryLeft, u.UserRole); err != nil {
			return c.JSON(http.StatusForbidden, ErrorResponse{Message:err.Error()})
		}

		err = rp.AddOrUpdate(vehicle)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusAccepted, ResourceResponse{Id:vehicle.Id()})
	})

	e.PUT("/vehicles/:id/collect", func(c echo.Context) error {

		id, err := uuid.FromString(c.Param("id"))
		if err != nil {
			return err
		}

		u := new(BaseRequest)
		if err = c.Bind(u); err != nil {
			return err
		}

		vehicle, err := rp.GetById(id)
		if err != nil {
			return err
		}

		if err = vehicle.Collect(u.UserRole); err != nil {
			return c.JSON(http.StatusForbidden, ErrorResponse{Message:err.Error()})
		}

		err = rp.AddOrUpdate(vehicle)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusAccepted, ResourceResponse{Id:vehicle.Id()})
	})

	e.PUT("/vehicles/:id/drop", func(c echo.Context) error {

		id, err := uuid.FromString(c.Param("id"))
		if err != nil {
			return err
		}

		u := new(BaseRequest)
		if err = c.Bind(u); err != nil {
			return err
		}

		vehicle, err := rp.GetById(id)
		if err != nil {
			return err
		}

		if err = vehicle.Drop(u.UserRole); err != nil {
			return c.JSON(http.StatusForbidden, ErrorResponse{Message:err.Error()})
		}

		err = rp.AddOrUpdate(vehicle)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusAccepted, ResourceResponse{Id:vehicle.Id()})
	})

	e.PUT("/vehicles/:id/ready", func(c echo.Context) error {

		id, err := uuid.FromString(c.Param("id"))
		if err != nil {
			return err
		}

		u := new(BaseRequest)
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

	e.PUT("/vehicles/:id/bounty", func(c echo.Context) error {

		id, err := uuid.FromString(c.Param("id"))
		if err != nil {
			return err
		}

		u := new(BaseRequest)
		if err = c.Bind(u); err != nil {
			return err
		}

		vehicle, err := rp.GetById(id)
		if err != nil {
			return err
		}

		if err = vehicle.SetBounty(u.UserRole); err != nil {
			return c.JSON(http.StatusForbidden, ErrorResponse{Message:err.Error()})
		}

		err = rp.AddOrUpdate(vehicle)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusAccepted, ResourceResponse{Id:vehicle.Id()})
	})

	e.PUT("/vehicles/:id/batteryLow", func(c echo.Context) error {

		id, err := uuid.FromString(c.Param("id"))
		if err != nil {
			return err
		}

		u := new(BaseRequest)
		if err = c.Bind(u); err != nil {
			return err
		}

		vehicle, err := rp.GetById(id)
		if err != nil {
			return err
		}

		if err = vehicle.SetBatteryLow(u.UserRole); err != nil {
			return c.JSON(http.StatusForbidden, ErrorResponse{Message:err.Error()})
		}

		err = rp.AddOrUpdate(vehicle)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusAccepted, ResourceResponse{Id:vehicle.Id()})
	})

	e.PUT("/vehicles/:id/unknown", func(c echo.Context) error {

		id, err := uuid.FromString(c.Param("id"))
		if err != nil {
			return err
		}

		u := new(BaseRequest)
		if err = c.Bind(u); err != nil {
			return err
		}

		vehicle, err := rp.GetById(id)
		if err != nil {
			return err
		}

		if err = vehicle.Unknown(u.UserRole); err != nil {
			return c.JSON(http.StatusForbidden, ErrorResponse{Message:err.Error()})
		}

		err = rp.AddOrUpdate(vehicle)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusAccepted, ResourceResponse{Id:vehicle.Id()})
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
				return err
			}
			rsp = append(rsp, ResourceResponse{Id: v.Id()})
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
				return err
			}
			rsp = append(rsp, ResourceResponse{Id: v.Id()})
		}

		return c.JSON(http.StatusAccepted, rsp)
	})

	err := e.Start(":8080")
	return err
}