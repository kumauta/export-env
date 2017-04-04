package main

import (
	"code.cloudfoundry.org/cli/plugin"
	"fmt"
	"os"
	"syscall"
)

type ExportEnvPlugin struct{}

func (c *ExportEnvPlugin) Run(cliConnection plugin.CliConnection, args []string) {
	if args[0] == "export-env" || args[0] == "ee" {

		un, _ := cliConnection.Username()
		ug, _ := cliConnection.UserGuid()
		ue, _ := cliConnection.UserEmail()
		ae, _ := cliConnection.ApiEndpoint()
		av, _ := cliConnection.ApiVersion()
		le, _ := cliConnection.LoggregatorEndpoint()
		de, _ := cliConnection.DopplerEndpoint()
		at, _ := cliConnection.AccessToken()

		os.Setenv("CF_USER_NAME", un)
		os.Setenv("CF_USER_GUID", ug)
		os.Setenv("CF_USER_EMAIL", ue)
		os.Setenv("CF_API_ENDPOINT", ae)
		os.Setenv("CF_API_VERSION", av)
		os.Setenv("CF_LOGGREGATOR_ENDPOINT", le)
		os.Setenv("CF_DOPPLER_ENDPOINT", de)
		os.Setenv("CF_ACCESS_TOKEN", at)

		fmt.Println("export env ...\n")
		fmt.Println("env\t\t\tvalue")
		fmt.Println("CF_USER_NAME\t\t" + un)
		fmt.Println("CF_USER_GUID\t\t" + ug)
		fmt.Println("CF_USER_EMAIL\t\t" + ue)
		fmt.Println("CF_API_ENDPOINT\t\t" + ae)
		fmt.Println("CF_API_VERSION\t\t" + av)
		fmt.Println("CF_LOGGREGATOR_ENDPOINT\t" + le)
		fmt.Println("CF_DOPPLER_ENDPOINT\t" + de)
		fmt.Println("CF_ACCESS_TOKEN\t\t" + at)

		syscall.Exec(os.Getenv("SHELL"), []string{os.Getenv("SHELL")}, syscall.Environ())

	}
}

func (c *ExportEnvPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "ExportEnvPlugin",
		Version: plugin.VersionType{
			Major: 1,
			Minor: 0,
			Build: 0,
		},
		MinCliVersion: plugin.VersionType{
			Major: 6,
			Minor: 7,
			Build: 0,
		},
		Commands: []plugin.Command{
			{
				Name:     "export-env",
				Alias:    "ee",
				HelpText: "Export CF Environment",

				UsageDetails: plugin.Usage{
					Usage: "export-env,ee\n   cf export-env",
				},
			},
		},
	}
}

func main() {
	plugin.Start(new(ExportEnvPlugin))
}
