package database

import (
	"errors"
)

var (
	ErrCantFindProduct    = errors.New("Can't find product")
	ErrCantDecodeProducts = errors.New("Can't find product")
	ErrUserIdIsNotValid   = errors.New("User is not valid")
	ErrCantUpdateUser     = errors.New("Can't add product to cart")
	ErrCantRemoveItemCart = errors.New("Can't remove item from cart")
	ErrCantGetItem        = errors.New("Unable to get item from cart")
	ErrCantBuyCartItem    = errors.New("Can't purchase")
)

func AddProductToCart() {
	return
}

func RemoveCartItem() {
	return
}

func BuyItemFormCart() {
	return
}

func InstantBuyer() {
	return
}
