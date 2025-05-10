package sliceutils_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/amadejkastelic/spar-api/internal/sliceutils"
)

func TestMap(t *testing.T) {
	t.Run("Integer to string", func(t *testing.T) {
		input := []int{1, 2, 3}
		expected := []string{"1", "2", "3"}

		result := sliceutils.Map(input, func(i int) string {
			return fmt.Sprintf("%d", i)
		})

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})

	t.Run("String to length", func(t *testing.T) {
		input := []string{"apple", "banana", "cherry"}
		expected := []int{5, 6, 6}

		result := sliceutils.Map(input, func(s string) int {
			return len(s)
		})

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})

	t.Run("Integer pow", func(t *testing.T) {
		input := []int{1, 2, 3}
		expected := []int{1, 4, 9}

		result := sliceutils.Map(input, func(i int) int {
			return i * i
		})

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})

	t.Run("Empty slice", func(t *testing.T) {
		input := []int{}
		expected := []string{}

		result := sliceutils.Map(input, func(i int) string {
			return fmt.Sprintf("%d", i)
		})

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})
}
