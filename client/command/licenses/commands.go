package licenses

import (
	"github.com/spf13/cobra"

	"github.com/Oni-kuki/sliver/client/command/help"
	"github.com/Oni-kuki/sliver/client/console"
	consts "github.com/Oni-kuki/sliver/client/constants"
	"github.com/Oni-kuki/sliver/client/licenses"
)

// Commands returns the `licences` command.
func Commands(con *console.SliverClient) []*cobra.Command {
	licensesCmd := &cobra.Command{
		Use:   consts.LicensesStr,
		Short: "Open source licenses",
		Long:  help.GetHelpFor([]string{consts.LicensesStr}),
		Run: func(cmd *cobra.Command, args []string) {
			con.Println(licenses.All)
		},
		GroupID: consts.GenericHelpGroup,
	}

	return []*cobra.Command{licensesCmd}
}
