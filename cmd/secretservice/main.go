package main

import (
	"net/http"

	secretservice "github.com/AndrewCopeland/secret-service"
)

func main() {
	r := secretservice.NewRouter()
	http.ListenAndServe(":3000", r)
}
