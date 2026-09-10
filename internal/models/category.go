package models

import "time"

type Category struct {
	ID        int64
	UserID    int64
	Name      string
	CreatedAt time.Time
}
