package root

import (
	"strings"

	"github.com/sikalabs/dogsay/pkg/dogsay"
	"github.com/sikalabs/dogsay/version"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "dogsay <text>",
	Short: "dogsay, " + version.Version,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dogsay.PrintDogSay(strings.Join(args, " "))
	},
}
