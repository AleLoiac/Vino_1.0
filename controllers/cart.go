package controllers

import (
	"Vino/database"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"time"
)

type Application struct {
	prodCollection *mongo.Collection
	userCollection *mongo.Collection
}

func NewApplication(prodCollection, userCollection *mongo.Collection) *Application {
	return &Application{
		prodCollection: prodCollection,
		userCollection: userCollection,
	}
}

func (app *Application) AddToCart() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Retrieve the 'ID' query parameter from the request URL
		productQueryID := c.Query("id")

		if productQueryID == "" {
			log.Println("product ID is empty")

			// If 'id' is empty, abort the request with a 400 Bad Request
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("product ID is empty"))
			return
		}

		userQueryID := c.Query("userID")

		if userQueryID == "" {
			log.Println("user ID is empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user ID is empty"))
			return
		}

		// Convert the 'productQueryID' (a string) to a MongoDB ObjectID
		productID, err := primitive.ObjectIDFromHex(productQueryID)

		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)

		defer cancel()

		// Call the 'AddProductToCart' function, passing in the relevant data and context
		err = database.AddProductToCart(ctx, app.prodCollection, app.userCollection, productID, userQueryID)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}
		c.IndentedJSON(200, "Successfully added to cart")
	}
}

func (app *Application) RemoveItem() gin.HandlerFunc {

	return func(c *gin.Context) {

		// Retrieve the 'ID' query parameter from the request URL
		productQueryID := c.Query("id")

		if productQueryID == "" {
			log.Println("product ID is empty")

			// If 'id' is empty, abort the request with a 400 Bad Request
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("product ID is empty"))
			return
		}

		userQueryID := c.Query("userID")

		if userQueryID == "" {
			log.Println("user ID is empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user ID is empty"))
			return
		}

		// Convert the 'productQueryID' (a string) to a MongoDB ObjectID
		productID, err := primitive.ObjectIDFromHex(productQueryID)

		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)

		defer cancel()

		err = database.RemoveCartItem(ctx, app.prodCollection, app.userCollection, productID, userQueryID)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}
		c.IndentedJSON(200, "Successfully removed item from cart")

	}
}

func GetItemFromCart() gin.HandlerFunc {

}

func BuyFromCart() gin.HandlerFunc {

}

func (app *Application) InstantBuy() gin.HandlerFunc {

	return func(c *gin.Context) {

		// Retrieve the 'ID' query parameter from the request URL
		productQueryID := c.Query("id")

		if productQueryID == "" {
			log.Println("product ID is empty")

			// If 'id' is empty, abort the request with a 400 Bad Request
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("product ID is empty"))
			return
		}

		userQueryID := c.Query("userID")

		if userQueryID == "" {
			log.Println("user ID is empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user ID is empty"))
			return
		}

		// Convert the 'productQueryID' (a string) to a MongoDB ObjectID
		productID, err := primitive.ObjectIDFromHex(productQueryID)

		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)

		defer cancel()

		err = database.InstantBuy(ctx, app.prodCollection, app.userCollection, productID, userQueryID)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}
		c.IndentedJSON(200, "Successfully bought item instantly")

	}
}
