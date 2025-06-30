// SPDX-License-Identifier: AGPL-3.0-or-later
// Copyright © 2025 Digital Curation Centre (UK) and contributors

package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "path/filepath"
    "strings"
    "crypto/sha256"
    "encoding/hex"
)


type ApiRequestData struct {
    Cmd   string          `json:"cmd"`
    Value json.RawMessage `json:"value"`  // holds raw JSON — can be string or array
}

func handleSubmit(w http.ResponseWriter, r *http.Request) {
  if r.Method != "POST" {
      http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
      return
  }

  body, err := ioutil.ReadAll(r.Body)
  if err != nil {
      http.Error(w, "Error reading request body", http.StatusInternalServerError)
      return
  }
  defer r.Body.Close()

  var apiRequestData ApiRequestData
  if err := json.Unmarshal(body, &apiRequestData); err != nil {
      http.Error(w, "Invalid JSON", http.StatusBadRequest)
      return
  }

  response, err := processCommand(apiRequestData)
  if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
  }

  w.Header().Set("Content-Type", "application/json")
  jsonResponse, err := json.Marshal(response)
  if err != nil {
      http.Error(w, "Error creating response JSON", http.StatusInternalServerError)
      return
  }

  w.Write(jsonResponse)
}

func processCommand(data ApiRequestData) (interface{}, error) {
    switch data.Cmd {
    case "search_by_domain":
        var domain string
        if err := json.Unmarshal(data.Value, &domain); err != nil {
            return nil, fmt.Errorf("expected a string for domain: %v", err)
        }
        fmt.Println("Searching by domain:", domain)
        return getOrgsByDomain(domain)

    case "search_by_ror_id":
        // Try to parse as array
        var ids []string
        if err := json.Unmarshal(data.Value, &ids); err == nil {
            fmt.Println("Searching by multiple ROR IDs:", ids)
            return getOrgsById(ids)
        }

        // If not an array, try to parse as a single ID
        var id string
        if err := json.Unmarshal(data.Value, &id); err != nil {
            return nil, fmt.Errorf("invalid value for search_by_ror_id: must be string or array")
        }

        fmt.Println("Searching by single ROR ID:", id)
        return getOrgsById([]string{id})

    default:
        return nil, fmt.Errorf("unknown command %s", data.Cmd)
    }
}

func getOrgsById(rorIDs []string) (interface{}, error) {
    var result []interface{}

    for _, rorID := range rorIDs {
        // Build file path
        dirPath := filepath.Join("/storage/orgs", rorID[:2], rorID[2:4])
        filePath := filepath.Join(dirPath, rorID + ".json")

        // Read the file
        file, err := ioutil.ReadFile(filePath)
        if err != nil {
            return nil, fmt.Errorf("error reading file for ROR ID %s: %v", rorID, err)
        }

        var data interface{}
        if err := json.Unmarshal(file, &data); err != nil {
            return nil, fmt.Errorf("error parsing JSON file for ROR ID %s: %v", rorID, err)
        }

        result = append(result, data)
    }

    return result, nil
}


func getOrgsByDomain(domain string) (interface{}, error) {
  // Hash the domain
  hash := sha256.Sum256([]byte(domain))
  hashedDomain := hex.EncodeToString(hash[:])

  // Build file path
  dirPath := filepath.Join("/storage/domains", hashedDomain[:2], hashedDomain[2:4])
  filePath := filepath.Join(dirPath, hashedDomain + ".txt")

    // Read the file
    content, err := ioutil.ReadFile(filePath)
    if err != nil {
        return nil, fmt.Errorf("error reading file for domain %s: %v", domain, err)
    }

    rorIDs := strings.Split(strings.TrimSpace(string(content)), "\n")

    // Get org details for those ROR IDs
    orgDetails, err := getOrgsById(rorIDs)
    if err != nil {
        return nil, fmt.Errorf("error retrieving org details: %v", err)
    }

    // Construct the final response
    result := map[string]interface{}{
        "ids":  rorIDs,
        "orgs": orgDetails,
    }

    return result, nil
}


func main() {
  http.HandleFunc("/submit", handleSubmit)

  fmt.Println("Server is starting on port 8080...")
  err := http.ListenAndServe(":8080", nil)
  if err != nil {
      log.Fatal("ListenAndServe:", err)
  }
}
