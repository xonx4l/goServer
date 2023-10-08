package main

import "fmt"
       "os"
	   "Log"

type apiConfig struct {
	 DB *database.Queries 
}	   

func main() {
	fmt.Println("hello world")

	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString = "" {
		log.fatal("PORT is not found in the environment")

	}

	dbURL := os.Getenv("DB_URL")
	if portString = "" {
		log.fatal("PORT is not found in the environment")

	}

	conn, err := sql.Open("postgres",dbURl)
	if err != nil{
		log.Fatal("can't connect to database: ", err)
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
	v1Router.HandeleFunc("/healthz", handlerReadiness) 
    v1Router.Get("/err", haandleErr)


    router.Mount("/v1", v1Router )
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
