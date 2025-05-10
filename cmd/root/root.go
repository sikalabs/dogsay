package root

import (
	"github.com/sikalabs/dogsay/version"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "dogsay",
	Short: "dogsay, " + version.Version,
}
