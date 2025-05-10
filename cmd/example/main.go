package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/amadejkastelic/spar-api/pkg/sparsi"
)

var (
	queryPtr     = flag.String("query", "*", "Search query")
	pageSizePtr  = flag.Int("page-size", 10, "Number of results per page")
	pagePtr      = flag.Int("page", 0, "Page number")
	userAgentPtr = flag.String(
		"user-agent",
		"Mozilla/5.0 (X11; Linux i686; rv:128.5) Gecko/20100101 Firefox/128.5",
		"User agent string",
	)
	timeoutPtr  = flag.Int("timeout", 15, "Timeout in seconds")
	sortPtr     = flag.String("sort", "", "Sort order (e.g., 'best-price:asc')")
	filtersPtr  = flag.String("filters", "", "Filters (e.g., 'ecr-brand:3M,is-on-promotion:true')")
	minPricePtr = flag.Float64("min-price", 0, "Minimum price")
	maxPricePtr = flag.Float64("max-price", 0, "Maximum price")
)

func main() {
	flag.Parse()

	// Create a new client with custom timeout
	client := sparsi.NewClient(
		sparsi.WithTimeout(time.Duration(*timeoutPtr)*time.Second),
		sparsi.WithUserAgent(*userAgentPtr),
	)

	// Perform a search
	resp, err := client.Search(sparsi.SearchRequest{
		Query:       *queryPtr,
		HitsPerPage: *pageSizePtr,
		Page:        *pagePtr,
		Sort:        sparsi.ParseSort(*sortPtr),
		Filters:     sparsi.ParseFilters(*filtersPtr),
		PriceRange: &sparsi.PriceRangeFilter{
			MinPrice: *minPricePtr,
			MaxPrice: *maxPricePtr,
		},
	})
	if err != nil {
		log.Fatalf("Error searching: %v", err)
	}

	fmt.Printf("Found %d products (total: %d)\n", len(resp.Hits), resp.TotalHits)

	for i, hit := range resp.Hits {
		fmt.Printf("%d. %s - %.2f\n", i+1, hit.Product.Name, hit.Product.Price)
	}
}
