package main

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestLint(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, analyzer, "decl", "assign", "strct", "imprt", "main")
}