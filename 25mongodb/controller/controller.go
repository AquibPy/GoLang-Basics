package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/AquibPy/mongoapi/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://aquibpy:<password>@cluster0.sdx7i86.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB Connection Successfuly !!!!!")

	collection = client.Database(dbName).Collection(colName)

	fmt.Println("Collection instance is ready !!!")
}

// MongoDb helpers -file

// insert one record

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 Movie in db with id:", inserted.InsertedID)
}

func updateOneMovie(movieID string) {
	id, _ := primitive.ObjectIDFromHex(movieID)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count :", result.ModifiedCount)
}

func deleteOneMovie(movieID string) {
	id, _ := primitive.ObjectIDFromHex(movieID)
	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie got deleted with delete count :", result.DeletedCount)
}

func deleteAllMovies() int64 {
	result, err := collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("All Records Deleted Successfully",result.DeletedCount)
	return result.DeletedCount

}

// get all movies from db

func getAllMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	var movies []primitive.M
	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	return movies
}


// @Summary Get all movies
// @Description Get all movies from the watchlist
// @Tags movies
// @Produce  json
// @Success 200 {array} model.Netflix
// @Router /api/movies [get]
func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

// @Summary Create a movie
// @Description Add a new movie to the watchlist
// @Tags movies
// @Accept  json
// @Produce  json
// @Param movie body model.Netflix true "Movie to create"
// @Success 201 {object} model.Netflix
// @Router /api/movie [post]
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)

}

// @Summary Mark a movie as watched
// @Description Update a movie's watched status to true
// @Tags movies
// @Accept  json
// @Produce  json
// @Param id path string true "Movie ID"
// @Success 200 {string} string "Movie ID"
// @Router /api/movie/{id} [put]
func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneMovie(params["id"])

	json.NewEncoder(w).Encode(params["id"])
}

// @Summary Delete a movie
// @Description Delete a movie from the watchlist by ID
// @Tags movies
// @Accept  json
// @Produce  json
// @Param id path string true "Movie ID"
// @Success 200 {string} string "Movie ID"
// @Router /api/movie/{id} [delete]
func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

// @Summary Delete all movies
// @Description Delete all movies from the watchlist
// @Tags movies
// @Accept  json
// @Produce  json
// @Success 200 {integer} int "Number of deleted movies"
// @Router /api/deleteallmovie [delete]
func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovies()
	json.NewEncoder(w).Encode(count)
}
