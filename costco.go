package main

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func GetCostcoResults(query Query) {
	htmlBody := getHTMLBody(query)
	defer htmlBody.Close()

	doc, err := goquery.NewDocumentFromReader(htmlBody)

	if err != nil {
		panic(err)
	}

	doc.Find(".product").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the sku
		sku, _ := s.Find("input[id^='product_sku_']").Attr("value")
		name, _ := s.Find("input[id^='product_name_']").Attr("value")
		price := strings.TrimSpace(s.Find(".price").Text())

		image_url, _ := s.Find("img").Attr("src")

		slog.Info("Found result", "looking for", query.Query, "found product", name, "sku", sku, "price", price, "image_url", image_url)
	})
}

func getHTMLBody(query Query) io.ReadCloser {
	base_url := "https://www.costco.ca/CatalogSearch?keyword=%s&costcoprogramtypes=costco-grocery&dept=All&deliveryFacetFlag=true&refine=||item_program_eligibility-ShipIt||item_program_eligibility-2DayDelivery"

	url := fmt.Sprintf(base_url, query.Query)

	// Create a new HTTP client
	client := &http.Client{}

	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}

	// costco blocks you if you do not use a human looking UA
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.81 Safari/537.36")

	// Send the request using the client
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	return resp.Body
}
