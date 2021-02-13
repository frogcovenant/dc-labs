package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"math"
)

type Point struct {
	X, Y float64
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//generatePoints array
func generatePoints(s string) ([]Point, error) {

	points := []Point{}

	s = strings.Replace(s, "(", "", -1)
	s = strings.Replace(s, ")", "", -1)
	vals := strings.Split(s, ",")
	if len(vals) < 2 {
		return []Point{}, fmt.Errorf("Point [%v] was not well defined", s)
	}

	var x, y float64

	for idx, val := range vals {

		if idx%2 == 0 {
			x, _ = strconv.ParseFloat(val, 64)
		} else {
			y, _ = strconv.ParseFloat(val, 64)
			points = append(points, Point{x, y})
		}
	}
	return points, nil
}

// getArea gets the area inside from a given shape
func getArea(points []Point) float64 {
	area := 0.0
	points = append(points, points[0]) //add last point to array
	for i := 0; i < len(points)-1; i++{
		area += ((points[i].Y + points[i+1].Y)/2) * (points[i+1].X-points[i].X)
	}
	return (math.Round(area*10)/10)
}

// getPerimeter gets the perimeter from a given array of connected points
func getPerimeter(points []Point) float64 {
	perimeter := 0.0
	points = append(points, points[0]) //add last point to array
	for i := 0; i < len(points)-1; i++{
		perimeter += math.Sqrt(math.Pow((points[i+1].X-points[i].X),2)+math.Pow((points[i+1].Y-points[i].Y),2))
	}
	return (math.Round(perimeter*10)/10)
}

// handler handles the web request and reponds it
func handler(w http.ResponseWriter, r *http.Request) {

	var vertices []Point
	for k, v := range r.URL.Query() {
		if k == "vertices" {
			points, err := generatePoints(v[0])
			if err != nil {
				fmt.Fprintf(w, fmt.Sprintf("error: %v", err))
				return
			}
			vertices = points
			break
		}
	}

	// Results gathering
	area := getArea(vertices)
	perimeter := getPerimeter(vertices)

	// Logging in the server side
	log.Printf("Received vertices array: %v", vertices)

	// Response construction
	response := fmt.Sprintf("Welcome to the Remote Shapes Analyzer\n")
	response += fmt.Sprintf(" - Your figure has : [%v] vertices\n", len(vertices))
	response += fmt.Sprintf(" - Vertices        : %v\n", vertices)
	if len(vertices) > 2{
		response += fmt.Sprintf(" - Perimeter       : %v\n", perimeter)
		response += fmt.Sprintf(" - Area            : %v\n", area)
	}else{
		response += fmt.Sprintf(" - ERROR - Your shape is not compliying with the minimum number of vertices.\n")
	}

	// Send response to client
	fmt.Fprintf(w, response)
}
