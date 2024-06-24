# star-citizen

Handlers

AddPlanet: POST /planets

GetPlanets: GET /planets/:id?

UpdatePlanet: PUT /planets/:id

DeletePlanet: DELETE /planets/:id

GetFuelEstimate: GET /planets/:id/fuel

Repositories
Repositories actually store the data
Planet Repository is a interface
and InMemoryPlanetRepo is its implementation

Util
ValidatePlanet: Validates planet data according to the requirements
CalculateFuel: Estimates fuel required

Run
go run main.go

Build the Docker image:

docker build -t star-citizen-app .

Run the Docker container:
docker run -p 3000:3000 star-citizen-app

Testing
go test ./...