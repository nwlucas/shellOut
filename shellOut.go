// Package is an attempt to shamelessly create a wrapper around os exec functions so that a configurable command can be embedded in
// an application for easily configurable and prototyping of common commands that are shelled out by that application.
// It is intended for personal use in projects.

package shellOut

import (
    "os"
    "os/exec"
)

type Cmd struct {
    Name string
    Path string
    Args []string
    Env  []string
}

var c *Cmd

// Function calls the proper initiator for the package to work.
func init() {
    c = New()
}

// Function properly initializes each part of the containing struct.
func New() *Cmd {
    c := new(Cmd)
    c.Env = os.Environ()
    return c
}

// Function takes no parameters and calls the underlying os/exec function to do the work.
// It Runs the command and gives no output. Returns an error for unsuccessful runs otherwise
// returns nil.
func Run() error          { return c.Run() }
func (c *Cmd) Run() error { return exec.Command(c.Name, c.Args...).Run() }

// Function takes no parameters and calls the underlying os/exec function to do the work.
// It Runs the given command and returns the output and nil for the error for successful runs,
// otherwise returns an error.
func Output() ([]byte, error) { return c.Output() }
func (c *Cmd) Output() ([]byte, error) {
    output, err := exec.Command(c.Name, c.Args...).Output()
    return output, err
}

// Function takes a single string paramter and adds it to the Args slice if not empty
// This function can be called multiple times to keep adding arguments
func AddArg(arg string) { c.AddArg(arg) }
func (c *Cmd) AddArg(arg string) {
    if arg != "" {
        c.Args = append(c.Args, arg)
    }
}
