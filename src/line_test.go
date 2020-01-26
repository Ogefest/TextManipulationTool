package main

import (
	"testing"
)

func TestReplace(t *testing.T) {
	pd := ParamDefinition{
		line: "ATEST",
	}

	params := []string{"A", "1"}
	Replace(&pd, params)
	if pd.line != "1TEST" {
		t.Errorf("Error replace")
	}
	params = []string{"T", "1"}
	Replace(&pd, params)
	if pd.line != "11ES1" {
		t.Errorf("Error replace")
	}
}

func TestFilter1(t *testing.T) {
	pd := ParamDefinition{
		line: "TEST",
	}

	params := []string{"TEST"}
	Filter(&pd, params)
	if pd.stopProcessing == true {
		t.Errorf("Error filter1")
	}

	params = []string{"ABC"}
	Filter(&pd, params)
	if pd.stopProcessing == false {
		t.Errorf("Error filter2")
	}
}

func TestIf(t *testing.T) {
	pd := ParamDefinition{
		line: "TEST",
	}

	params := []string{"TEST"}
	FIf(&pd, params)
	if pd.skipCurrentLine == true {
		t.Errorf("Error if")
	}

	params = []string{"ABC"}
	FIf(&pd, params)
	if pd.skipCurrentLine == false {
		t.Errorf("Error if")
	}
}

func TestFIfNot(t *testing.T) {
	pd := ParamDefinition{
		line: "TEST",
		// skipCurrentLine: true,
	}

	params := []string{"TEST"}
	FIfNot(&pd, params)
	if pd.skipCurrentLine == false {
		t.Errorf("Error if not")
	}

	params = []string{"ABC"}
	FIfNot(&pd, params)
	if pd.skipCurrentLine == false {
		t.Errorf("Error if not")
	}
}

func TestLower(t *testing.T) {
	pd := ParamDefinition{
		line: "TEST",
	}

	params := []string{}
	Lower(&pd, params)
	if pd.line != "test" {
		t.Errorf("Error lower")
	}
}

func TestUpper(t *testing.T) {
	pd := ParamDefinition{
		line: "test",
	}

	params := []string{}
	Upper(&pd, params)
	if pd.line != "TEST" {
		t.Errorf("Error upper")
	}
}

func TestTrim(t *testing.T) {
	pd := ParamDefinition{
		line: "  TEST  ",
	}

	params := []string{}
	Trim(&pd, params)
	if pd.line != "TEST" {
		t.Errorf("Error trim")
	}
}

func TestDuplicate(t *testing.T) {
	pd := ParamDefinition{
		line: "TEST",
	}

	params := []string{}
	Duplicate(&pd, params)
	if pd.line != "TEST\nTEST" {
		t.Errorf("Error duplicate")
	}
}

func TestPrefix(t *testing.T) {
	pd := ParamDefinition{
		line: "TEST",
	}

	params := []string{"111"}
	Prefix(&pd, params)
	if pd.line != "111TEST" {
		t.Errorf("Error prefix")
	}
}

func TestSuffix(t *testing.T) {
	pd := ParamDefinition{
		line: "TEST",
	}

	params := []string{"111"}
	Suffix(&pd, params)
	if pd.line != "TEST111" {
		t.Errorf("Error suffix")
	}
}

func TestCut(t *testing.T) {
	pd := ParamDefinition{
		line: "TEST",
	}

	params := []string{"0:2"}
	Cut(&pd, params)
	if pd.line != "TE" {
		t.Errorf("Error cut")
	}
}
