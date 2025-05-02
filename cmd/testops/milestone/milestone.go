package milestone

import (
	"github.com/soulfulnerv/qasectl/cmd/testops/milestone/create"
	"github.com/spf13/cobra"
)

// Command returns a new cobra command for milestones
func Command() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "milestone",
		Short: "Manage milestones",
	}

	cmd.AddCommand(create.Command())

	return cmd
}
