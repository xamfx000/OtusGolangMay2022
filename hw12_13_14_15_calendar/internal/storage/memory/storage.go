package memorystorage

import (
	"sync"

	"github.com/xamfx/OtusGolangMay2022/hw12_13_14_15_calendar/internal/models"
	"github.com/xamfx/OtusGolangMay2022/hw12_13_14_15_calendar/internal/pkg"
	"github.com/xamfx/OtusGolangMay2022/hw12_13_14_15_calendar/internal/storage"
)

type Storage struct {
	m  map[string]models.Event
	mu sync.RWMutex
}

func New(m map[string]models.Event) *Storage {
	return &Storage{m: m, mu: sync.RWMutex{}}
}

func (s *Storage) Create(e models.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, v := range s.m {
		if v.OwnerID != e.OwnerID {
			continue
		}
		if pkg.IntervalsIntersect(e.StartTime, e.EndTime, v.StartTime, v.EndTime) {
			return storage.ErrDateBusy
		}
	}
	s.m[e.ID] = e
	return nil
}

func (s *Storage) Update(e models.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, v := range s.m {
		if v.OwnerID != e.OwnerID {
			continue
		}
		if pkg.IntervalsIntersect(e.StartTime, e.EndTime, v.StartTime, v.EndTime) && v.ID != e.ID {
			return storage.ErrDateBusy
		}
	}
	s.m[e.ID] = e
	return nil
}

func (s *Storage) Delete(Id string) error {
	s.mu.Lock()
	delete(s.m, Id)
	s.mu.Unlock()
	return nil
}

func (s *Storage) EventsOnDay(dateTime int64) ([]models.Event, error) {
	s.mu.RLock()
	events := []models.Event{}
	for _, e := range s.m {
		if e.StartTime >= dateTime && e.StartTime <= dateTime+24*3600 {
			events = append(events, e)
		}
	}
	s.mu.RUnlock()
	return events, nil
}

func (s *Storage) EventsOnWeek(startOfWeekDate int64) ([]models.Event, error) {
	s.mu.RLock()
	events := []models.Event{}
	for _, e := range s.m {
		if e.StartTime >= startOfWeekDate && e.StartTime <= startOfWeekDate+7*24*3600 {
			events = append(events, e)
		}
	}
	s.mu.RUnlock()
	return events, nil
}

func (s *Storage) EventsOnMonth(startOfMonthDate int64) ([]models.Event, error) {
	s.mu.RLock()
	events := []models.Event{}
	for _, e := range s.m {
		if e.StartTime >= startOfMonthDate && e.StartTime <= startOfMonthDate+30*24*3600 {
			events = append(events, e)
		}
	}
	s.mu.RUnlock()
	return events, nil
}
