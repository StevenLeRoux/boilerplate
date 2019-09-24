package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/stevenleroux/boilerplate/core"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	Run: func(cmd *cobra.Command, arguments []string) {
		fmt.Printf(projectName+" version %s %s\n", core.Version, core.Githash)
		fmt.Printf(projectName+" build date %s\n", core.BuildDate)
		fmt.Printf("go version %s %s/%s\n", runtime.Version(), runtime.GOOS, runtime.GOARCH)
	},
}
