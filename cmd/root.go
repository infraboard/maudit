package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/infraboard/maudit/cmd/bus"
	"github.com/infraboard/maudit/cmd/service"
	"github.com/infraboard/maudit/version"
)

var vers bool

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "maudit",
	Short: "审计中心",
	Long:  "审计中心",
	RunE: func(cmd *cobra.Command, args []string) error {
		if vers {
			fmt.Println(version.FullVersion())
			return nil
		}
		return cmd.Help()
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	RootCmd.AddCommand(bus.Cmd, service.Cmd)
	RootCmd.PersistentFlags().BoolVarP(&vers, "version", "v", false, "the maudit version")
}
