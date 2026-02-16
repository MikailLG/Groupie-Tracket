package helper

import (
    "bytes"
    "fmt"
    "html/template"
    "log"
    "net/http"
    "strings"
)

var listeTemplate *template.Template

func Load() {
    temp, tempErr := template.ParseGlob("templates/*.html")
    if tempErr != nil {
        log.Fatalf("Erreur template - %s", tempErr.Error())
        return
    }
    listeTemplate = temp
    fmt.Println("Template - chargement des templates terminé")
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, name string, data interface{}) {
    var buffer bytes.Buffer
    errRender := listeTemplate.ExecuteTemplate(&buffer, name, data)
    if errRender != nil {
        http.Error(w, "Erreur lors du chargement du template", http.StatusInternalServerError)
        return
    }
    buffer.WriteTo(w)
}

func Lettre(texte, recherche string) bool {
    return strings.Contains(strings.ToLower(texte), strings.ToLower(recherche))
}
