package messages

import "testing"

func TestGreet(t *testing.T) {
	got := Greet("Gopher")
	expect := "Hello, Gopher!\n"

	if got != expect {
		t.Errorf("Did not get expected result. Wanted %q, got %q\n", expect, got)
	}
}

func TestGreetTableDriven(t *testing.T) {
	scenarios := []struct {
		input string
		expect string
	}{
		{input: "Gopher", expect: "Hello, Gopher!\n"},
		{input: "", expect: "Hello, !\n"},
	}

	for _, s := range scenarios {
		got := Greet(s.input)
		if got != s.expect {
			t.Errorf("Did not get expected result for input '%v'. Wanted %q, got %q", 
				s.input, got, s.expect)
		}
	}
	
}

func TestDepart(t *testing.T) {
	got := depart("Gopher")
	expect := "Goodbye, Gopher\n"

	if got != expect {
		t.Errorf("Did not get expected result. Wanted %q, got %q\n", expect, got)
	}
}

func TestFailureTypes(t *testing.T) {
	// t.Error("Error signals a failed test, but does not stop the resto of the execution")
	// t.Fatal("Fatal will fail the test and stop its execution")
	// t.Error("This will never print since it is preceeded by an immediate failure")
}
