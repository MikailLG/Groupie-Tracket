package routes

import (
    "net/http"
)

func MainRouter() *http.ServeMux {
    mainRouter := http.NewServeMux()
    routePersonnages(mainRouter)
    fileServerHandler := http.FileServer(http.Dir("assets"))
    mainRouter.Handle("/static/", http.StripPrefix("/static/", fileServerHandler))
    
    return mainRouter
}
