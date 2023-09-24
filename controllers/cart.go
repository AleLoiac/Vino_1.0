package controllers

import (
	"Vino/database"
	"Vino/models"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
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

// GetItemFromCart : we want to retrieve the user first, then the cart associated with it and ungroup it to see the individual values
func GetItemFromCart() gin.HandlerFunc {

	return func(c *gin.Context) {
		user_id := c.Query("id")

		if user_id == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"error": "invalid id"})
			c.Abort()
			return
		}

		usert_id, _ := primitive.ObjectIDFromHex(user_id)

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var filledCart models.User
		err := UserCollection.FindOne(ctx, bson.D{primitive.E{
			Key:   "_id",
			Value: usert_id,
		}}).Decode(&filledCart)

		if err != nil {
			log.Println(err)
			c.IndentedJSON(500, "not found")
			return
		}

		filterMatch := bson.D{{Key: "$match", Value: bson.D{primitive.E{Key: "_id", Value: user_id}}}}
		unwind := bson.D{{Key: "$unwind", Value: bson.D{primitive.E{Key: "path", Value: "$usercart"}}}}
		grouping := bson.D{{Key: "$group", Value: bson.D{primitive.E{Key: "_id", Value: "$_id"}, {Key: "total", Value: bson.D{primitive.E{Key: "$sum", Value: "$usercart.price"}}}}}}

		pointCursor, err := UserCollection.Aggregate(ctx, mongo.Pipeline{filterMatch, unwind, grouping})
		if err != nil {
			log.Println(err)
		}
		var listing []bson.M
		if err = pointCursor.All(ctx, &listing); err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		for _, json := range listing {
			c.IndentedJSON(200, json["total"])
			c.IndentedJSON(200, filledCart.User_Cart)
		}
		ctx.Done()
	}
}

func (app *Application) BuyFromCart() gin.HandlerFunc {

	return func(c *gin.Context) {
		userQueryID := c.Query("id")

		if userQueryID == "" {
			log.Panic("user id is empty")
			c.AbortWithError(http.StatusBadRequest, errors.New("USerID is empty"))
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		defer cancel()

		err := database.BuyItemFromCart(ctx, app.userCollection, userQueryID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}

		c.IndentedJSON(200, "successfully placed the order")
	}
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
		c.IndentedJSON(200, "Successfully placed order")

	}
}
