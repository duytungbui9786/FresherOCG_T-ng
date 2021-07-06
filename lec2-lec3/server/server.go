package server

import (
	"fmt"
	"net/http"

	"github.com/techmaster.vn/helloworld/handlers"
)

func RunServer() {
	fmt.Println("Starting server. Please open http://localhost:8090/")
	// Defer function will be called when process exits
	defer func() {
		fmt.Println("Server is stopped")
	}()
	// --> browse to http://localhost:8090/hello will return Hello guys
	//  handler will read query param
	http.HandleFunc("/sum", handlers.SumWithParam)
	http.HandleFunc("/minus", handlers.MinuWithParam)
	http.HandleFunc("/multiply", handlers.MultiplyWithParam)
	http.HandleFunc("/divide", handlers.DivideWithParam)
	// run server
	if err := http.ListenAndServe(":8090", nil); err != nil {
		panic("Error when running server")
	}
}
