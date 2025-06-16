package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/url"
	"strings"
	"fmt"
	"io"
	"log"
	"os"
	"orion/rordb" // your package with Organization struct
)

func extractDomain(rawurl string) (string, error) {
    parsed, err := url.Parse(rawurl)
    if err != nil {
        return "", err
    }
    host := parsed.Hostname()
    if strings.HasPrefix(host, "www.") {
        host = strings.TrimPrefix(host, "www.")
    }
    return host, nil
}

func main() {
	filePath := "ror-data.json" // Your JSON file path

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	// Read opening array bracket
	_, err = decoder.Token()
	if err != nil {
		log.Fatalf("Error reading opening array token: %v", err)
	}

	count := 0
	calculated := ""
	// At the top of main, before the loop
	listDomains := make(map[string][]string)
	for decoder.More() {
		var raw json.RawMessage

		// Decode next JSON object as raw bytes
		err := decoder.Decode(&raw)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("Error decoding raw JSON: %v", err)
		}

		// Calculate SHA256 hash of the raw JSON bytes
		hash := sha256.Sum256(raw)
		hashString := hex.EncodeToString(hash[:])

		// Unmarshal raw JSON into your Organization struct
		var org rordb.Organization
		err = json.Unmarshal(raw, &org)
		if err != nil {
			log.Printf("Warning: failed to unmarshal organization #%d: %v", count, err)
		}
		
		if (org.Status == "active") {
			if len(org.Domains) == 0 {
				// Extract website link
				for _, link := range org.Links {
					if link.Type == "website" {
						if link.Value != "" {
							domain, err := extractDomain(link.Value)
							if err == nil && domain != "" {
									org.Domains = append(org.Domains, domain)
									listDomains[domain] = append(listDomains[domain], org.ID)
								}
							}
						}
					}
					calculated = "yes"
			} else {
				for _, domain := range org.Domains {
					listDomains[domain] = append(listDomains[domain], org.ID)
				}
				calculated = ""
			}
			id := strings.TrimPrefix(org.ID, "https://ror.org/")
			if (len(id) >=4) {
					_ = calculated
					_ = hashString
					// fmt.Printf("Org #%d, Hash: %s, ID: %s, Calculated %s, Domains %s\n", count, hashString, org.ID, calculated, org.Domains)
	//				rordb.SaveJSON("/storage", id, raw);
				} else {
					// here we should delete the 
			}
		}
		
		count++
	}

	// Read closing array bracket
	_, err = decoder.Token()
	if err != nil && err != io.EOF {
		log.Fatalf("Error reading closing array token: %v", err)
	}
	// fmt.Println("All org domains:")
	for domain, rorID := range listDomains {
			fmt.Printf("Domain - ROR: %s - %s\n", domain, rorID)
	}

	fmt.Printf("Processed %d organizations\n", count)

}
