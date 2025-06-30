# Orion

Orion (Roadmap ROR Integration) is a lightweight service designed to integrate the [Research Organization Registry](https://ror.org/) (ROR) into systems like [Roadmap](https://github.com/DMPRoadmap/roadmap).
Built in Go, Orion exposes a simple HTTP API for linking email domains to research organizations.

Built with efficiency and minimal resource consumption in mind, Orion avoids using a traditional database, aligning with principles of sustainable and lightweight infrastructure.
Instead, it leverages a structured local filesystem-based storage to serve organization records and domain-to-ROR mappings.

This design enables fast, efficient lookups with minimal dependencies and overhead.
This is ideal for lightweight deployments or containerized environments.

## Getting Started

### Requirements
- [Go](https://golang.org/doc/install)
- [Docker](https://docs.docker.com/get-docker/) (optional)
- Input data to generate ROR records and domain mappings

### Generate Storage Data

Run the following to build the `/storage` structure used by the API:

```bash
go run json.go
```

This will create:

    /storage/orgs/ → individual JSON files for each ROR ID

    /storage/domains/ → hashed domain .txt files mapping to ROR IDs

### Start the API Server
Locally:

`go run api.go`

Or in Docker:

```bash
docker build -t orion .

docker run -it --rm \
  -v $HOME/tmp/storage:/storage \
  -v $(pwd):/app \
  -w /app \
  -p 8080:8080 \
  orion bash

# then inside the container:
go run api.go
```

### Testing the API

Once the server is running (either locally or inside Docker), you can test the endpoints using curl or tools like Postman.

#### Test 1: Search by Domain

This returns the list of ROR IDs associated with a domain.

**Sample Request:**
```bash
curl -X POST http://localhost:8080/submit \
  -H "Content-Type: application/json" \
  -d '{
        "cmd": "search_by_domain",
        "value": "dcc.ac.uk"
      }'
```
**Sample Response**
```json
{
  ids: ["01k9d6864"],
  orgs: [
    {
    ...

    "id": "https://ror.org/01k9d6864",
    "links": [
      {
        "type": "website",
        "value": "http://www.dcc.ac.uk/"
      },
      {
        "type": "wikipedia",
        "value": "https://en.wikipedia.org/wiki/Digital_Curation_Centre"
      }
    ],

    ...

    "names": [
      {
        "lang": null,
        "types": [
          "acronym"
        ],
        "value": "DCC"
      },
      {
        "lang": "en",
        "types": [
          "ror_display",
          "label"
        ],
        "value": "Digital Curation Centre"
      }
    ],

    ...
  }
  ]
```

#### Test 2: Search by ROR ID

This returns the full JSON record of a single organization.

**Sample Request:**
```bash
curl -X POST http://localhost:8080/submit \
  -H "Content-Type: application/json" \
  -d '{
        "cmd": "search_by_ror_id",
        "value": "01k9d6864"
      }'
```

**Sample Response:**
```json
[
  {
    ...

    "id": "https://ror.org/01k9d6864",
    "links": [
      {
        "type": "website",
        "value": "http://www.dcc.ac.uk/"
      },
      {
        "type": "wikipedia",
        "value": "https://en.wikipedia.org/wiki/Digital_Curation_Centre"
      }
    ],

    ...

    "names": [
      {
        "lang": null,
        "types": [
          "acronym"
        ],
        "value": "DCC"
      },
      {
        "lang": "en",
        "types": [
          "ror_display",
          "label"
        ],
        "value": "Digital Curation Centre"
      }
    ],

    ...
  }
]
```

**Note:** You can also send an array of ROR IDs to get multiple org records in one call.

**Request format:**
```bash
{
    "cmd": "search_by_ror_id",
    "value": ["rorID1", "rorID2"]
}
```

**Response format:**
```json
[
  {
      ...
    "id": "https://ror.org/rorID1",
    "name": "First Organisation",
      ...
  },
  {
      ...
    "id": "https://ror.org/rorID2",
    "name": "Second Organisation",
      ...
  }
]
```

### License
SPDX-License-Identifier: AGPL-3.0-or-later

Copyright © 2025 Digital Curation Centre (UK) and contributors
