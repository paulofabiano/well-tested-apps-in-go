// Interesting external projects:
// Testify - github.com/stretchrcom/testify
// Ginkgo - github.com/onsi/ginkgo (BDD)
// GoConvey - goconvey.co
// httpexpect - github.com/gavv/httpexpect
// gomock - code.google.com/p/gomock
// go-sqlmock - github.com/DATA-DOG/go-sqlmock (in-memory sql DB)

package main_test

import "testing"

func TestAddition(t *testing.T) {
	got := 2 + 2
	expected := 4
	if got != expected {
		t.Errorf("Did not get expected result. Got: '%v', wanted: '%v'", got, expected)
	}
}

func TestSubtraction(t *testing.T) {
	got := 10 - 5
	expected := 4
	if got != expected {
		t.Errorf("Did not get expected result. Got: '%v', wanted: '%v'", got, expected)
	}
}
