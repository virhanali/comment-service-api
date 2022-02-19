package main

import "fmt"


type App struct {}


func (app *App)  Run() error{
	fmt.Println("Setting the APP")
	return nil
}

func main() {
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up the app")
		fmt.Println(err)
	}
}