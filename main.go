package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"reflect"
)

func main() {

	// Create Database if not exist and execute pre-written initial query from "config.go"
	createDB(initialQuery)

	// Get port from dedicated server environment, if running locally assign port to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Starting handler to run our front-end application
	http.Handle("/", http.FileServer(http.Dir("./front")))

	// Our API endpoints
	endpoint("/api/register", register)
	endpoint("/api/login", login)
	endpoint("/api/logout", logout, "check JWT")
	endpoint("/api/addpost", addpost, "check JWT")
	endpoint("/api/updpost", updpost, "check JWT")
	endpoint("/api/viewposts", viewposts)
	endpoint("/api/readpost", readpost)

	//TODO
	// Pagination for viewposts
	// Show comments by postID
	// Write comment (secure)
	// Post reaction (secure)

	// Listen server
	log.Println("Running http://localhost:" + port)
	e := http.ListenAndServe(":"+port, nil)
	log.Println(e)

}

type ctxKey string

// Middleware handler
func endpoint(path string, page func(w http.ResponseWriter, r *http.Request), secure ...interface{}) {

	fn := func(w http.ResponseWriter, r *http.Request) {

		// Set headers for CORS enabling
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5000")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		isValid, id := validateJWT(w, r)

		// JWT validation handler
		if len(secure) > 0 && !isValid {
			log.Println("User Error: JWT is not valid")
			http.Error(w, http.StatusText(403), 403)
			return
		}

		// Save userID to context
		var key ctxKey = "userID"
		ctx := context.WithValue(r.Context(), key, id)
		req := r.WithContext(ctx)

		// Error handler
		defer func() {
			if err := recover(); err != nil {
				switch err.(type) {

				case *reflect.ValueError:
					http.Error(w, http.StatusText(404), 404)
				case *json.SyntaxError:
					http.Error(w, http.StatusText(400), 400)
				default:
					log.Printf("Server Error: %+v", err)
					http.Error(w, http.StatusText(500), 500)
				}
			}
		}()

		http.HandlerFunc(page).ServeHTTP(w, req)
	}
	http.Handle(path, http.HandlerFunc(fn))
}
