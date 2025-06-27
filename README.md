## 🚀 Getting Started

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
```

Then inside the container:

`go run api.go`

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
```
["01k9d6864"]
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

This project is licensed under the **GNU Affero General Public License v3.0**.

You may copy, modify, and distribute this program under the terms of the license as published by the Free Software Foundation.

See the full license text here: [GNU AGPL v3.0](https://www.gnu.org/licenses/agpl-3.0.html)
