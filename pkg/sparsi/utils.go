package sparsi

import (
	"strings"

	"github.com/amadejkastelic/spar-api/internal/format"
	"github.com/amadejkastelic/spar-api/internal/sliceutils"
)

func (s *Sort) String() string {
	if s == nil {
		return ""
	}

	res := string(s.Field)
	if s.Order != "" {
		res += ":" + string(s.Order)
	}
	return res
}

func ParseSort(s string) *Sort {
	if s == "" {
		return nil
	}
	parts := strings.Split(s, ":")
	if len(parts) < 1 {
		return nil
	}
	order := OrderAsc
	if len(parts) == 2 && parts[1] == "desc" {
		order = OrderDesc
	}
	return &Sort{
		Field: SortField(parts[0]),
		Order: order,
	}
}

func (f *Filter) String() string {
	if f == nil || f.Name == "" || len(f.Values) == 0 {
		return ""
	}
	var values []string
	for _, v := range f.Values {
		values = append(values, v.Value)
	}
	return f.Name + ":" + strings.Join(values, "~~~")
}

func ParseFilters(s string) []Filter {
	if s == "" {
		return nil
	}
	parts := strings.Split(s, ",")
	var filters []Filter
	for _, part := range parts {
		subParts := strings.Split(part, ":")
		if len(subParts) < 2 {
			continue
		}
		name := subParts[0]
		values := strings.Split(subParts[1], "~~~")
		var filterValues []FilterValue
		for _, value := range values {
			filterValues = append(filterValues, FilterValue{Value: value})
		}
		filters = append(filters, Filter{
			Name:   name,
			Values: filterValues,
		})
	}
	return filters
}

func (f *PriceRangeFilter) ToFilter() *Filter {
	if f == nil || (f.MinPrice == 0 && f.MaxPrice == 0) || (f.MinPrice > f.MaxPrice) {
		return nil
	}
	return &Filter{
		Name: "price",
		Values: []FilterValue{{
			Value: "[" + format.FormatFloat(
				f.MinPrice,
				".",
			) + ", " + format.FormatFloat(
				f.MaxPrice,
				".",
			) + "]",
		}},
	}
}

func facetToCategory(f *Facet) *Category {
	if f == nil {
		return nil
	}
	return &Category{
		Name:       f.Name,
		FilterName: f.AssociatedFieldName,
		FilterValues: sliceutils.Map(f.Elements, func(e *FacetElement) string {
			return e.Text
		}),
	}
}
