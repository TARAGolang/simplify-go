package simplifier

import (
	"log"
	"testing"
)

func TestSimplifier(t *testing.T) {
	orig, simpl_5f, simpl_3t := GetTestData()

	simpl_5f_test := Simplify(orig, 5, false)
	simpl_3t_test := Simplify(orig, 3, true)

	if !CompareSlices(simpl_5f_test, simpl_5f) {
		t.Fatalf(" 5 false | Something went wrong")
	}

	if !CompareSlices(simpl_3t_test, simpl_3t) {
		t.Fatalf("3 true | Something went wrong")
	}

}

func TestSimplex(t *testing.T) {

	testData := [][]float64{{1, 2}, {2, 4}, {3, 5}}

	res := Simplify(testData, 3, false)
	log.Println()
}
