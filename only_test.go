package iterutil_test

import (
	"iter"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tingtt/iterutil"
)

func TestOnly(t *testing.T) {
	t.Run("may yield only matched", func(t *testing.T) {
		t.Parallel()

		list := []string{"a", "b", "c"}
		for item := range iterutil.Only(slices.Values(list), "a") {
			assert.Equal(t, item, "a")
		}
	})

	t.Run("may yield only", func(t *testing.T) {
		t.Parallel()

		list := []string{"a", "b", "c"}
		mock := new(Mock)
		mock.On("Func")

		iter := func() iter.Seq[string] {
			return func(yield func(string) bool) {
				for _, v := range list {
					mock.Func()
					if !yield(v) {
						break
					}
				}
			}
		}()

		mock2 := new(Mock)
		mock2.On("Func")
		for range iterutil.Only(iter, "b") {
			mock2.Func()
		}

		mock.AssertNumberOfCalls(t, "Func", 2)
		mock2.AssertNumberOfCalls(t, "Func", 1)
	})
}
