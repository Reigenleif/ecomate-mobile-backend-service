package models

import (
	"time"

	"github.com/gobuffalo/nulls"
	"github.com/gofrs/uuid"
)

var UserRole []string = []string{"COSTUMER", "ADMIN"}

var Gender = []string{"MALE", "FEMALE"}

type Example struct {
	ID        uuid.UUID `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type User struct {
	ID        nulls.String `db:"id"`

	Name          nulls.String `db:"name"`
	Email         nulls.String `db:"email"`
	EmailVerified CustomTime   `db:"emailverified"`
	HashPassword  nulls.String `db:"hashpassword"`
	ImageUrl      nulls.String `db:"imageurl"`

	Role nulls.String `db:"role"`
}

// News Feature Schemas
type News struct {
	ID        nulls.String `db:"id"`
	CreatedAt time.Time    `db:"createdat"`
	UpdatedAt time.Time    `db:"updatedat"`

	Title       nulls.String `db:"title"`
	ImageUrl    nulls.String `db:"imageurl"`
	Content     nulls.String `db:"content"`
	IsPublished nulls.Bool   `db:"ispublished"`
}

type NewsComment struct {
	ID        uuid.UUID `db:"id"`
	CreatedAt time.Time `db:"createdat"`
	UpdatedAt time.Time `db:"updated_at"`

	Content nulls.String `db:"content"`

	NewsID uuid.UUID `db:"newsid"`
	UserID uuid.UUID `db:"userid"`
}

// Flashcard Feature Schemas

type Flashcard struct {
	ID        uuid.UUID `db:"id"`
	CreatedAt time.Time `db:"createdat"`
	UpdatedAt time.Time `db:"updatedat"`

	Title       nulls.String `db:"title"`
	Description nulls.String `db:"description"`
	ImageUrl    nulls.String `db:"imageurl"`
	IsPublished nulls.Bool   `db:"ispublished"`
}

// Carbon Footprint Calculator Schemas

type CarbonFootprintCalculation struct {
	ID        uuid.UUID `db:"id"`
	CreatedAt time.Time `db:"createdat"`
	UpdatedAt time.Time `db:"updatedat"`

	UserId   uuid.UUID `db:"userid"`
	ItemId   uuid.UUID `db:"itemid"`
	Quantity int       `db:"quantity"`
}

type CarbonFootprintAttribute struct {
	ID        uuid.UUID `db:"id"`
	CreatedAt time.Time `db:"createdat"`
	UpdatedAt time.Time `db:"updatedat"`

	Name        nulls.String  `db:"name"`
	Description nulls.String  `db:"description"`
	Unit        nulls.String  `db:"unit"`
	Coefficient nulls.Float64 `db:"coefficient"`
}

type CarbonFootprintFormula struct {
	ID        uuid.UUID `db:"id"`
	CreatedAt time.Time `db:"createdat"`
	UpdatedAt time.Time `db:"updatedat"`

	Name        nulls.String `db:"name"`
	Description nulls.String `db:"description"`
	Formula     nulls.String `db:"formula"`
	Unit        nulls.String `db:"unit"`
}

// Marketplace Schemas - Item list
type Category struct {
	ID        uuid.UUID `db:"id"`
	CreatedAt time.Time `db:"createdat"`
	UpdatedAt time.Time `db:"updatedat"`

	Name        nulls.String `db:"name"`
	Description nulls.String `db:"description"`
}

type Item struct {
	ID        uuid.UUID `db:"id"`
	CreatedAt time.Time `db:"createdat"`
	UpdatedAt time.Time `db:"updatedat"`

	Name        nulls.String  `db:"name"`
	Description nulls.String  `db:"description"`
	Price       nulls.Float64 `db:"price"`
	ImageUrl    nulls.String  `db:"imageurl"`
	IsPublished nulls.Bool    `db:"ispublished"`

	CategoryId uuid.UUID `db:"categoryid"`
	UserId     uuid.UUID `db:"userid"`
}

// Marketplace Schemas - Cart
type Cart struct {
	ID        uuid.UUID `db:"id"`
	CreatedAt time.Time `db:"createdat"`
	UpdatedAt time.Time `db:"updatedat"`

	UserId uuid.UUID `db:"userid"`
}

type CartItem struct {
	ID        uuid.UUID `db:"id"`
	CreatedAt time.Time `db:"createdat"`
	UpdatedAt time.Time `db:"updatedat"`

	Quantity int `db:"quantity"`

	CartId uuid.UUID `db:"cartid"`
	ItemId uuid.UUID `db:"itemid"`
}
