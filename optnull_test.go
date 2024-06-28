package optnull_test

func ptr[T any](x T) *T { return &x }
