package qrng

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Define the structure to match the API response
type QRNGResponse struct {
	Type    string `json:"type"`
	Length  int    `json:"length"`
	Data    []int  `json:"data"`
	Success bool   `json:"success"`
}

func GetQuantumRandomNumbers(length int) ([]int, error) {
	// URL for ANU QRNG
	url := fmt.Sprintf("https://qrng.anu.edu.au/API/jsonI.php?length=%d&type=uint8", length)

	// Create a new HTTP request
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check HTTP status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parse the JSON response
	var qrngResponse QRNGResponse
	err = json.Unmarshal(body, &qrngResponse)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %v", err)
	}

	// Return the random data
	return qrngResponse.Data, nil
}


func main() {
	randomNumbers, err := getQuantumRandomNumbers()
	if err != nil {
		fmt.Println("Error fetching QRNG data:", err)
		return
	}
	fmt.Println("Quantum Random Numbers:", randomNumbers)
}
