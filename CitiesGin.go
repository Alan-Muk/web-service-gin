package main

import (
    "encoding/json"
    "github.com/gin-gonic/gin"
    "os"
)

type City struct {
    Name       string  `json:"name"`
    Country    string  `json:"country"`
    Population int     `json:"population"`
    Lat        float64 `json:"lat"`
    Lon        float64 `json:"lon"`
    Timezone   string  `json:"timezone"`
}

type Data struct {
    Cities []City `json:"cities"`
}

var cities []City

func loadData() {
    file, err := os.ReadFile("world-cities.json")
    if err != nil {
        panic(err)
    }

    var data Data
    if err := json.Unmarshal(file, &data); err != nil {
        panic(err)
    }

    cities = data.Cities
}

func main() {
    loadData()

    r := gin.Default()
    
    r.GET("/search", func(c *gin.Context) {
    country := c.Query("country")

    var result []City
    for _, city := range cities {
        if city.Country == country {
            result = append(result, city)
        }
    }

    c.JSON(200, result)
    })

    r.Run(":8080")

}

/*
PROGRAM OVERVIEW

- Loads city data from "world-cities.json" into memory when the application starts.
- Uses the Gin web framework to expose an HTTP API.
- Provides a GET /search endpoint that accepts a "country" query parameter.
- Searches through the in-memory city dataset and returns cities that match the specified country.
- Responds with a JSON array containing the matching city records.
- Runs an HTTP server on port 8080.

DATA STRUCTURES

City:
- Represents a city record.
- Contains name, country, population, coordinates, and timezone information.

Data:
- Represents the root JSON object.
- Contains a list of City records.

FUNCTIONS

loadData():
- Reads the JSON file from disk.
- Unmarshals the JSON into Go structs.
- Stores the loaded cities in the global cities slice.

main():
- Loads the dataset.
- Creates and configures the Gin router.
- Registers the /search endpoint.
- Starts the web server on port 8080.

ENDPOINTS

GET /search?country=<country>

Example:
GET /search?country=Netherlands

Response:
[
  {
    "name": "Amsterdam",
    "country": "Netherlands",
    "population": 821752,
    "lat": 52.374,
    "lon": 4.8897,
    "timezone": "Europe/Amsterdam"
  }
]
*/
