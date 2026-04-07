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
