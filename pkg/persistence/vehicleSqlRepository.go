package persistence

import (
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"github.com/pmorelli92/go-state-machine-two/pkg/domain"
	"github.com/satori/go.uuid"
	"time"
)

type VehicleSQLRepository struct {
	Options PostgresOptions
}

func (rp *VehicleSQLRepository) AddOrUpdate(vehicle *domain.Vehicle) error {

	db, err := sql.Open("postgres", rp.Options.getConnection())
	panicWhenError(err)

	defer db.Close()

	result, err := db.Exec(
		"INSERT INTO go.vehicles(id, battery, current_state, last_change_state) "+
			"VALUES ($1, $2, $3, $4) ON CONFLICT(id) DO UPDATE "+
			"SET battery = excluded.battery, current_state = excluded.current_state, last_change_state = excluded.last_change_state",
		vehicle.ID(), vehicle.Battery(), vehicle.GetCurrentState(), vehicle.LastChangeOfState())

	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("no rows added")
	}

	return nil
}

func (rp *VehicleSQLRepository) GetByID(id uuid.UUID) (*domain.Vehicle, error) {

	db, err := sql.Open("postgres", rp.Options.getConnection())
	panicWhenError(err)

	defer db.Close()

	result := db.QueryRow("SELECT id, battery, current_state, last_change_state FROM go.vehicles WHERE id = $1", id)

	var vID uuid.UUID
	var battery int
	var currentState string
	var lastChangeOfState time.Time

	err = result.Scan(&vID, &battery, &currentState, &lastChangeOfState)
	if err != nil {
		return nil, err
	}

	return domain.RecreateVehicle(vID, battery, lastChangeOfState, currentState), err
}

func (rp *VehicleSQLRepository) GetAllWhereReadyState() ([]*domain.Vehicle, error) {

	db, err := sql.Open("postgres", rp.Options.getConnection())
	panicWhenError(err)

	defer db.Close()

	rows, err := db.Query("SELECT id, battery, current_state, last_change_state FROM go.vehicles WHERE current_state = 'ready'")
	if err != nil {
		return nil, err
	}

	var vehicles []*domain.Vehicle

	for rows.Next() {
		var id uuid.UUID
		var battery int
		var currentState string
		var lastChangeOfState time.Time
		err = rows.Scan(&id, &battery, &currentState, &lastChangeOfState)
		if err != nil {
			return nil, err
		}

		vehicles = append(vehicles, domain.RecreateVehicle(id, battery, lastChangeOfState, currentState))
	}

	return vehicles, nil
}

func (rp *VehicleSQLRepository) GetAllWithLastChangeOfStateOlderThanTwoDays() ([]*domain.Vehicle, error) {

	db, err := sql.Open("postgres", rp.Options.getConnection())
	panicWhenError(err)

	defer db.Close()

	rows, err := db.Query("SELECT id, battery, current_state, last_change_state FROM go.vehicles WHERE now() >= last_change_state + interval '48 hours' AND current_state = 'ready'")
	if err != nil {
		return nil, err
	}

	var vehicles []*domain.Vehicle

	for rows.Next() {
		var id uuid.UUID
		var battery int
		var currentState string
		var lastChangeOfState time.Time
		err = rows.Scan(&id, &battery, &currentState, &lastChangeOfState)
		if err != nil {
			return nil, err
		}

		vehicles = append(vehicles, domain.RecreateVehicle(id, battery, lastChangeOfState, currentState))
	}

	return vehicles, nil
}
