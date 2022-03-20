// Recipes API
//
// This is a samle recipes API. You can find out more about the API at https://github.com/khhini/go-riset/gin-rest-api
//
// Schemes: http
// Host: localhost:8080
// BasePath: /
// Version: 1.0.0
// Contact: Kiki<kiki.h.hutapea@gmail.com> https://github.com/khhini
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
// swagger:meta
package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/khhini/go-riset/gin-rest-api/handlers"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var recipesHandler *handlers.RecipesHandler

// IndexHandler ...
// swagger:operation POST / index Index
// ---
// produces:
// - application/json
// responses:
//   '200':
//     description: Successful operation
func IndexHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"ping": "pong",
	})
}

// DeleteRecipeHandler ...
// swagger:operation Delete /recipes/{id} recipes deleteRecipe
// Delete existing recipe
// ---
// parameters:
// - name: id
//   in: path
//   description: ID of the recipe
//   required: true
//   type: string
// produces:
// - application/json
// responses:
//   '200':
//      description: Successful operation
//   '404':
//      description: Invalid recipe ID

// func DeleteRecipeHandler(c *gin.Context) {
// 	id := c.Param("id")
// 	objectID, _ := primitive.ObjectIDFromHex(id)
// 	index := -1
// 	for i := 0; i < len(recipes); i++ {
// 		if recipes[i].ID == objectID {
// 			index = i
// 		}
// 	}
// 	if index == -1 {
// 		c.IndentedJSON(http.StatusNotFound, gin.H{
// 			"error": "Recipe not found",
// 		})
// 		return
// 	}

// 	recipes = append(recipes[:index], recipes[index+1:]...)
// 	c.IndentedJSON(http.StatusOK, gin.H{
// 		"message": "Recipe has been deleted",
// 	})
// }

// SearchRecipesHandler ...
// swagger:operation GET /recipes/search recipes searchRecipe
// Return recipe by its tag
// ---
// parameters:
// - name: tag
//   in: query
//   description: recipe tag
//   required: false
//   type: string
//
// produces:
// - application/json
// responses:
//   '200':
//      description: Successful operation

// func SearchRecipesHandler(c *gin.Context) {
// 	tag := c.Query("tag")
// 	listOfRecipes := make([]Recipe, 0)

// 	for i := 0; i < len(recipes); i++ {
// 		found := false
// 		for _, t := range recipes[i].Tags {
// 			if strings.EqualFold(t, tag) {
// 				found = true
// 			}
// 		}
// 		if found {
// 			listOfRecipes = append(listOfRecipes, recipes[i])
// 		}
// 	}

// 	c.IndentedJSON(http.StatusOK, listOfRecipes)
// }

func init() {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection("recipes")
	recipesHandler = handlers.NewRecipeHandler(ctx, collection)

	// var listOfRecipes []interface{}

	// for _, recipe := range recipes {
	// 	listOfRecipes = append(listOfRecipes, recipe)
	// }

	// insertManyResult, err := collection.InsertMany(ctx, listOfRecipes)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("inserted recipes: ", len(insertManyResult.InsertedIDs))
}

func main() {
	router := gin.Default()
	router.GET("/", IndexHandler)
	router.GET("/recipes", recipesHandler.ListRecipesHandler)
	router.POST("/recipes", recipesHandler.NewRecipeHandler)
	router.PUT("/recipes/:id", recipesHandler.UpdateRecipeHandler)

	router.Run("localhost:8080")
}
