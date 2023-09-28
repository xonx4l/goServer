package main

import "fmt"
       "os"
	   "Log"

func main() {
	fmt.Println("hello world")

	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString = "" {
		log.fatal("PORT is not found in the environment")

	}

	router := chi.NewRouter()

	srv:= &http.server{
		Handler:router,
		Addr: ":" + portString,
	}

	 log .printf("Server starting on port %v", portString)
     err = srv.listenAndServer()
	 if err != nil{
		log.fatal(err)
	 }

	fmt.println("Port:", portString)
}
