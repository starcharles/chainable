package chaninable

type chainable[T any] struct {
	values []T
}

func NewChaninable[T any](values []T) *chainable[T] {
	return &chainable[T]{
		values: values,
	}
}

func (c *chainable[T]) Filter(f func(T) bool) *chainable[T] {

	results := make([]T, 0)
	for i, v := range c.values {

		if f(v) {
			results = append(results, c.values[i])
		}
	}
	return NewChaninable(results)
}

// func (c *Chainable[T]) Map(f func(T) U) *Chainable[U] {
// 	results := make([]U, len(c.values))
// 	for i, v := range c.values {
// 		results[i] = f(v)
// 	}
// 	return NewChainable(results)
// }
// func (c *chainable[T]) Map(f func(T) any) *chainable{
// 	return Map(c.values, f)

// }

func (c *chainable[T]) Values() []T {
	return c.values
}

// func Map[T, U any](values []T, f func(T) U) *chainable[U] {
// 	results := make([]U, len(values))
// 	for i, v := range values {
// 		results[i] = f(v)
// 	}
// 	return NewChaninable(results)
// }
