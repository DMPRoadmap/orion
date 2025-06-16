// Copyright 2025 Digital Curation Centre (UK) and contributors, Licence AGPLv3

package rordb

type Organization struct {
	ID            string         `json:"id"`
	Locations     []Location     `json:"locations"`
	Established   *int           `json:"established"`
	ExternalIDs   []ExternalID   `json:"external_ids"`
	Domains       []string       `json:"domains"`
	Links         []Link         `json:"links"`
	Names         []Name         `json:"names"`
	Relationships []Relationship `json:"relationships"`
	Status        string         `json:"status"`
	Types         []string       `json:"types"`
	Admin         Admin          `json:"admin"`
}

type Location struct {
	GeonamesID      int            `json:"geonames_id"`
	GeonamesDetails GeoNameDetails `json:"geonames_details"`
}

type GeoNameDetails struct {
	ContinentCode           string  `json:"continent_code"`
	ContinentName           string  `json:"continent_name"`
	CountryCode             string  `json:"country_code"`
	CountryName             string  `json:"country_name"`
	CountrySubdivisionCode  string  `json:"country_subdivision_code"`
	CountrySubdivisionName  string  `json:"country_subdivision_name"`
	Lat                     float64 `json:"lat"`
	Lng                     float64 `json:"lng"`
	Name                    string  `json:"name"`
}

type ExternalID struct {
	Type      string   `json:"type"`
	All       []string `json:"all"`
	Preferred *string  `json:"preferred"`
}

type Link struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Name struct {
	Value string   `json:"value"`
	Types []string `json:"types"`
	Lang  *string  `json:"lang"`
}

type Relationship struct {
	Type  string `json:"type"`
	Label string `json:"label"`
	ID    string `json:"id"`
}

type Admin struct {
	Created      AdminDate `json:"created"`
	LastModified AdminDate `json:"last_modified"`
}

type AdminDate struct {
	Date          string `json:"date"`
	SchemaVersion string `json:"schema_version"`
}
