// Copyright 2025 Digital Curation Centre (UK) and contributors, Licence AGPLv3

package rordb

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
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
