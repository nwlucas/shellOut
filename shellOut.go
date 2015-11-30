// Package is an attempt to shamelessly create a wrapper around os exec functions so that a configurable command can be embedded in
// an application for easily configurable and prototyping of common commands that are shelled out by that application.
// It is intended for personal use in projects.

package shellOut

import (
    "errors"
    "os"
    "os/exec"
)

type CustomCommand struct {
    Name string
    Path string
    Args []string
    Env  []string
}

var c *CustomCommand

// Function calls the proper initiator for the package to work.
func init() {
    c = New()
}

// Function properly initializes each part of the containing struct.
func New() *CustomCommand {
    c := new(CustomCommand)
    c.Env = os.Environ()
    return c
}

func GetPath() (string, error) { return c.GetPath() }
func (c *CustomCommand) GetPath() (string, error) {
    // If Path is set, return that
    if c.Path != "" {
        return c.Path, nil
    }
    // If Path is not set AND Name is not empty, attempt to find the path and return that
    // or error if not found
    if c.Name != "" {
        return exec.LookPath(c.Name)
    }
    // If Path is not set AND Name is empty, return an error
    return "", errors.New("Binary name not set. Unable to look up path.")
}

// Function takes no parameters and calls the underlying os/exec function to do the work.
// It Runs the command and gives no output. Returns an error for unsuccessful runs otherwise
// returns nil.
func Run() error                    { return c.Run() }
func (c *CustomCommand) Run() error { return exec.Command(c.Name, c.Args...).Run() }

// Function takes no parameters and calls the underlying os/exec function to do the work.
// It Runs the given command and returns the output and nil for the error for successful runs,
// otherwise returns an error.
func Output() ([]byte, error) { return c.Output() }
func (c *CustomCommand) Output() ([]byte, error) {
    output, err := exec.Command(c.Name, c.Args...).Output()
    return output, err
}

// Function takes a single string paramter and adds it to the Args slice if not empty
// This function can be called multiple times to keep adding arguments
func AddArg(arg string) { c.AddArg(arg) }
func (c *CustomCommand) AddArg(arg string) {
    if arg != "" {
        c.Args = append(c.Args, arg)
    }
}
