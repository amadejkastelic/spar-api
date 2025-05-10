package sparsi

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/amadejkastelic/spar-api/internal/format"
	"github.com/amadejkastelic/spar-api/internal/sliceutils"
)

const searchURL = "search/products_lmos_si"

var defaultSubstringFilter = Filter{
	Name:      "pos-visible",
	Values:    []FilterValue{{Value: "81701"}},
	Substring: true,
}

// Categories lists all available categories that can be used for filtering
// products.
func (c *client) Categories() (*CategoriesResponse, error) {
	query := url.Values{}
	query.Set("substringFilter", defaultSubstringFilter.String())
	query.Set("query", "*") // needed so we obtain all facets

	req, err := c.newRequest("GET", searchURL, query, nil)
	if err != nil {
		return nil, err
	}

	var resp *SearchResponse
	status, err := c.do(req, &resp)
	if err != nil {
		return nil, err
	}
	if status.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected status: %s", status.Status)
	}

	categories := sliceutils.Map(resp.Facets, facetToCategory)

	return &CategoriesResponse{
		Values: categories,
	}, nil
}

// Search performs a search query against the Spar API.
func (c *client) Search(in SearchRequest) (*SearchResponse, error) {
	query := url.Values{}
	in.Filters = append(in.Filters, defaultSubstringFilter)
	if in.Query != "" {
		query.Set("query", in.Query)
	} else {
		query.Set("query", "*")
	}
	if in.HitsPerPage > 0 {
		query.Set("hitsPerPage", strconv.Itoa(in.HitsPerPage))
	}
	if in.Page > 0 {
		query.Set("page", strconv.Itoa(in.Page))
	}
	if in.Sort != nil {
		query.Set("sort", in.Sort.String())
	}
	if in.PriceRange != nil {
		if in.PriceRange.MinPrice > 0 {
			query.Set("price_min", format.FormatFloat(in.PriceRange.MinPrice, ","))
		}
		if in.PriceRange.MaxPrice > 0 {
			query.Set("price_max", format.FormatFloat(in.PriceRange.MaxPrice, ","))
		}

		priceRangeFilter := in.PriceRange.ToFilter()
		if priceRangeFilter != nil {
			in.Filters = append(in.Filters, *priceRangeFilter)
		}
	}
	for _, filter := range in.Filters {
		if filter.Substring {
			query.Set("substringFilter", filter.String())
		} else {
			query.Set("filter", filter.String())
		}
	}

	req, err := c.newRequest("GET", searchURL, query, nil)
	if err != nil {
		return nil, err
	}

	var resp *SearchResponse
	status, err := c.do(req, &resp)
	if err != nil {
		return nil, err
	}
	if status.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected status: %s", status.Status)
	}

	return resp, nil
}
