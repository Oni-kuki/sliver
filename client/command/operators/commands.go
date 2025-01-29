package operators

import (
	"github.com/Oni-kuki/sliver/client/command/flags"
	"github.com/Oni-kuki/sliver/client/command/help"
	"github.com/Oni-kuki/sliver/client/console"
	consts "github.com/Oni-kuki/sliver/client/constants"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// Commands returns the “ command and its subcommands.
func Commands(con *console.SliverClient) []*cobra.Command {
	operatorsCmd := &cobra.Command{
		Use:   consts.OperatorsStr,
		Short: "Manage operators",
		Long:  help.GetHelpFor([]string{consts.OperatorsStr}),
		Run: func(cmd *cobra.Command, args []string) {
			OperatorsCmd(cmd, con, args)
		},
		GroupID: consts.GenericHelpGroup,
	}
	flags.Bind("operators", false, operatorsCmd, func(f *pflag.FlagSet) {
		f.IntP("timeout", "t", flags.DefaultTimeout, "grpc timeout in seconds")
	})

	return []*cobra.Command{operatorsCmd}
}
