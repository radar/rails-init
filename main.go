package main

import (
	"fmt"

	"github.com/apex/log"
	apexCli "github.com/apex/log/handlers/cli"
	"github.com/radar/rails-init/output"
	"github.com/radar/rails-init/runner"
	"github.com/radar/rails-init/tool"
	"github.com/radar/rails-init/utils"
)

func installNode() {
	nodeVersion := utils.GetEnv("NODE_VERSION", "13.6.0")

	output.Info("Now attempting Node installation: "+nodeVersion, 0)
	output.Info("Node is used in modern Rails applications for Webpack / JavaScript assets", 2)

	output.Info("Before we install Node, we need to install a few homebrew packages:", 2)
	runner.StreamWithInfo("brew install coreutils gpg", 4)

	output.Info("Now we can install Node...", 2)

	node := tool.Tool{
		Name:            "Node",
		PackageName:     "nodejs",
		Executable:      "node",
		VersionCommand:  "node -v",
		VersionRegexp:   `v([\d\.]{3,})`,
		ExpectedVersion: utils.GetEnv("NODE_VERSION", "13.6.0"),
	}

	node.Install()
}

func installYarn() {
	output.Info("Now attempting Yarn installation (necessary for Webpacker)", 2)

	runner.StreamWithInfo("npm install -g yarn", 4)
}

func installRuby() {
	rubyVersion := utils.GetEnv("RUBY_VERSION", "2.6.5")

	output.Info("Now attempting Ruby installation: "+rubyVersion, 0)
	ruby := tool.Tool{
		Name:            "Ruby",
		PackageName:     "ruby",
		Executable:      "ruby",
		VersionCommand:  "ruby -v",
		VersionRegexp:   `ruby ([\d\.]{3,})`,
		ExpectedVersion: rubyVersion,
	}

	ruby.Install()
}

func installRails() {
	railsVersion := utils.GetEnv("RAILS_VERSION", "6.0.2.1")
	output.Info("Now attempting Rails installation: "+railsVersion, 0)

	runner.StreamWithInfo(fmt.Sprintf("gem install rails -v %s", railsVersion), 2)

}

func main() {

	log.SetHandler(apexCli.Default)

	installNode()
	installYarn()
	installRuby()
	installRails()

	output.Success("You are now ready to use Rails!", 0)
}
