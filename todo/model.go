package todo

import "time"

type Todo struct {
	title       string
	description string
	createdAt   time.Time
}
