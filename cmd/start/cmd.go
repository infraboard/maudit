package start

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/infraboard/mcenter/clients/rpc/hooks"
	"github.com/infraboard/mcube/ioc/config/application"

	// 注册所有服务
	_ "github.com/infraboard/maudit/apps"
)

// Cmd represents the start command
var Cmd = &cobra.Command{
	Use:   "start",
	Short: "maudit API服务",
	Long:  "maudit API服务",
	Run: func(cmd *cobra.Command, args []string) {
		hooks.NewMcenterAppHook().SetupAppHook()
		cobra.CheckErr(application.App().Start(context.Background()))
	},
}
