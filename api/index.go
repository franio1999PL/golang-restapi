package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {

    searchParam := r.URL.Query().Get("search")
    limitParam := r.URL.Query().Get("limit")
    sortParam := r.URL.Query().Get("sort")
    pageParam := r.URL.Query().Get("page")

    targetURL := "https://cms.bladywebdev.pl/items/pocketposts"

    queryParams := make(map[string]string)

    if searchParam != "" {
        queryParams["search"] = searchParam
    }
    if limitParam != "" {
        queryParams["limit"] = limitParam
    }
    if sortParam != "" {
        queryParams["sort"] = sortParam
    }
    if pageParam != "" {
        queryParams["page"] = pageParam
    }

    queryStr := ""
    for key, value := range queryParams {
        if queryStr == "" {
            queryStr += "?"
        } else {
            queryStr += "&"
        }
        queryStr += key + "=" + value
    }

    targetURL += queryStr

    response, err := http.Get(targetURL)
    if err != nil {
        http.Error(w, "Błąd podczas wykonywania zapytania GET", http.StatusInternalServerError)
        return
    }
    defer response.Body.Close()

    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        http.Error(w, "Błąd podczas odczytywania treści odpowiedzi", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")

    // Wysyłamy odpowiedź na zapytanie klienta
    w.Write(body)
    // end

        
}

func main() {
    // Tworzymy nowy router HTTP
    http.HandleFunc("/", Handler)

    // Uruchamiamy serwer na porcie 8080
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Błąd serwera:", err)
    }
}