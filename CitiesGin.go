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