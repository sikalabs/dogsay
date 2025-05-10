package cmd

import (
	"github.com/sikalabs/dogsay/cmd/root"
	_ "github.com/sikalabs/dogsay/cmd/version"
	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(root.Cmd.Execute())
}
