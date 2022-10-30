package models

type Event struct {
	ID          string `db:"id"`
	Title       string `db:"title"`
	StartTime   int64  `db:"start_date"`
	EndTime     int64  `db:"end_date"`
	Description string `db:"description,omitempty"`
	OwnerID     int    `db:"owner_id"`
	AlertBefore int64  `db:"alert_before,omitempty"`
}
