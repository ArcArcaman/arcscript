package cmd

import (
	"github.com/ArcArcaman/arcscript/cmd/randoms"
	"github.com/spf13/cobra"
)

var (
	RandomCmd = &cobra.Command{
		Use:   "random",
		Short: "Random generator scripts",
	}
)

func init() {
	RandomCmd.AddCommand(randoms.RandomStringCmd)
	RandomCmd.AddCommand(randoms.RandomEdDSACmd)
}
