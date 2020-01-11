package ubuntu

import (
	"os"

	"github.com/radar/rails-init/output"
	"github.com/radar/rails-init/runner"
)

func Install() {
	err := checkForGit()
	if err != nil {
		output.Fail("Git is not installed. Please install it with `sudo apt-get install git` to before trying this again.")
		os.Exit(1)

		return
	}
}

func checkForGit() error {
	err := runner.LookPath("git")
	return err
}
