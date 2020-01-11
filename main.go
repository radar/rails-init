package main

import (
	"runtime"

	"github.com/apex/log"
	apexCli "github.com/apex/log/handlers/cli"
	"github.com/radar/rails-init/mac"
	"github.com/radar/rails-init/output"
)

func main() {

	log.SetHandler(apexCli.Default)

	if runtime.GOOS == "darwin" {
		mac.Install()
	} else {
		output.Fail("Unsupported OS: " + runtime.GOOS)
		return
	}

	output.Success("You are now ready to use Rails!", 0)
}
