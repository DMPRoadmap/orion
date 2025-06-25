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

// RequestData structures the expected JSON payload
type RequestData struct {
    Cmd   string `json:"cmd"`
    Value string `json:"value"`
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

  var requestData RequestData
  if err := json.Unmarshal(body, &requestData); err != nil {
      http.Error(w, "Invalid JSON", http.StatusBadRequest)
      return
  }

  response, err := processCommand(requestData)  // Ensure you capture both the response and error
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

func processCommand(data RequestData) (interface{}, error) {
  switch data.Cmd {
  case "search_by_domain":
      fmt.Println("Searching by domain:", data.Value)
      return getOrgId(data.Value)
  case "search_by_ror_id":
      fmt.Println("Searching by ROR ID:", data.Value)
      return getOrgDetails(data.Value)
  default:
      return nil, fmt.Errorf("unknown command %s", data.Cmd)
  }
}

func getOrgDetails(rorID string) (interface{}, error) {
  dirPath := filepath.Join("/storage/orgs", rorID[:2], rorID[2:4])
	filePath := filepath.Join(dirPath, rorID + ".json")
  
  file, err := ioutil.ReadFile(filePath)
  if err != nil {
      return nil, fmt.Errorf("error reading file for ROR ID %s: %v", rorID, err)
  }

  var data interface{}
  if err := json.Unmarshal(file, &data); err != nil {
      return nil, fmt.Errorf("error parsing JSON file for ROR ID %s: %v", rorID, err)
  }

  return []interface{}{data}, nil
}

func getOrgId(domain string) (interface{}, error) {
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

    // Split file content by new lines
    lines := strings.Split(strings.TrimSpace(string(content)), "\n")

    return lines, nil
}


func main() {
  http.HandleFunc("/submit", handleSubmit)

  fmt.Println("Server is starting on port 8080...")
  err := http.ListenAndServe(":8080", nil)
  if err != nil {
      log.Fatal("ListenAndServe:", err)
  }
}
