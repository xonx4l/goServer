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

	router.Use(cors.Handler(cors.Optional{
		AllowedOriginal: []string {"http://*","http://*"},
		AllowedMethods: []string{"GET","POST","PUT","DELETE","OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposeHeaders: []string {"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router :=chi.NewRouter()

	v1Router.HandeleFunc("/ready", handlerReadiness) 

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
