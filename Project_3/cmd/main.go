package main

import (
	"fmt"
	"net/http"

	"github.com/samarth8765/Project_3/pkg/routes"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStores(r)
	http.Handle("/", r)

	fmt.Print("Listening at PORT 8080\n")

	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Printf("Error occured %s", err)
	}
}
