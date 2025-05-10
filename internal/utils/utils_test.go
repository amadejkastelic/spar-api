package utils_test

import (
	"testing"

	"github.com/amadejkastelic/spar-api/internal/utils"
)

type mockCloser struct {
	closed bool
}

func (m *mockCloser) Close() error {
	m.closed = true
	return nil
}

func TestCloseQuetly(t *testing.T) {
	mockCloser := &mockCloser{}
	utils.CloseQuetly(mockCloser)
	if !mockCloser.closed {
		t.Errorf("Expected Close to be called, but it wasn't")
	}

	// Test with a non-io.Closer type
	utils.CloseQuetly("not a closer")
}
