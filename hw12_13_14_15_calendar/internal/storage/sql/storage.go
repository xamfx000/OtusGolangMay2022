package sqlstorage

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/xamfx/OtusGolangMay2022/hw12_13_14_15_calendar/internal/models"
	"github.com/xamfx/OtusGolangMay2022/hw12_13_14_15_calendar/internal/storage"
)

type Storage struct {
	db *sqlx.DB
}

func New(db sqlx.DB) *Storage {
	return &Storage{
		db: &db,
	}
}

func (s *Storage) Connect(ctx context.Context) error {
	_, err := s.db.Conn(ctx)
	return err
}

func (s *Storage) Close() error {
	return s.db.Close()
}

func (s *Storage) Create(e models.Event) error {
	result, err := s.db.Exec(
		`INSERT INTO events (
            id,
            title,
            start_time,
            end_time,
            description,
        	owner_id,
            alert_before
            ) select $1, $2, $3, $4, $5, $6,$7
            where not exists (
                select 1 from events 
                         where owner_id = $6
                and (DATE ($3), DATE ($4)) OVERLAPS
                (date_from, date_to)
            )`,
		e.ID,
		e.Title,
		e.StartTime,
		e.EndTime,
		e.Description,
		e.OwnerID,
		e.AlertBefore,
	)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return storage.ErrDateBusy
	}
	return err
}

func (s *Storage) Update(e models.Event) error {
	row, err := s.db.Exec(
		`update events 
			set id = $1, 
			    title = $2, 
			    start_time = $3, 
			    end_time = $4, 
			    description = $5, 
			    owner_id = $6,
			    alert_before = $7
            where id = $1 
            and not exists (
                select 1 from events 
                         where owner_id = $6
                        and not id = $1
                and (DATE ($3), DATE ($4)) OVERLAPS
                (date_from, date_to)
            )`,
		e.ID,
		e.Title,
		e.StartTime,
		e.EndTime,
		e.Description,
		e.OwnerID,
		e.AlertBefore,
	)
	if err != nil {
		return err
	}
	rowsAffected, err := row.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return storage.ErrDateBusy
	}
	return err
}

func (s *Storage) Delete(ID string) error {
	_, err := s.db.Exec(
		`delete from events 
			where id = $1`,
		ID,
	)
	return err
}

func (s *Storage) EventsOnDay(date int64) ([]models.Event, error) {
	res := []models.Event{}
	err := s.db.Select(
		&res,
		`select 
    		id,
    		title,
            start_time,
            end_time,
            description,
            owner_id,
            alert_before
            from events
    		where tsrange(start_time, end_time) && tsrange($1, INTERVAL '1 day')`,
		date,
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Storage) EventsOnWeek(startOfWeekDate int64) ([]models.Event, error) {
	res := []models.Event{}
	err := s.db.Select(
		&res,
		`select 
    		id,
    		title,
            start_time,
            end_time,
            description,
            owner_id,
            alert_before
            from events
    		where tsrange(start_time, end_time) && tsrange($1, INTERVAL '1 week')`,
		startOfWeekDate,
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Storage) EventsOnMonth(startOfMonthDate int64) ([]models.Event, error) {
	res := []models.Event{}
	err := s.db.Select(
		&res,
		`select 
    		id,
    		title,
            start_time,
            end_time,
            description,
            owner_id,
            alert_before
            from events
    		where tsrange(start_time, end_time) && tsrange($1, INTERVAL '1 month')`,
		startOfMonthDate,
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}
