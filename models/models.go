package models

import (
	"time"

	"github.com/gobuffalo/nulls"
	"github.com/gofrs/uuid"
)

type Emoji struct {
	ID        uuid.UUID    `db:"id"`
	Name      nulls.String `db:"name"`
	Image     nulls.String `db:"image"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
}
