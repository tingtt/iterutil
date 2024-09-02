package iterutil

import "iter"

func Filter[V comparable](S iter.Seq[V], matchValue V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range S {
			if v == matchValue {
				yield(v)
			}
		}
	}
}

func FilterFunc[V any](S iter.Seq[V], match func(V) bool) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range S {
			if match(v) {
				yield(v)
			}
		}
	}
}

func FilterKey[K comparable, V any](S iter.Seq2[K, V], matchKey K) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range S {
			if k == matchKey {
				yield(k, v)
			}
		}
	}
}

func FilterKeyFunc[K, V any](S iter.Seq2[K, V], match func(K) bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range S {
			if match(k) {
				yield(k, v)
			}
		}
	}
}

func FilterValue[K any, V comparable](S iter.Seq2[K, V], matchValue V) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range S {
			if v == matchValue {
				yield(k, v)
			}
		}
	}
}

func FilterValueFunc[K, V any](S iter.Seq2[K, V], match func(V) bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range S {
			if match(v) {
				yield(k, v)
			}
		}
	}
}
