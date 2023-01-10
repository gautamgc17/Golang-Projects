package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gautamgc75/mongoapi/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://admin:hZ2U8pxUCEeARF67@cluster0.rzm4ekz.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const collectinName = "watchList"

// reference of mongoDB collection
var collection *mongo.Collection

// connect with the mongoDB
// initialization function that runs only once at start of this application
func init() {

	// create a new client option and ApplyURI parses the given URI and sets options accordingly for the client
	clientOption := options.Client().ApplyURI(connectionString)
	fmt.Printf("Type of Option is: %T\n", clientOption)

	// connect to mongodb - Connect creates a new Client and then initializes it using the Connect method
	// TODO returns a non-nil, empty Context. It's unclear which Context to use or it is not yet available
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Type of client is: %T\n", client)
	fmt.Println("MongoDB client instance created and connection success")

	// handler for db and collection inside it
	collection = client.Database(dbName).Collection(collectinName)
	fmt.Println(collection)
	fmt.Printf("Type of Collection is: %T\n", collection)
	fmt.Println("Collection instance is ready")
}

// helper functions - not exporting them

// insert 1 record into DB
func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Type of Insertion is: %T\n", inserted)
	fmt.Println("Inserted one movie in DB with ID: ", inserted.InsertedID)
}

// update 1 record
func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}
	// pass the context and query filter and data with which update is required , matchCount of 1 returned if operation successful
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Matched Count:", result.MatchedCount)
	fmt.Printf("Type of ID is: %T", id)
	fmt.Println(id)
	fmt.Printf("Type of filer is: %T", filter)
	fmt.Println(filter)
	fmt.Printf("Type of update is: %T", update)
	fmt.Println(update)
}

// delete 1 record
func deleteOneRecord(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie got delete with delete count: ", deleteCount)
}

// delete all records from MongoDB
func deleteAllMovie() int64 {
	// D represents order and M is non-order representation , bson.D{{}} with emptys set means select everything
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Number of Movies Deleted:", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

// get all movies from database
func getAllMovies() []primitive.M {
	// executes a find command and returns a Cursor over the matching documents in the collection
	//  not get actual values but return of a cursor which is a type of MongoDB object
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	var movies []primitive.M

	// gets the next document for this cursor
	for cursor.Next(context.Background()) {
		var movie bson.M
		// Decode will unmarshal the current document into val
		if err := cursor.Decode(&movie); err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	return movies
}

// Actual controllers - exportable
// Get all movies
func GetAllMovies(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

// Insert movie
func CreateMovie(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type" , "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods" , "POST")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

// Mark movie as watched
func MarkAsWatched(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type" , "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods" , "PUT")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

// Delete a movie
func DeleteAMovie(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type" , "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods" , "DELETE")

	params := mux.Vars(r)
	deleteOneRecord(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

// Delete all movies
func DeleteAllMovies(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type" , "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods" , "DELETE")

	count := deleteAllMovie()
	json.NewEncoder(w).Encode(count)
}



