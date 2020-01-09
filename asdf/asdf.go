package asdf

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/radar/rails-init/output"
	"github.com/radar/rails-init/runner"
)

type Tool struct {
	Name     string
	Versions []string
}

func CheckInstallation(name string, expectedVersion string) error {
	tool := Tool{Name: name}

	output.Info(fmt.Sprintf("Checking if %s (%s) is installed...", name, expectedVersion), 4)

	if !tool.checkForPlugin() {
		output.Fail(fmt.Sprintf("Could not find the %s plugin for asdf", tool.Name), 4)
		tool.installPlugin()
	}

	tool.setupSteps()
	tool.ensureInstalled(expectedVersion, false)

	return nil
}

func (tool Tool) setupSteps() {
	if tool.Name == "nodejs" {
		output.Info("Importing release team keyring for Node JS", 2)
		asdfDir := os.Getenv("ASDF_DIR")
		asdfDataDir := os.Getenv("ASDF_DATA_DIR")
		fmt.Println("asdf data dir")
		fmt.Println(asdfDataDir)
		runner.StreamWithInfo(fmt.Sprintf("ls -alH %s", asdfDir), 2)
		runner.StreamWithInfo(fmt.Sprintf("ls -alH ~/.asdf"), 2)
		runner.StreamWithInfo(fmt.Sprintf("bash %s/plugins/nodejs/bin/import-release-team-keyring", asdfDir), 2)
	}
}

func (tool Tool) installPlugin() {
	output.Info(fmt.Sprintf("Adding plugin for %s to asdf.", tool.Name), 6)
	runner.StreamWithInfo("asdf plugin-add "+tool.Name, 6)
	runner.StreamWithInfo("asdf plugin list", 6)
	runner.StreamWithInfo("asdf where "+tool.Name, 6)
}

func (tool Tool) ensureInstalled(expectedVersion string, attempted bool) error {
	asdfTool := tool.listVersions()
	if asdfTool.CheckInstalled(expectedVersion) {
		return nil
	}

	errorMsg := fmt.Sprintf("You do not have %s (%s) installed.", tool.Name, expectedVersion)
	output.Fail(errorMsg, 2)
	if attempted {
		output.Fail("Prior installation attempt failed. Please try it yourself with 'asdf install'", 6)
		return errors.New(fmt.Sprintf("Could not install %s (%s)", tool.Name, expectedVersion))
	}
	asdfTool.Install(expectedVersion)

	return tool.ensureInstalled(expectedVersion, true)
}

func (tool Tool) checkForPlugin() bool {
	pluginListOutput, _, _ := runner.Run("asdf plugin-list")
	plugins := strings.Split(strings.TrimSpace(pluginListOutput), "\n")
	pluginInstalled := false

	for _, plugin := range plugins {
		if tool.Name == plugin {
			pluginInstalled = true
		}
	}

	return pluginInstalled
}

func (tool Tool) listVersions() Tool {
	listOutput, _, _ := runner.Run("asdf list " + tool.Name)
	rawVersions := strings.Split(strings.TrimSpace(listOutput), "\n")

	versions := make([]string, len(rawVersions))
	for i, v := range rawVersions {
		versions[i] = strings.TrimSpace(v)
	}

	tool.Versions = versions
	return tool
}

func (t Tool) CheckInstalled(expectedVersion string) bool {
	for _, actualVersion := range t.Versions {
		if expectedVersion == actualVersion {
			return true
		}
	}
	return false
}

func (t Tool) Install(version string) {
	installCommand := fmt.Sprintf("asdf install %s %s", t.Name, version)
	output.Info("Attempting installation:", 4)
	runner.StreamWithInfo(installCommand, 6)
	runner.StreamWithInfo("asdf list nodejs", 6)
}
