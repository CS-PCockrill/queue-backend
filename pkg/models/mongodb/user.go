package mongodb

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/CS-PCockrill/queue/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

//UserFunctions is used to inject mongo driver client and context into the application
type UserFunctions struct {
	CLIENT *mongo.Client
}

//Insert method is used to insert new user into the User collection
func (u *UserFunctions) Insert(username, firstname, lastname, email, password string) error {

	//Insert user to the database
	newUser := u.CLIENT.Database("queue")
	userCollection := newUser.Collection("user")
	var user models.User

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}
	user.UserName = username
	user.FirstName = firstname
	user.LastName = lastname
	user.Email = email
	user.Password = hashedPassword
	user.Created = time.Now().UTC()
	user.Active = true

	//Insert the user into the database
	result, err := userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		fmt.Println(err)
	}

	//Check ID of the inserted document
	insertedID := result.InsertedID.(primitive.ObjectID)
	fmt.Println(insertedID)

	return nil
}

func (u *UserFunctions) Update(street, city, state, zip string) error {
	user := u.CLIENT.Database("queue")
	userCollection := user.Collection("user")
	
	var address models.Address
	address.Street = street
	address.City = city
	address.State = state
	address.Zip = zip

	aByte, err := bson.Marshal(&address)
	if err != nil {
		return err
	}

	var update bson.M
	err = bson.Unmarshal(aByte, &update)
	if err != nil {
		return err
	}
	// var user models.User

	// TODO: change idStr to actual user _id string to filter update
	idStr := "5f904fffa9ef191c89c8af65" // id.Hex()
	user_id, _ := primitive.ObjectIDFromHex(idStr)
	filter := bson.M{"_id": bson.M{"$eq": user_id}}

	_, err = userCollection.UpdateOne(
		context.TODO(),
		filter, 
		bson.D{{Key: "$set", Value: bson.M{"address": update},
	}})

	if err != nil {
        fmt.Println("UpdateOne() result ERROR:", err)
        return err
    }
	return nil
}

//Authenticate method to confirm if a user exists in the database
func (u *UserFunctions) Authenticate(email, password string) (primitive.ObjectID, error) {
	//Authenticate user before login by retrieving the user id and hashed password from database
	//Hash the password entered and compare it to the one retrieved from database

	newUser := u.CLIENT.Database("queue")
	userCollection := newUser.Collection("user")

	var output models.User
	filter := bson.M{"email": email}

	err := userCollection.FindOne(context.TODO(), filter).Decode(&output)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			return output.ID, errors.New("filter did not match any documents in the collection")
		}
		log.Fatal(err)
	}
	// Check whether the hashed password and plain-text password provided match
	err = bcrypt.CompareHashAndPassword(output.Password, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return output.ID, models.ErrInvalidCredentials
		} else {
			return output.ID, err
		}
	}
	fmt.Println("Authenticated")
	return output.ID, nil
}

//GetUsers to get all users
func (u *UserFunctions) GetUsers() []models.User {
	myUser := u.CLIENT.Database("queue")
	userCollection := myUser.Collection("user")
	var newUser []models.User
	cursor, err := userCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		panic(err)
	}
	defer cursor.Close(context.TODO())
	if err = cursor.All(context.TODO(), &newUser); err != nil {
		fmt.Println(err)
		panic(err)
	}
	return newUser
}

//GetUser takes a username and return that user
func (u *UserFunctions) GetUser(username string) models.User {
	myUser := u.CLIENT.Database("queue")
	userCollection := myUser.Collection("user")
	var newUser models.User
	err := userCollection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&newUser)
	if err != nil {
		panic(err)
	}
	return newUser
}

//GetUser to get all users
// func (u *UserFunctions) GetUser() []models.Students {
// 	//Get a user with its id
// 	myUser := u.CLIENT.Database("students")
// 	moviesCollection := myUser.Collection("grades")
// 	var students []models.Students
// 	options := options.Find()
// 	options.SetLimit(10)
// 	cursor, err := moviesCollection.Find(u.CTX, bson.M{}, options)

// 	if err != nil {
// 		panic(err)
// 	}
// 	defer cursor.Close(u.CTX)
// 	//for cursor.Next(u.CTX){
// 	//	var oneMovie models.Movie
// 	//	cursor.Decode(oneMovie)
// 	//	movies = append(movies, oneMovie)
// 	//}
// 	if err = cursor.All(u.CTX, &students); err != nil {
// 		fmt.Println(err)
// 		panic(err)
// 	}
// 	return students
// }
