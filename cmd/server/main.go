package main

import (
	"fmt"
	"net/http"

	transportHttp "github.com/virhanali/comment-service-api/internal/transport/http"
)

// App - the struct which contains things like pointers
// to database connection
type App struct{}


// run -  sets up our connection
func (app *App) Run() error {
	fmt.Println("Setting the APP")
	handler := transportHttp.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}
	return nil
}

func main() {
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up the app")
		fmt.Println(err)
	}
}
