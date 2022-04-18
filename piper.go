package piper

import (
	"fmt"
)

// Pipe input output manager
type Pipe struct {
	stdin  string
	stdout string
}

// Do return pipe instance
func Do(in string) *Pipe {
	return &Pipe{
		stdin:  "",
		stdout: in,
	}
}

// Then of pipes
func (c *Pipe) Then(in string) *Pipe {
	res := fmt.Sprintf("%s | %s ", c.stdout, in)
	return &Pipe{
		stdin:  in,
		stdout: res,
	}
}

func (c *Pipe) String() string {
	return fmt.Sprintf("%v", c.stdout)
}

// Output return built pipeline
func (c *Pipe) Output() {
	fmt.Printf("%v\n", c.stdout)
}
