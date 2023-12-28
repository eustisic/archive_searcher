package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"milestones/constants"
	"net/http"
)

func SearchArchive() {
	fmt.Println(constants.RootUrl)
	// Implement the API request here
	// Parse the response and display it
}

func ListRecordings() {
	// var data constants.ApiResponse
	data := constants.ApiResponse{}

	fmt.Println("Listing recordings")
	resp, err := http.Get(constants.ListUrl)

	if err != nil {
		return
	}

	if resp.StatusCode >= 300 {
		fmt.Println("Status code %d recieved", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()
	if err != nil {
		fmt.Println("Error reading body", err)
		return
	}

	err = json.Unmarshal(body, &data)
	fmt.Println(data)
	if err != nil {
		fmt.Println("Error parsing JSON", err)
		return
	}
}
