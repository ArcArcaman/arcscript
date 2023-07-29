package cmd

import "github.com/spf13/cobra"

var (
	rootCmd = &cobra.Command{
		Use:   "arcscript",
		Short: "Utility Scripts by ArcArcaman",
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(ConvertCmd, RandomCmd)
}
