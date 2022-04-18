package piper

import "testing"

func TestDo(t *testing.T) {
	res := Do("ls -l").Then("wc -l").Then("uniq").String()
	expected := "ls -l | wc -l | uniq"
	if res != expected {
		t.Error("String() output not as expected")
	}
}
