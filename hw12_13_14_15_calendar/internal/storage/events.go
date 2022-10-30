package storage

import (
	"github.com/pkg/errors"
	"github.com/xamfx/OtusGolangMay2022/hw12_13_14_15_calendar/internal/models"
)

type EventsStorage interface {
	Create(e models.Event) error
	Update(e models.Event) error
	Delete(ID string) error
	EventsOnDay(date int64) ([]models.Event, error)
	EventsOnWeek(startOfWeekDate int64) ([]models.Event, error)
	EventsOnMonth(startOfMonthDate int64) ([]models.Event, error)
}

var ErrDateBusy = errors.New("This time slot is busy by another event")
