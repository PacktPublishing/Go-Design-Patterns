package pipelines

import "testing"

func TestLaunchPipeline(t *testing.T) {
	tableTest := [][]int{
		{3, 14},
		{5, 55},
	}

	// ...
	var res int
	for _, test := range tableTest {
		res = LaunchPipeline(test[0])
		if res != test[1] {
			t.Fatal()
		}

		t.Logf("%d == %d\n", res, test[1])
	}
}