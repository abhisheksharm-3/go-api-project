package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/abhisheksharm-3/go-api-project/router"
)

func main() {
	fmt.Println("MONGO GO API")
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000 ...")
}
