package sparsi_test

import (
	"testing"

	"github.com/amadejkastelic/spar-api/pkg/sparsi"
)

func TestSort_String(t *testing.T) {
	testCases := []struct {
		name string
		s    *sparsi.Sort
		want string
	}{
		{
			name: "nil",
			s:    nil,
			want: "",
		},
		{
			name: "empty",
			s:    &sparsi.Sort{},
			want: "",
		},
		{
			name: "field only",
			s:    &sparsi.Sort{Field: sparsi.SortFieldCreatedAt},
			want: string(sparsi.SortFieldCreatedAt),
		},
		{
			name: "field and order",
			s:    &sparsi.Sort{Field: sparsi.SortFieldPrice, Order: sparsi.OrderDesc},
			want: string(sparsi.SortFieldPrice) + ":" + string(sparsi.OrderDesc),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.s.String(); got != tc.want {
				t.Errorf("Sort.String() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestParseSort(t *testing.T) {
	testCases := []struct {
		name string
		s    string
		want *sparsi.Sort
	}{
		{
			name: "empty",
			s:    "",
			want: nil,
		},
		{
			name: "field only",
			s:    string(sparsi.SortFieldCreatedAt),
			want: &sparsi.Sort{Field: sparsi.SortFieldCreatedAt, Order: sparsi.OrderAsc},
		},
		{
			name: "field and order",
			s:    string(sparsi.SortFieldPrice) + ":" + string(sparsi.OrderDesc),
			want: &sparsi.Sort{Field: sparsi.SortFieldPrice, Order: sparsi.OrderDesc},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := sparsi.ParseSort(tc.s)
			if got == nil && tc.want == nil {
				return
			}
			if got.Field != tc.want.Field || got.Order != tc.want.Order {
				t.Errorf("ParseSort(%v) = %v, want %v", tc.s, got, tc.want)
			}
		})
	}
}

func TestFilter_String(t *testing.T) {
	testCases := []struct {
		name string
		f    *sparsi.Filter
		want string
	}{
		{
			name: "nil",
			f:    nil,
			want: "",
		},
		{
			name: "empty",
			f:    &sparsi.Filter{},
			want: "",
		},
		{
			name: "single value",
			f:    &sparsi.Filter{Name: "test", Values: []sparsi.FilterValue{{Value: "value"}}},
			want: "test:value",
		},
		{
			name: "multiple values",
			f:    &sparsi.Filter{Name: "test", Values: []sparsi.FilterValue{{Value: "value1"}, {Value: "value2"}}},
			want: "test:value1~~~value2",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.f.String(); got != tc.want {
				t.Errorf("Filter.String() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestParseFilters(t *testing.T) {
	testCases := []struct {
		name string
		s    string
		want []sparsi.Filter
	}{
		{
			name: "empty",
			s:    "",
			want: nil,
		},
		{
			name: "single filter",
			s:    "test:value1~~~value2",
			want: []sparsi.Filter{{Name: "test", Values: []sparsi.FilterValue{{Value: "value1"}, {Value: "value2"}}}},
		},
		{
			name: "multiple filters",
			s:    "test:value1~~~value2,another:value3",
			want: []sparsi.Filter{
				{Name: "test", Values: []sparsi.FilterValue{{Value: "value1"}, {Value: "value2"}}},
				{Name: "another", Values: []sparsi.FilterValue{{Value: "value3"}}},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := sparsi.ParseFilters(tc.s)
			if len(got) != len(tc.want) {
				t.Errorf("ParseFilters(%v) = %v, want %v", tc.s, got, tc.want)
				return
			}
			for i := range got {
				if got[i].Name != tc.want[i].Name || len(got[i].Values) != len(tc.want[i].Values) {
					t.Errorf("ParseFilters(%v)[%d] = %v, want %v", tc.s, i, got[i], tc.want[i])
				}
				for j := range got[i].Values {
					if got[i].Values[j].Value != tc.want[i].Values[j].Value {
						t.Errorf("ParseFilters(%v)[%d][%d] = %v, want %v", tc.s, i, j, got[i].Values[j], tc.want[i].Values[j])
					}
				}
			}
		})
	}
}

func TestPriceRangeFilter_ToFilter(t *testing.T) {
	testCases := []struct {
		name string
		f    *sparsi.PriceRangeFilter
		want *sparsi.Filter
	}{
		{
			name: "nil",
			f:    nil,
			want: nil,
		},
		{
			name: "empty",
			f:    &sparsi.PriceRangeFilter{},
			want: nil,
		},
		{
			name: "valid range",
			f:    &sparsi.PriceRangeFilter{MinPrice: 10, MaxPrice: 20},
			want: &sparsi.Filter{Name: "price", Values: []sparsi.FilterValue{{Value: "[10.00, 20.00]"}}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.f.ToFilter()
			if got == nil && tc.want == nil {
				return
			}
			if got.Name != tc.want.Name || len(got.Values) != len(tc.want.Values) {
				t.Errorf("PriceRangeFilter.ToFilter() = %v, want %v", got, tc.want)
				return
			}
			for i := range got.Values {
				if got.Values[i].Value != tc.want.Values[i].Value {
					t.Errorf("PriceRangeFilter.ToFilter()[%d] = %v, want %v", i, got.Values[i], tc.want.Values[i])
				}
			}
		})
	}
}
