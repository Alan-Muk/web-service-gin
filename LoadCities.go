package main

import (
    "encoding/json"
    "fmt"
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

func main() {
    data, err := os.ReadFile("world-cities.json")
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

    var result Data

    err = json.Unmarshal(data, &result)
    if err != nil {
        fmt.Println("Error parsing JSON:", err)
        return
    }
    // Example: print all cities
    for _, city := range result.Cities {
        fmt.Printf("%s (%s): %d people\n",
            city.Name, city.Country, city.Population)
    }
}

/*
PROGRAM SUMMARY

This program reads city information from the file "world-cities.json"
and converts the JSON data into Go structures.

STRUCTURES

City
- Stores information about a city:
  - Name
  - Country
  - Population
  - Latitude
  - Longitude
  - Timezone

Data
- Represents the root JSON object.
- Contains a slice of City records.

PROGRAM FLOW

1. Read the contents of "world-cities.json".
2. Check for file-reading errors.
3. Unmarshal the JSON data into a Data struct.
4. Check for JSON parsing errors.
5. Loop through all cities in the dataset.
6. Print each city's name, country, and population.

EXAMPLE OUTPUT

Amsterdam (Netherlands): 821752 people
Rotterdam (Netherlands): 623652 people
Paris (France): 2148327 people

PURPOSE

Demonstrates how to:
- Read files in Go.
- Parse JSON using encoding/json.
- Store data in structs.
- Iterate through slices and display formatted output.
*/
