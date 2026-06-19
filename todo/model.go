package todo

import "time"

type Todo struct {
	Title       string
	Description string
	CreatedAt   time.Time
}
