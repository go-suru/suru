package cmd_test

import (
	"testing"

	"gopkg.in/suru.v0/cmd"
)

func TestHelp(t *testing.T) {
	var expect = `
 - config	Configure Suru
 - do		Do some Task
 - help (h, ?)	Help for Suru commands (with shortcuts)
 - init		Initialize Suru for a repo
 - live		Enter live mode
 - mode		Set the Mode (default Private)
 - pub		Publish to a Suru channel
 - sub		Subscribe to a Suru channel
 - task		Schedule a new Task
`[1:]

	if outs := new(cmd.Help).Help(); outs != expect {
		t.Fail()
		t.Logf("Unexpected Help output:\n%s\n"+
			"should be:\n%s\n", outs, expect)

		if len(outs) != len(expect) {
			t.Logf("Output length (%d) exceeded "+
				"expected length (%d)\n",
				len(outs), len(expect),
			)
			t.FailNow()
		}

		for i, c := range outs {
			if e := []rune(expect)[i]; e != c {
				t.Logf("expect[%d] (%#q) not matched "+
					"by output[%d] (%#q)\n",
					i, e,
					i, c,
				)
			}
		}
	}
}
