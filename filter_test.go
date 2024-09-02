package iterutil_test

import (
	"iter"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tingtt/iterutil"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) Func() {
	m.Called()
}

func TestFilter(t *testing.T) {
	t.Run("may filter", func(t *testing.T) {
		t.Parallel()

		list := []string{"a", "b", "c"}
		for item := range iterutil.Filter(slices.Values(list), "a") {
			assert.Equal(t, item, "a")
		}
	})

	t.Run("may filter", func(t *testing.T) {
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
		for range iterutil.Filter(iter, "a") {
			mock2.Func()
		}

		mock.AssertNumberOfCalls(t, "Func", len(list))
		mock2.AssertNumberOfCalls(t, "Func", 1)
	})
}
