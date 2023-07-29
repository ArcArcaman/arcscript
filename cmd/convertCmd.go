package cmd

import (
	"github.com/ArcArcaman/arcscript/cmd/converters"
	"github.com/spf13/cobra"
)

var (
	ConvertCmd = &cobra.Command{
		Use:   "convert",
		Short: "Converter scripts",
	}
)

func init() {
	ConvertCmd.AddCommand(converters.Base64ConvertCmd)
}
