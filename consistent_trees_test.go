package consistent_trees

import (
	"fmt"
	"sort"
	"testing"
)

var files = []string {
	"/project/surph/diemer/Box_L0063_N1024_CBol/Rockstar/trees/tree_0_0_0.dat",
	"/project/surph/diemer/Box_L0063_N1024_CBol/Rockstar/trees/tree_0_0_1.dat",
	"/project/surph/diemer/Box_L0063_N1024_CBol/Rockstar/trees/tree_0_1_0.dat",
	"/project/surph/diemer/Box_L0063_N1024_CBol/Rockstar/trees/tree_0_1_1.dat",
	"/project/surph/diemer/Box_L0063_N1024_CBol/Rockstar/trees/tree_1_0_0.dat",
	"/project/surph/diemer/Box_L0063_N1024_CBol/Rockstar/trees/tree_1_0_1.dat",
	"/project/surph/diemer/Box_L0063_N1024_CBol/Rockstar/trees/tree_1_1_0.dat",
	"/project/surph/diemer/Box_L0063_N1024_CBol/Rockstar/trees/tree_1_1_1.dat",
}

func TestReadHalos(t *testing.T) {
	for _, file := range files {
		ReadTree(file)
		fmt.Println("Halos Read.")
		list := FindClosestScale(1.0)
		n := list.NumHalos()
		
		ms := make([]float64, n)
		for i := range ms {
			ms[i] = list.Halos(i).MVir()
		}
		sort.Float64Slice(ms).Sort()
		fmt.Println(n)
		fmt.Printf("%.4g %.4g\n", ms[0], ms[len(ms) - 1])

		DeleteTree()
	}
	fmt.Println("Everything worked?!?")
}
