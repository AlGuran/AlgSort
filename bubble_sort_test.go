package AlgSort

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestBubbleSortInt(t *testing.T) {
	t.Run("10 val", func(t *testing.T) {
		expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		require.Equal(t, expected, BubbleSortInt([]int{5, 3, 4, 2, 1, 10, 8, 9, 6, 7}))
	})

	t.Run("1 val", func(t *testing.T) {
		expected := []int{5}
		require.Equal(t, expected, BubbleSortInt([]int{5}))
	})

	t.Run("zeros", func(t *testing.T) {
		expected := []int{0, 0, 0, 0, 0}
		require.Equal(t, expected, BubbleSortInt([]int{0, 0, 0, 0, 0}))
	})

	t.Run("doubles", func(t *testing.T) {
		expected := []int{4, 5, 5, 6, 7, 7, 7, 8, 9, 10}
		require.Equal(t, expected, BubbleSortInt([]int{5, 4, 5, 10, 7, 7, 8, 9, 6, 7}))
	})

	t.Run("wrong numbers", func(t *testing.T) {
		expected := []int{1, 2, 3, 4, 5}
		require.Equal(t, expected, BubbleSortInt([]int{05, 03, 04, 02, 01}))
	})

	t.Run("negatives", func(t *testing.T) {
		expected := []int{-999, -8, -1, 0, 8, 999}
		require.Equal(t, expected, BubbleSortInt([]int{999, -1, 0, -999, 8, -8}))
	})

	t.Run("extremums", func(t *testing.T) {
		expected := []int{-9223372036854775807, 9223372036854775807}
		require.Equal(t, expected, BubbleSortInt([]int{9223372036854775807, -9223372036854775807}))
	})
}

func BenchmarkBubbleSortInt(b *testing.B) {
	var testSet []int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 256; i++ {
		testSet = append(testSet, rand.Intn(1000))
	}

	for i := 0; i < b.N; i++ {
		BubbleSortInt(testSet)
	}
}
