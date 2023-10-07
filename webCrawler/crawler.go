package webCrawler

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func Crawl(url string) error {
	body, err := Fetch(url)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("%s : %s", err, url)
	}

	campusInformationElement := doc.Find("h3.wp-block-heading:contains('Campus Information')")
	campusInformationText := campusInformationElement.Text()
	fmt.Println(campusInformationText)

	fmt.Printf("found: %s %q\n", url, body)
	return nil
}

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("%s : %s", err, url)
	}
	if resp.StatusCode != 200 {
		fmt.Println("Request failed with status code", resp.StatusCode)
		return nil, fmt.Errorf("Request failed with status code %d: %s", resp.StatusCode, url)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("%s : %s", err, url)
	}
	return body, nil
}
