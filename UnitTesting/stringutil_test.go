package stringutil

import "testing"

//How to run
// Just cd to UnitTesting folder and run the below command, it identifies *_test.go files and run the unit tests
// > go test



func TestReverseString_If_Strings_AreNotEqual(t *testing.T){
	cases :=  []struct{
		in,want string
	}{
		{"testing","gniset"},
	}

	for _,c := range cases{
		got := Reverse(c.in)
		if got == c.want {
			t.Errorf("Reverse (%q) == %q, want %q", c.in,got,c.want)
		}
	}
}

func TestReverseString_If_Strings_AreEqual(t *testing.T){
	cases :=  []struct{
		in,want string
	}{
		{"Hello world","dlrow olleH"},
		{"Hello, 世界", "界世 ,olleH"},
	}

	for _,c := range cases{
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse (%q) == %q, want %q", c.in,got,c.want)
		}
	}
}

func TestReverseString_If_Strings_AreNonEnglish(t *testing.T){
	cases :=  []struct{
		in,want string
	}{
		{"Hello, 世界", "界世 ,olleH"},
	}

	for _,c := range cases{
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse (%q) == %q, want %q", c.in,got,c.want)
		}
	}
}