package command

import (
    "github.com/spf13/cobra"
    "monitor/command/common"
)

var RoleCmd = &common.Command{
    Flags: make(map[string]interface{}, 0),
    Command: &cobra.Command{
        Use: common.CMD_ROLE,
        Short: "verion this monitor indentify",
        Long: "verion this monitor indentify",
        RunE: func(cmd *cobra.Command, args []string) error {
            // todo
            return nil
        },
    },
}