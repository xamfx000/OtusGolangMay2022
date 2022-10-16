package memorystorage

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/xamfx/OtusGolangMay2022/hw12_13_14_15_calendar/internal/models"
	storage2 "github.com/xamfx/OtusGolangMay2022/hw12_13_14_15_calendar/internal/storage"
)

func TestStorage_CreateEvent_TimeSlotIsFree(t *testing.T) {
	storage := New(map[string]models.Event{})
	event := models.Event{
		ID:          "lkdsgds",
		Title:       "test event",
		StartTime:   1663525226,
		EndTime:     1663528226,
		Description: "",
		OwnerID:     1,
		AlertBefore: 0,
	}
	err := storage.Create(event)
	require.Equal(t, event, storage.m[event.ID])
	require.NoError(t, err)
}

func TestStorage_CreateEvent_TimeSlotIsBusy(t *testing.T) {
	storage := New(map[string]models.Event{
		"abc": {
			ID:          "abc",
			Title:       "test event",
			StartTime:   1663526226,
			EndTime:     1663529226,
			Description: "",
			OwnerID:     1,
			AlertBefore: 0,
		},
	})
	event := models.Event{
		ID:          "lkdsgds",
		Title:       "test event",
		StartTime:   1663525226,
		EndTime:     1663528226,
		Description: "",
		OwnerID:     1,
		AlertBefore: 0,
	}
	err := storage.Create(event)
	require.Equal(t, models.Event{}, storage.m[event.ID])
	require.ErrorIs(t, err, storage2.ErrDateBusy)
}

func TestStorage_UpdateEvent(t *testing.T) {
	storage := New(map[string]models.Event{
		"abc": {
			ID:          "abc",
			Title:       "test event",
			StartTime:   1663526226,
			EndTime:     1663529226,
			Description: "",
			OwnerID:     1,
			AlertBefore: 0,
		},
		"lkdsgds": {
			ID:          "lkdsgds",
			Title:       "test event",
			StartTime:   1663525226,
			EndTime:     1663528226,
			Description: "",
			OwnerID:     1,
			AlertBefore: 0,
		},
	})
	event := models.Event{
		ID:          "lkdsgds",
		Title:       "test event",
		StartTime:   1663530000,
		EndTime:     1663550000,
		Description: "",
		OwnerID:     1,
		AlertBefore: 0,
	}
	err := storage.Update(event)
	require.Equal(t, event, storage.m[event.ID])
	require.NoError(t, err)
}

func TestStorage_UpdateEvent_ToBusyTime(t *testing.T) {
	storage := New(map[string]models.Event{
		"abc": {
			ID:          "abc",
			Title:       "test event",
			StartTime:   1663526226,
			EndTime:     1663529226,
			Description: "",
			OwnerID:     1,
			AlertBefore: 0,
		},
		"lkdsgds": {
			ID:          "lkdsgds",
			Title:       "test event",
			StartTime:   1663525226,
			EndTime:     1663528226,
			Description: "",
			OwnerID:     1,
			AlertBefore: 0,
		},
	})
	event := models.Event{
		ID:          "lkdsgds",
		Title:       "test event",
		StartTime:   1663526326,
		EndTime:     1663529826,
		Description: "",
		OwnerID:     1,
		AlertBefore: 0,
	}
	err := storage.Update(event)
	require.NotEqual(t, event, storage.m[event.ID])
	require.ErrorIs(t, err, storage2.ErrDateBusy)
}

func TestStorage_DeleteEvent(t *testing.T) {
	storage := New(map[string]models.Event{
		"abc": {
			ID:          "abc",
			Title:       "test event",
			StartTime:   1663526226,
			EndTime:     1663529226,
			Description: "",
			OwnerID:     1,
			AlertBefore: 0,
		},
		"lkdsgds": {
			ID:          "lkdsgds",
			Title:       "test event",
			StartTime:   1663525226,
			EndTime:     1663528226,
			Description: "",
			OwnerID:     1,
			AlertBefore: 0,
		},
	})

	err := storage.Delete("lkdsgds")
	require.Equal(t, models.Event{}, storage.m["lkdsgds"])
	require.Equal(t, 1, len(storage.m))
	require.NoError(t, err)
}
