package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version  = "v0.0.0"
	revision = "HEAD"
)

func NewVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("git-org version: %s (rev: %s)\n", version, revision)
		},
	}
}
