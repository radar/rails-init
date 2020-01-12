package main

import (
	"runtime"

	"github.com/apex/log"
	apexCli "github.com/apex/log/handlers/cli"
	"github.com/radar/rails-init/detect"
	"github.com/radar/rails-init/mac"
	"github.com/radar/rails-init/output"
	"github.com/radar/rails-init/ubuntu"
)

func main() {

	log.SetHandler(apexCli.Default)

	if detect.IsMac() {
		mac.Install()
	} else if detect.IsUbuntu() {
		ubuntu.Install()
	} else {
		output.Fail("Unsupported OS: "+runtime.GOOS, 0)
	}

	return

	output.Success("You are now ready to use Rails!", 0)
}
