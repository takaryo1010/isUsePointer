package isUsePointer_test

import (
	"testing"

	"isUsePointer"

	"github.com/gostaticanalysis/testutil"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	t.Run("a", func(t *testing.T) {
		analysistest.Run(t, testdata, isUsePointer.Analyzer, "a")
	})
	t.Run("methods", func(t *testing.T) {
		analysistest.Run(t, testdata, isUsePointer.Analyzer, "methods")
	})
	t.Run("noMethods", func(t *testing.T) {
		analysistest.Run(t, testdata, isUsePointer.Analyzer, "noMethods")
	})
	t.Run("noPointer", func(t *testing.T) {
		analysistest.Run(t, testdata, isUsePointer.Analyzer, "noPointer")
	})
	t.Run("funcs", func(t *testing.T) {
		analysistest.Run(t, testdata, isUsePointer.Analyzer, "funcs")
	})
	t.Run("methods_reverse", func(t *testing.T) {
		analysistest.Run(t, testdata, isUsePointer.Analyzer, "methods_reverse")
	})
}
