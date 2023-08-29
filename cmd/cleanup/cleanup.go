package cleanup

import (
	"github.com/fededp/dbg-go/pkg/cleanup"
	"github.com/fededp/dbg-go/pkg/root"
	"github.com/spf13/cobra"
)

func NewCleanupCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cleanup",
		Short: "Cleanup outdated dbg configs",
		RunE:  execute,
	}
	return cmd
}

func execute(c *cobra.Command, args []string) error {
	return cleanup.Run(cleanup.Options{Options: root.LoadRootOptions()}, cleanup.NewFileCleaner())
}
