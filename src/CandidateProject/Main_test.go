/*

Created By: Jeff Sadler

Note: Govatar library must be installed on local PC for this to work!
	  See https://github.com/o1egl/govatar
	  go get -u github.com/o1egl/govatar/...

This file contains the test units and benchmarks for the wrappers defined
in Main.go.

Command lines for executing tests and benchmarks

Run unit tests (verbose)
  go test -v CandidateProject

Run unit tests and calc code coverage and put result to file results.txt
  go test -v CandidateProject -cover > results.txt

Run benchmarks and append results to results.txt
  go test CandidateProject -bench="." >> results.txt

*/

package main

import (
	"fmt"
	"testing"

	"github.com/o1egl/govatar"
)

// TestCanGenerateGovatarMale is the test for function GovatarGenerateMale
func TestCanGenerateGovatarMale(t *testing.T) {
	if _, err := GovatarGenerateMale(); err != nil {
		t.Log("Failed to generate male Govatar.")
		t.Fail()
	}
}

// TestCanGenerateGovatarFemale is the test for function GovatarGenerateFemale
func TestCanGenerateGovatarFemale(t *testing.T) {
	if _, err := GovatarGenerateFemale(); err != nil {
		t.Log("Failed to generate female Govatar.")
		t.Fail()
	}
}

// TestCanGenerateGovatar is the test for function GovatarGenerate
func TestCanGenerateGovatar(t *testing.T) {
	var gender govatar.Gender

	gender = govatar.MALE
	if _, err := GovatarGenerate(gender); err != nil {
		t.Log(fmt.Errorf("Failed to generate Govatar with gender = %d", gender))
		t.Fail()
	}

	gender = govatar.FEMALE
	if _, err := GovatarGenerate(gender); err != nil {
		t.Log(fmt.Errorf("Failed to generate Govatar with gender = %d", gender))
		t.Fail()
	}

}

// BenchmarkGovatarGenerateMale is the benchmark test for function BenchmarkGovatarGenerateMale
func BenchmarkGovatarGenerateMale(b *testing.B) {
	// Let's report memory allocations
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		GovatarGenerateMale()
	}
}

// BenchmarkGovatarGenerateFemale is the benchmark test for function BenchmarkGovatarGenerateFemale
func BenchmarkGovatarGenerateFemale(b *testing.B) {
	// Let's report memory allocations
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		GovatarGenerateFemale()
	}
}

// BenchmarkGovatarGenerate is the benchmark test for function BenchmarkGovatarGenerate
func BenchmarkGovatarGenerate(b *testing.B) {
	// Let's report memory allocations
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		GovatarGenerate(govatar.MALE)
	}
}
