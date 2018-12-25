package main

import (
	_ "database/sql"
	"encoding/json"
	"log"

	"github.com/jackwhelpton/fasthttp-routing/cors"

	routing "github.com/jackwhelpton/fasthttp-routing"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/valyala/fasthttp"
)

var (
	connStr = "user=root password=root dbname=recipes sslmode=disable port=26257"
)

// Recipes estructura de la tabla con el mismo nombre
type Recipes struct {
	ID           *string `json:"id"`
	Name         string  `json:"name"`
	Calories     float64 `json:"calories"`
	Img          string  `json:"img"`
	Ingredients  string  `json:"ingredients"`
	Instructions string  `json:"instructions"`
}

func main() {
	router := routing.New()
	router.Use(cors.Handler(cors.AllowAll))
	recipes := router.Group("/recipes")
	recipes.Get("", getAll).Post(post)
	recipes.Get("/", getAll).Post(post)
	recipes.Get("/<id>", query).Put(put).Delete(delete)

	fasthttp.ListenAndServe(":8082", router.HandleRequest)
}

func getAll(c *routing.Context) error {
	name := c.FormValue("name")
	c.SetContentType("application/json")
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalf("error initializing database: %v\n", err)
	}
	recipes := []Recipes{}
	if name == nil {
		db.Select(&recipes, "SELECT * FROM recipes")
	} else {
		db.Select(&recipes, "SELECT * FROM recipes WHERE LOWER(name) LIKE ('%' || LOWER($1) || '%')", string(name))
	}
	db.Close()
	body, err := json.Marshal(recipes)
	if err != nil {
		responseError(c, err)
		return nil
	}
	c.SetStatusCode(fasthttp.StatusOK)
	c.SetBody(body)
	return nil
}

func query(c *routing.Context) error {
	id := c.Param("id")
	c.SetContentType("application/json")
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalf("error initializing database: %v\n", err)
	}
	recipe := Recipes{}
	err = db.Get(&recipe, "SELECT * FROM recipes where id=$1", id)
	var body []byte
	if err != nil {
		body, _ = json.Marshal(map[string]string{})
	} else {
		body, _ = json.Marshal(recipe)
	}
	db.Close()
	c.SetStatusCode(fasthttp.StatusOK)
	c.SetBody(body)
	return nil
}

func post(c *routing.Context) error {
	request := Recipes{}
	c.SetContentType("application/json")
	err := json.Unmarshal(c.PostBody(), &request)
	if err != nil {
		responseError(c, err)
		return nil
	}
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalf("error initializing database: %v\n", err)
	}
	_, err = db.NamedExec("INSERT INTO recipes (name, calories, img, ingredients, instructions) VALUES (:name, :calories, :img, :ingredients, :instructions)", request)
	if err != nil {
		db.Close()
		responseError(c, err)
		return nil
	}
	db.Close()
	body, _ := json.Marshal(map[string]interface{}{"success": true})
	c.SetContentType("application/json")
	c.SetStatusCode(fasthttp.StatusCreated)
	c.SetBody(body)
	return nil
}

func put(c *routing.Context) error {
	id := c.Param("id")
	c.SetContentType("application/json")
	request := Recipes{}
	err := json.Unmarshal(c.PostBody(), &request)
	if err != nil {
		responseError(c, err)
		return nil
	}
	recipe := Recipes{}
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalf("error initializing database: %v\n", err)
	}
	err = db.Get(&recipe, "SELECT * FROM recipes where id=$1", id)
	if err != nil {
		db.Close()
		responseError(c, err)
		return nil
	}
	recipe.Name = request.Name
	recipe.Calories = request.Calories
	recipe.Img = request.Img
	recipe.Ingredients = request.Ingredients
	recipe.Instructions = request.Instructions

	_, err = db.NamedExec("UPDATE recipes SET name=:name, calories=:calories, img=:img, ingredients=:ingredients, instructions=:instructions WHERE id=:id", recipe)
	if err != nil {
		db.Close()
		responseError(c, err)
		return nil
	}
	db.Close()
	body, _ := json.Marshal(map[string]interface{}{"success": true})
	c.SetStatusCode(fasthttp.StatusOK)
	c.SetBody(body)
	return nil
}

func delete(c *routing.Context) error {
	id := c.Param("id")
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalf("error initializing database: %v\n", err)
	}
	_, err = db.Exec("DELETE FROM recipes WHERE id=$1", id)
	if err != nil {
		db.Close()
		responseError(c, err)
		return nil
	}
	db.Close()
	body, _ := json.Marshal(map[string]interface{}{"success": true})
	c.SetStatusCode(fasthttp.StatusOK)
	c.SetBody(body)
	return nil
}

func responseError(c *routing.Context, err error) {
	body, _ := json.Marshal(map[string]interface{}{"success": false, "error": err.Error()})
	c.SetStatusCode(fasthttp.StatusInternalServerError)
	c.SetBody(body)
}
