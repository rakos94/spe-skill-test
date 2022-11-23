package trial

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Post struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserId int    `json:"userId"`
}

func Trial() {

	// Encode the data
	postBody, _ := json.Marshal(map[string]string{
		"resource_type": "get_weather",
	})
	responseBody := bytes.NewBuffer(postBody)
	req, err := http.NewRequest(http.MethodPost, "https://www.spesolution.com/trial/get-weather", responseBody)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer OtQxKb09kfN7ZSaO3mX8qJXS0BLwupix")
	req.Header.Add("x-api-key", "a2FtYWx3bzVqNjpzbUl1SERsOVE2T3d6eGNM")

	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}
	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))

}
