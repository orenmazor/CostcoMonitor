package main

import (
	"fmt"
	"io"
	"net/http"
)

func GetCostcoResults(query Query) {
	base_url := "https://www.costco.ca/CatalogSearch?keyword=%s&costcoprogramtypes=costco-grocery&dept=All&deliveryFacetFlag=true&refine=||item_program_eligibility-ShipIt||item_program_eligibility-2DayDelivery"

	url := fmt.Sprintf(base_url, query.Query)

	// Create a new HTTP client
	client := &http.Client{}

	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}

	// costco blocks you if you do not use a human looking UA
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.81 Safari/537.36")

	// Send the request using the client
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Print the response body
	fmt.Println(string(body))
}
