package cmd

import (
	"context"
	"github.com/openimsdk/chat/internal/api/chat"
	"github.com/openimsdk/chat/pkg/common/config"
	"github.com/openimsdk/tools/system/program"
	"github.com/spf13/cobra"
)

type ChatApiCmd struct {
	*RootCmd
	ctx       context.Context
	configMap map[string]any
	apiConfig chat.Config
}

func NewChatApiCmd() *ChatApiCmd {
	var ret ChatApiCmd
	ret.configMap = map[string]any{
		ShareFileName:           &ret.apiConfig.Share,
		ChatAPIChatCfgFileName:  &ret.apiConfig.ApiConfig,
		ZookeeperConfigFileName: &ret.apiConfig.ZookeeperConfig,
	}
	ret.RootCmd = NewRootCmd(program.GetProcessName(), WithConfigMap(ret.configMap))
	ret.ctx = context.WithValue(context.Background(), "version", config.Version)
	ret.Command.RunE = func(cmd *cobra.Command, args []string) error {
		return ret.preRunE()
	}
	return &ret
}

func (a *ChatApiCmd) Exec() error {
	return a.Execute()
}

func (a *ChatApiCmd) preRunE() error {
	return chat.Start(a.ctx, a.Index(), &a.apiConfig)
}
