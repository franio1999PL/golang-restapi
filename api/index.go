package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func pocketpostsHandler(w http.ResponseWriter, r *http.Request) {
    // Odczytujemy parametry z zapytaniax
    searchParam := r.URL.Query().Get("search")
    limitParam := r.URL.Query().Get("limit")
    sortParam := r.URL.Query().Get("sort")

    // Adres URL docelowy
    targetURL := "https://cms.bladywebdev.pl/items/pocketposts"

    // Tworzymy mapę parametrów, które chcemy przekazać
    queryParams := make(map[string]string)

    // Dodajemy parametry do mapy, jeśli są dostępne w zapytaniu
    if searchParam != "" {
        queryParams["search"] = searchParam
    }
    if limitParam != "" {
        queryParams["limit"] = limitParam
    }
    if sortParam != "" {
        queryParams["sort"] = sortParam
    }

    // Tworzymy ciąg z parametrami do dodania do adresu URL
    queryStr := ""
    for key, value := range queryParams {
        if queryStr == "" {
            queryStr += "?"
        } else {
            queryStr += "&"
        }
        queryStr += key + "=" + value
    }

    // Dodajemy ciąg z parametrami do adresu URL docelowego
    targetURL += queryStr

    // Wykonujemy zapytanie GET do adresu docelowego
    response, err := http.Get(targetURL)
    if err != nil {
        http.Error(w, "Błąd podczas wykonywania zapytania GET", http.StatusInternalServerError)
        return
    }
    defer response.Body.Close()

    // Odczytujemy odpowiedź
    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        http.Error(w, "Błąd podczas odczytywania treści odpowiedzi", http.StatusInternalServerError)
        return
    }

    // Ustalamy nagłówek Content-Type jako "application/json", zakładając, że odpowiedź to JSON
    w.Header().Set("Content-Type", "application/json")

    // Wysyłamy odpowiedź na zapytanie klienta
    w.Write(body)
}

func main() {
    // Tworzymy nowy router HTTP
    http.HandleFunc("/", pocketpostsHandler)

    // Uruchamiamy serwer na porcie 8080
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Błąd serwera:", err)
    }
}
