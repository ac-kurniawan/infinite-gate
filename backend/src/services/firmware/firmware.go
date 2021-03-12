package firmware

import "time"

type Firmware struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
