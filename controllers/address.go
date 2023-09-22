package controllers

import (
	"Vino/models"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

func AddAddress() gin.HandlerFunc {

}

func EditHomeAddress() gin.HandlerFunc {

}

func EditWorkAddress() gin.HandlerFunc {

}

func DeleteAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Query("id")

		if userID == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"Error": "Invalid search"})
			c.Abort()
			return
		}

		//create an address with value 0
		addresses := make([]models.Address, 0)
		user_id, err := primitive.ObjectIDFromHex(userID)
		if err != nil {
			c.IndentedJSON(500, "Internal Server Error")
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		filter := bson.D{primitive.E{
			Key:   "_id",
			Value: user_id,
		}}

		//update the db with an empty address
		update := bson.D{{Key: "$set", Value: bson.D{primitive.E{
			Key:   "address",
			Value: addresses}}}}

		_, err = UserCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.IndentedJSON(404, "Wrong command")
			return
		}
		defer cancel()
		ctx.Done()
		c.IndentedJSON(200, "Successfully deleted")
	}
}
