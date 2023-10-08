package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/PratikforCoding/Syntaxia/controller" 
	"github.com/PratikforCoding/Syntaxia/database"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongouri := os.Getenv("connectlink")

	AttendeesCol, err := database.CreateDB(mongouri)
	if err != nil {
		log.Fatal("Didn't create connection to mongodb")
	}
	defer database.CloseDB()

	apiCfg := controller.NewAPIConfig(AttendeesCol)

	fmt.Println("MongoDB API")
	router := chi.NewRouter()
	apiRouter := chi.NewRouter()

	fsHandler := http.StripPrefix("/app", http.FileServer(http.Dir(".")))
	router.Handle("/app/", fsHandler)
	router.Handle("/app/*", fsHandler)
	
	apiRouter.Get("/getByYear", apiCfg.HandlerGetAttendeesbyYear)
	apiRouter.Get("/getAllAttendees", apiCfg.HandlerGetAllAttendees)
	apiRouter.Get("/getAttendee", apiCfg.HandlerGetAttendeebyId)
	apiRouter.Post("/register", apiCfg.HandlerRegister)
	apiRouter.Put("/done", apiCfg.HandlerTaken)

	
	router.Mount("/api", apiRouter)
	

	corsMux := middlewareCors(router)
	server := &http.Server {
		Addr: ":8080",
		Handler: corsMux,
	}

	log.Println("Server is getting started at port: 8080 ....")
	log.Fatal(server.ListenAndServe())
	log.Println("Server is runnig at port: 8080 ....")
}