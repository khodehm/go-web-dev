package main

import (
	"html/template"
	"lerning-go-mongodb/controller"
	"lerning-go-mongodb/models"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// simple CRUD opprations with map and mongodb
var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("static/index.html"))
}
func main() {
	r := httprouter.New()
	uc := controller.NewUserController(getSession())
	r.GET("/", index)
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	log.Fatal(http.ListenAndServe(":9091", r))
}
func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.LoadUsers()
	w.Header().Add("Content-Type", "text/html")
	tmpl.ExecuteTemplate(w, "index.html", u)
}

// init mongo db - user map
func getSession() map[string]models.User {
	return models.LoadUsers()

	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	// c, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// return c.Database("lerning-go-mongodb")
}
