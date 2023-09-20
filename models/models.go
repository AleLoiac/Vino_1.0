package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID              primitive.ObjectID
	First_Name      *string
	Last_Name       *string
	Password        *string
	Email           *string
	Phone           *string
	Token           *string
	Refresh_Token   *string
	Created_At      time.Time
	Updated_At      time.Time
	User_ID         string
	User_Cart       []ProductUser
	Address_Details []Address
	Order_Status    []Order
}

type Product struct {
	Product_ID   primitive.ObjectID
	Product_Name *string
	Price        *uint64
	Rating       *uint8
	Image        *string
}

type ProductUser struct {
	Product_ID   primitive.ObjectID
	Product_Name *string
	Price        *uint64
	Rating       *uint
	Image        *string
}

type Address struct {
	Address_ID
	House
	Street
	City
	Pincode
}

type Order struct {
	Order_ID
	Order_Cart
	Ordered_At
	Price
	Discount
	Payment_Method
}

type Payment struct {
	Digital bool
	COD     bool
}
