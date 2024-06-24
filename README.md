# Star Citizen

## Handlers

### AddPlanet
adds a new planet
- **Method:** `POST`
- **Endpoint:** `/planets`

### GetPlanets
get all planets

if id is provided one specific planet with id will be fetched
else all planets will be fetched
and following filters can be used in query parameters when
fetching all planets

minMass <br />
maxMass <br />
minRadius <br />
maxRadius <br />

- **Method:** `GET`
- **Endpoint:** `/planets/:id?`

### UpdatePlanet
details of of planet with :id will be updated
- **Method:** `PUT`
- **Endpoint:** `/planets/:id`

### DeletePlanet
planet with :id will be deleted
- **Method:** `DELETE`
- **Endpoint:** `/planets/:id`

### GetFuelEstimate
fuel estimation for going to planet with id will be calculated
- **Method:** `GET`
- **Endpoint:** `/planets/:id/fuel`

## Repositories

Repositories handle data storage.

- **Planet Repository:** Interface for managing planet data.
- **InMemoryPlanetRepo:** Implementation using in-memory storage.

## Utilities

- **ValidatePlanet:** Validates planet data according to specified requirements.
- **CalculateFuel:** Estimates fuel required for a planet.

## Run

To run the application:

```bash
go run main.go
```

### Build the Docker image:

docker build -t star-citizen-app .

### Run the Docker container: 
```bash
docker run -p 3000:3000 star-citizen-app
```

### Testing
```bash
go test ./...
```