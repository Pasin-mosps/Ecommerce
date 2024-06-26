package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id`
	First_Name      *string            `json:"first_name" 			validate:"required,min=2,,max=20"`
	Last_Name       *string            `json:"last_name 			validate:"required,min=5,max=20"`
	Password        *string            `json:"password" 			validate:"required,min=5"`
	Email           *string            `json:"email" 				validate:"email, required"`
	Phone           *string            `json:"phone" 				validate:"required"`
	Token           *string            `json:"token"`
	Refresh_Token   *string            `json:"refresh_token"`
	Created_At      time.Time          `json:"created_at"`
	Updated_At      time.Time          `json:"updated_at"`
	User_ID         *string            `json:"user_id"`
	UserCart        []ProductUser      `json:"usercart" bson:"usercart"`
	Address_Details []Address          `json:"address" bson:"address"`
	Order_Status    []Order            `json:"orders" bson:"orders"`
}

type Product struct {
	Product_ID   primitive.ObjectID `bson:"_id"`
	Product_Name *string            `json:"product_name"`
	Price        *uint              `json:"price"`
	Rating       *uint              `json:"rating"`
	Image        *string            `json:"image"`
}

type ProductUser struct {
	Product_ID   primitive.ObjectID `bson:"_id"`
	Product_Name *string            `json:"product_name" bson:"product_name"`
	Price        *uint              `json:"price" bson:"price"`
	Rating       *uint              `json:"rating" bson:"rating"`
	Image        *string            `json:"image" bson:"image"`
}
type Address struct {
	Address_id primitive.ObjectID `bson:"_id"`
	House      *string            `json:"house_name" bson:"house_name"`
	Street     *string            `json:"stress_name" bson:"stress_name"`
	City       *string            `json:"city_name" bson:"city_name"`
	Pincode    *uint              `json:"pincode" bson:"pincode"`
}

type Order struct {
	Order_ID       primitive.ObjectID `bson:"order_id"`
	Order_Cart     []ProductUser      `json:"order_id" bson:"order_id"`
	Ordered_At     time.Time          `json:"ordered_at bson:"ordered_at"`
	Price          int                `json:"total_price" bson:"total_price"`
	Discount       *uint              `json:"discount" bson:"discount"`
	Payment_Method Payment            `json:"payment_method" bson:"payment_method"`
}

type Payment struct {
	Digital  bool
	COD      bool
	Promtpay bool
}
