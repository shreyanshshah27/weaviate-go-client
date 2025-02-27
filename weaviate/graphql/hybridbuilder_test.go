package graphql

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHybridBuilder_build(t *testing.T) {
	t.Run("all parameters", func(t *testing.T) {
		hybrid := HybridArgumentBuilder{}
		str := hybrid.WithQuery("query").WithVector([]float32{1, 2, 3}).WithAlpha(0.6).build()
		expected := `hybrid:{query: "query", vector: [1,2,3], alpha: 0.6}`
		require.Equal(t, expected, str)
	})

	t.Run("only query", func(t *testing.T) {
		hybrid := HybridArgumentBuilder{}
		str := hybrid.WithQuery("query").build()
		expected := `hybrid:{query: "query"}`
		require.Equal(t, expected, str)
	})

	t.Run("query and vector", func(t *testing.T) {
		hybrid := HybridArgumentBuilder{}
		str := hybrid.WithQuery("query").WithVector([]float32{1, 2, 3}).build()
		expected := `hybrid:{query: "query", vector: [1,2,3]}`
		require.Equal(t, expected, str)
	})

	t.Run("query and alpha", func(t *testing.T) {
		hybrid := HybridArgumentBuilder{}
		str := hybrid.WithQuery("query").WithAlpha(0.6).build()
		expected := `hybrid:{query: "query", alpha: 0.6}`
		require.Equal(t, expected, str)
	})
}
