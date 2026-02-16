package routes

import (
    "Groupie-Tracker/src/controllers"
    "net/http"
)

func routePersonnages(router *http.ServeMux) {
    router.HandleFunc("/", controllers.HomeDisplay)
    router.HandleFunc("/collection", controllers.ListePersonnagesDisplay)
    router.HandleFunc("/details", controllers.DetailsDisplay)
    router.HandleFunc("/about", controllers.AboutDisplay)
}
