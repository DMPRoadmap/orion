package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

// RequestData struct for parsing incoming JSON
type RequestData struct {
    Data any `json:"data"`
}

// ResponseData struct for creating outgoing JSON
type ResponseData struct {
    OrgName string `json:"org_name"`
}

func main() {
    http.HandleFunc("/data", handler)
    fmt.Println("Server starting on port :8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
        return
    }

    var requestData RequestData
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&requestData); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    responseData := ResponseData{OrgName: "abcd"}
    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(responseData); err != nil {
        log.Printf("Error encoding response: %v", err)
    }
}