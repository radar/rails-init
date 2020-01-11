package main

import (
	"runtime"
	"strings"

	"github.com/apex/log"
	apexCli "github.com/apex/log/handlers/cli"
	"github.com/radar/rails-init/mac"
	"github.com/radar/rails-init/output"
	"github.com/radar/rails-init/runner"
)

func main() {

	log.SetHandler(apexCli.Default)

	if runtime.GOOS == "darwin" {
		mac.Install()

	} else if runtime.GOOS == "linux" {
		stdout, _, _ := runner.Run("lsb_release -a")
		if strings.Contains(stdout, "Ubuntu") {
			ubuntu.Install()
		} else {
			output.Fail("Unsupported OS: "+runtime.GOOS, 0)
		}

		return
	}

	output.Success("You are now ready to use Rails!", 0)
}
