// Copyright 2025 Digital Curation Centre (UK) and contributors, Licence AGPLv3

package rordb

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
    "strings"
)

func SaveJSON(basePath, orgID string, raw json.RawMessage) error {
	if len(orgID) < 4 {
		return fmt.Errorf("organization ID too short: %s", orgID)
	}

	dirPath := filepath.Join(basePath, orgID[:2], orgID[2:4])
	filePath := filepath.Join(dirPath, orgID + ".json")

	if err := os.MkdirAll(dirPath, 0755); err != nil {
		return fmt.Errorf("failed to create directories: %w", err)
	}

	if err := os.WriteFile(filePath, raw, 0644); err != nil {
		return fmt.Errorf("failed to write JSON file: %w", err)
	}

	return nil
}

func SaveHashedDomains(basePath string, domains map[string][]string) error {
    // Iterate through the domain map
    for domain, rorIDs := range domains {
        
        dirPath := filepath.Join(basePath, domain[:2], domain[2:4])
	    filePath := filepath.Join(dirPath, domain + ".txt")

        // create directories 
        if err := os.MkdirAll(dirPath, 0755); err != nil {
            return fmt.Errorf("failed to create directory for domain %s: %w", domain, err)
        }
        
        // convert the ROR IDs slice into a single string with each ID on a new line
        data := strings.Join(rorIDs, "\n") + "\n"

        // use os.WriteFile to write the data
        if err := os.WriteFile(filePath, []byte(data), 0644); err != nil {
            return fmt.Errorf("failed to write ROR IDs to file %s: %w", filePath, err)
        }
    }

	return nil
}
