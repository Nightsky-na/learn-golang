package cal

import "testing"

func TestSumOfFirstTree(t *testing.T) {
	given := 3
	want := 6

	get := sumOfFirst(given)
	if want != get {
		t.Errorf("given %d want %d but %d \n", given, want, get)
	}
}
