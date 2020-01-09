package tool

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/radar/rails-init/asdf"
	"github.com/radar/rails-init/output"
	"github.com/radar/rails-init/runner"
	"github.com/radar/rails-init/version"
)

type Tool struct {
	Name            string
	PackageName     string
	Executable      string
	ExpectedVersion string
	VersionCommand  string
	VersionRegexp   string
}

func (tool Tool) Install() error {
	tool.findExecutable()
	err := asdf.CheckInstallation(tool.PackageName, tool.ExpectedVersion)

	tool.setLocal()

	actualVersion, err := tool.actualVersion()
	if err != nil {
		return err
	}

	checker := version.Checker{
		tool.ExpectedVersion,
		actualVersion,
	}

	err = checker.Compare(tool.Name)
	if err != nil {
		return err
	}

	return nil
}

func (tool Tool) setLocal() {
	fmt.Println("Setting local version!!")
	command := fmt.Sprintf("asdf local %s %s", tool.PackageName, tool.ExpectedVersion)
	fmt.Println(command)
	runner.Run(command)

}

func (tool Tool) findExecutable() error {
	err := runner.LookPath(tool.VersionCommand)
	if err != nil {
		output.Fail(fmt.Sprintf("Could not find %s executable in PATH", tool.Executable), 4)
		return err
	}

	return nil
}

func (tool Tool) actualVersion() (string, error) {
	stdout, stderr, err := runner.Run(tool.VersionCommand)
	if err != nil {
		return stderr, err
	}

	re := regexp.MustCompile(tool.VersionRegexp)
	match := re.FindSubmatch([]byte(stdout))
	rubyVersion := strings.TrimSpace(string(match[1]))

	return rubyVersion, nil
}
