package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"crypto/tls"
)

type Fruit struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
	RomanName   string `json:"roman_name"`
}

type Crew struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	IsYonko     bool   `json:"is_yonko"`
}

type Character struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Job      string `json:"job"`
	Size     string `json:"size"`
	Birthday string `json:"birthday"`
	Age      string `json:"age"`
	Bounty   string `json:"bounty"`
	Status   string `json:"status"`
	Equipage Crew   `json:"crew"`
	Fruit    Fruit  `json:"fruit"`
}

type AllCharacters []Character

func GetAllCharacters() (*AllCharacters, int, error) {
    tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
    _client := http.Client{
        Timeout:   10 * time.Second,
        Transport: tr,
    }
    apiUrl := "https://api.api-onepiece.com/v2/characters/fr"
    req, reqErr := http.NewRequest(http.MethodGet, apiUrl, nil)
    if reqErr != nil {
        return nil, http.StatusInternalServerError, fmt.Errorf("GetAllCharacters - Erreur préparation : %s", reqErr)
    }
    
    res, resErr := _client.Do(req)
    if resErr != nil {
        return nil, http.StatusInternalServerError, fmt.Errorf("GetAllCharacters - Erreur envoi : %s", resErr)
    }
    defer res.Body.Close()

    if res.StatusCode != http.StatusOK {
        return nil, res.StatusCode, fmt.Errorf("GetAllCharacters - Code erreur : %d", res.StatusCode)
    }

    var data AllCharacters
    decodeErr := json.NewDecoder(res.Body).Decode(&data)
    if decodeErr != nil {
        return nil, http.StatusInternalServerError, fmt.Errorf("GetAllCharacters - Erreur décodage : %s", decodeErr.Error())
    }
    return &data, res.StatusCode, nil
}
