package gateway

import (
	"fmt"

	"github.com/sipeed/picoclaw/cmd/picoclaw/internal/status"
	"github.com/spf13/cobra"
)

func NewGatewayCommand() *cobra.Command {
	var debug bool

	cmd := &cobra.Command{
		Use:     "gateway",
		Aliases: []string{"g"},
		Short:   "Start picoclaw gateway",
		Args:    cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// If "status" is provided as an argument, delegate to status command
			if len(args) > 0 {
				if args[0] == "status" {
					status.StatusCmd()
					return nil
				}
				// Reject unknown arguments
				return fmt.Errorf("unknown argument: %s (did you mean 'picoclaw status'?)", args[0])
			}
			// Otherwise, start the gateway
			return gatewayCmd(debug)
		},
	}

	cmd.Flags().BoolVarP(&debug, "debug", "d", false, "Enable debug logging")

	return cmd
}
