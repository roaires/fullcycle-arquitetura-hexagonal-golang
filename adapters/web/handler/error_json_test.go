package handler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandler_jsonError(t *testing.T) {
	msg := "Test"
	result := jsonError(msg)
	require.Equal(t, string([]byte(`{"message":"Test"}`)), string(result))
}
