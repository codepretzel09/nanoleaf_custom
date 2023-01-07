package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Seed the random number generator with the current Unix time in nanoseconds
	rand.Seed(time.Now().UnixNano())
	// Get the panel layout and set the colors
	getpanels()
}

// Generate a random color value between 0 and 255 (inclusive)
func generateRandomColorValue() int {
	return rand.Intn(256)
}

// PositionData represents the data for a panel's position
type PositionData struct {
	PanelID int `json:"panelId"`
	X       int `json:"x"`
	Y       int `json:"y"`
	O       int `json:"o"`
	Shape   int `json:"shapeType"`
}

// Response represents the data returned in the response to a panel layout request
type Response struct {
	NumPanels    int            `json:"numPanels"`
	SideLength   int            `json:"sideLength"`
	PositionData []PositionData `json:"positionData"`
}

// getpanels gets the panel layout and sets the colors for each panel
func getpanels() {
	// URL for the panel layout endpoint
	url := "http://192.168.x.x:16021/api/v1/AUTHTOKEN/panelLayout/layout"
	// HTTP method for the request
	method := "GET"
	// Create a new HTTP client
	client := &http.Client{}
	// Create a new HTTP request with the specified method and URL
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	// Send the request and get the response
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Close the response body when the function returns
	defer res.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Unmarshal the response body into a Response struct
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Iterate over each panel in the PositionData slice
	for _, panel := range response.PositionData {
		// Convert the PanelID to a string
		panelIDStr := strconv.Itoa(panel.PanelID)

		// Generate a random color
		red := generateRandomColorValue()
		green := generateRandomColorValue()
		blue := generateRandomColorValue()

		// Convert the color values to strings
		redStr := strconv.Itoa(red)
		greenStr := strconv.Itoa(green)
		blueStr := strconv.Itoa(blue)

		fmt.Println(panelIDStr)
		fmt.Println(redStr)
		fmt.Println(greenStr)
		fmt.Println(blueStr)
		makeHTTPCall("http://192.168.x.x:16021/api/v1/AUTHTOKEN/effects", "PUT", `{"write":{"animData":"1 `+panelIDStr+` 1 `+redStr+` `+greenStr+` `+blueStr+` 0 1","animType":"static","command":"display","loop":false,"palette":[],"version":"1.0"}}`)
	}

}

// Function to make an HTTP call with the specified URL, method, and payload
func makeHTTPCall(url string, method string, payload string) {
	// Create a new HTTP client
	client := &http.Client{}
	// Create a new HTTP request with the specified method, URL, and payload
	req, err := http.NewRequest(method, url, strings.NewReader(payload))
	// Check if there was an error creating the request
	if err != nil {
		fmt.Println(err)
		return
	}
	// Send the HTTP request and get the response
	res, err := client.Do(req)
	// Check if there was an error sending the request
	if err != nil {
		fmt.Println(err)
		return
	}
	// Close the response body when the function returns
	defer res.Body.Close()
	// Read the response body
	body, err := ioutil.ReadAll(res.Body)
	// Check if there was an error reading the response body
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the response body
	fmt.Println(string(body))
}
