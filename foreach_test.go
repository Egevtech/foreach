package foreach

import "testing"

func TestForEach(t *testing.T) {
	arr := []int{0, 1, 2, 3}
	tarr := []int{}

	ForEach(arr, func(_ int, d int) {
		tarr = append(tarr, d)
	})

	if len(arr) != len(tarr) {
		t.Errorf("Lengths mismath %d != %d", len(arr), len(tarr))
		t.FailNow()
	}

	for index, val := range arr {
		if val != tarr[index] {
			t.Errorf("Don't match at index %d: %d != %d", index, val, tarr[index])
		}
	}
}
