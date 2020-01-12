package ubuntu

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/radar/rails-init/output"
	"github.com/radar/rails-init/runner"
	"github.com/radar/rails-init/tool"
	"github.com/radar/rails-init/utils"
)

func Install() {
	// Need to somehow make this less ghetto.
	// A good way would be auto-installing, but it needs sudo.
	// err := checkForGit()
	// if err != nil {
	// 	output.Fail("Git is not installed. Please install it with `sudo apt-get install git` to before trying this again.", 0)
	// 	os.Exit(1)
	// 	return
	// }

	// err = checkForCurl()
	// if err != nil {
	// 	output.Fail("Curl is not installed. Please install it with `sudo apt-get install curl` to before trying this again.", 0)
	// 	os.Exit(1)
	// 	return
	// }

	checkForAsdf()
	checkForAsdfConfig()
	if os.Getenv("ASDF_DIR") != "" {
		installNode()
		installYarn()
		installRuby()
		installRails()
	} else {
		fmt.Println("ASDF installed successfully.")
		fmt.Println("Restart your terminal (or open a new tab) to continue the setup of Rails.")
		fmt.Println("Run ./rails-init to resume.")
	}

}

func checkForGit() error {
	err := runner.LookPath("git")
	return err
}

func checkForCurl() error {
	err := runner.LookPath("git")
	return err
}

func checkForAsdf() {
	output.Info("Checking for asdf...", 0)
	err := runner.LookPath("asdf")
	if err != nil {
		output.Fail("asdf is not installed. I'll install it!", 2)
		asdfDir := utils.RelativeToHome(".asdf")
		runner.StreamWithInfo(fmt.Sprintf("git clone https://github.com/asdf-vm/asdf.git %s --branch v0.7.6", asdfDir), 4)
	}
}

func checkForAsdfConfig() bool {
	config, _ := ioutil.ReadFile(utils.RelativeToHome(".bash_profile"))
	if strings.Contains(string(config), ".asdf/asdf.sh") && strings.Contains(string(config), ".asdf/completions/asdf.bash") {
		return true
	}

	bashProfile := utils.RelativeToHome(".bash_profile")
	f, _ := os.OpenFile(bashProfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	output.Info("Adding asdf configuration to ~/.bash_profile", 2)
	f.WriteString("\n. ~/.asdf/asdf.sh" + "\n. ~/.asdf/completions/asdf.bash\n")
	defer f.Close()

	return false
}

func installNode() {
	nodeVersion := utils.GetEnv("NODE_VERSION", "13.6.0")

	output.Info("Now attempting Node installation: "+nodeVersion, 0)
	output.Info("Node is used in modern Rails applications for Webpack / JavaScript assets", 2)

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

	runner.StreamWithInfo(fmt.Sprintf("gem install rails -v %s --no-document", railsVersion), 2)

}
