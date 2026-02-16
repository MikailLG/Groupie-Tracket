package controllers

import (
	"Groupie-Tracker/src/helper"
	"Groupie-Tracker/src/services"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)
type PageData struct {
	Personnages    services.AllCharacters
	PageActuelle   int
	PagePrecedente int
	PageSuivante   int
	TotalPages     int
	FiltresURL     template.URL
}

func ListePersonnagesDisplay(w http.ResponseWriter, r *http.Request) {
	data, StatusCode, err := services.GetAllCharacters()
	if err != nil || StatusCode != http.StatusOK {
		http.Error(w, "Erreur d'obtention des personnages", StatusCode)
		return
	}
	recherche := r.URL.Query().Get("recherche")
	filtreStatut := r.URL.Query().Get("statut")
	filtreEquipage := r.URL.Query().Get("equipage")
	filtrePrime := r.URL.Query().Get("avec_prime")

	var resultats services.AllCharacters
	for _, perso := range *data {
		match := true
		if recherche != "" && !helper.Lettre(perso.Name, recherche) {
			match = false
		}
		if filtreStatut != "" && perso.Status != filtreStatut {
			match = false
		}
		if filtreEquipage != "" && !helper.Lettre(perso.Equipage.Name, filtreEquipage) {
			match = false
		}
		if filtrePrime == "true" && (perso.Bounty == "0" || perso.Bounty == "") {
			match = false
		}

		if match {
			resultats = append(resultats, perso)
		}
	}
	limite := 10
	totalResultats := len(resultats)
	pageStr := r.URL.Query().Get("page")
	pageActuelle, _ := strconv.Atoi(pageStr)
	if pageActuelle < 1 {
		pageActuelle = 1
	}
	totalPages := (totalResultats + limite - 1) / limite
	if pageActuelle > totalPages && totalPages > 0 {
		pageActuelle = totalPages
	}
	debut := (pageActuelle - 1) * limite
	fin := debut + limite
	if debut > totalResultats {
		debut = totalResultats
	}
	if fin > totalResultats {
		fin = totalResultats
	}
	var personnagesDeLaPage services.AllCharacters
	if totalResultats > 0 {
		personnagesDeLaPage = resultats[debut:fin]
	}
	filtresRaw := fmt.Sprintf("&recherche=%s&statut=%s&equipage=%s&avec_prime=%s",
		recherche, filtreStatut, filtreEquipage, filtrePrime)
	donnees := PageData{
		Personnages:    personnagesDeLaPage,
		PageActuelle:   pageActuelle,
		PagePrecedente: pageActuelle - 1,
		PageSuivante:   pageActuelle + 1,
		TotalPages:     totalPages,
		FiltresURL:     template.URL(filtresRaw),
	}
	helper.RenderTemplate(w, r, "liste_personnages", donnees)
}

func HomeDisplay(w http.ResponseWriter, r *http.Request) {
    helper.RenderTemplate(w, r, "home", nil)
}

func AboutDisplay(w http.ResponseWriter, r *http.Request) {
    helper.RenderTemplate(w, r, "about", nil)
}

func DetailsDisplay(w http.ResponseWriter, r *http.Request) {
    idString := r.URL.Query().Get("id")
    id, _ := strconv.Atoi(idString)

    data, _, _ := services.GetAllCharacters()
    var persoTrouve services.Character

    for _, p := range *data {
        if p.ID == id {
            persoTrouve = p
            break
        }
    }
    helper.RenderTemplate(w, r, "details", persoTrouve)
}

func CategoryDisplay(w http.ResponseWriter, r *http.Request) {
    data, _, _ := services.GetAllCharacters()
    helper.RenderTemplate(w, r, "categories", data)
}
