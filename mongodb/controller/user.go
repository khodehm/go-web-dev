package controller

import (
	"encoding/json"
	"fmt"
	"lerning-go-mongodb/models"
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

// UserController struct stores a map or mongodb  database
type UserController struct {
	// db *mongo.Database
	db map[string]models.User
}

// NewUserController inits app user-map or mongodb database
func NewUserController(db map[string]models.User) *UserController {
	return &UserController{db: db}
}

// GetUser gets data form map or mongodb
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// check user existens in map
	u, ok := uc.db[p.ByName("id")]
	if !ok {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	// Mongo db find user logic and validation
	// if !bson.IsObjectIdHex(p.ByName("id")) {
	// http.Error(w, "user not found", http.StatusNotFound)
	// return
	// }
	// id := bson.ObjectIdHex(p.ByName("id"))
	// uc.db.Collection("users").FindOne(r.Context(), bson.M{"id": id}).Decode(&u)
	ju, _ := json.Marshal(&u)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s \n", ju)
}

// CreateUser creates a user in mongodb or a map
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var u models.User
	decoder := json.NewDecoder(r.Body)
	//check request body and validation
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&u); err != nil {
		http.Error(w, "missing required fields", http.StatusBadRequest)
		return
	}
	_, ok := uc.db[u.ID]
	if ok {
		http.Error(w, "user already exists", http.StatusBadRequest)
		return
	}
	uid, _ := uuid.NewV6()
	u.ID = uid.String()
	uc.db[u.ID] = u
	models.StoreUser(uc.db)
	// Mongo db create user logic and validatiuon
	// u.ID = bson.NewObjectId().String()
	// _, err := uc.db.Collection("users").InsertOne(r.Context(), u)
	// if err != nil {
	// http.Error(w, err.Error(), http.StatusBadRequest)
	// return
	// }
	fmt.Fprintln(w, map[string]any{"user created": u})

}

// DeleteUser Delets a user in map or mongodb
func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// mongodb DELETE user logic and validation
	// if !bson.IsObjectIdHex(p.ByName("id")) {
	// http.Error(w, "user not found", http.StatusNotFound)
	// return
	// }
	// id := bson.ObjectIdHex(p.ByName("id"))
	// _, err := uc.db.Collection("users").DeleteOne(r.Context(), id)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// check user exists in map
	u, ok := uc.db[p.ByName("id")]
	if !ok {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	// delete user from map
	delete(uc.db, u.ID)
	models.StoreUser(uc.db)
	fmt.Fprintf(w, "user %v deleted! \n", u.ID)
}
