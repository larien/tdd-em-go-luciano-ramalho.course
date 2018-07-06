package main

import "testing"

func Example() {
	main()
	// Output: Please provide one or more words to search.
}

func TestExample(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for range tests {
		t.Run(tt.name, func(t *testing.T) {
			Example()
		})
	}
}
