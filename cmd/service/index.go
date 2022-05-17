package service

import (
	"github.com/spf13/cobra"
)

var (
	// pusher service config option
	confType string
	confFile string
	confETCD string
)

var Cmd = &cobra.Command{
	Use:   "service",
	Short: "服务管理",
	Long:  `服务管理`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {
	Cmd.PersistentFlags().StringVarP(&confType, "config-type", "t", "file", "the service config type [file/env/etcd]")
	Cmd.PersistentFlags().StringVarP(&confFile, "config-file", "f", "etc/maudit.toml", "the service config from file")
	Cmd.PersistentFlags().StringVarP(&confETCD, "config-etcd", "e", "127.0.0.1:2379", "the service config from etcd")
}
