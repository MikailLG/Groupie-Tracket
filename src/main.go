package main

import (
    "Groupie-Tracker/src/helper"
    "Groupie-Tracker/src/routes"
    "fmt"
    "net/http"
)

func main() {
    helper.Load()
    serveRouter := routes.MainRouter()
    fmt.Println("Serveur lancé sur le port 8080")
    fmt.Println("Clique ici : http://localhost:8080")
    http.ListenAndServe(":8080", serveRouter)
}
