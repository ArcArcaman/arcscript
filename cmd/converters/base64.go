package converters

import (
	"encoding/base64"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	Base64ConvertCmd = &cobra.Command{
		Use:   "base64 <value>",
		Short: "Base64 Encode / Decode",
		Args:  cobra.ExactArgs(1),
		Run:   convert,
	}
)

func init() {
	Base64ConvertCmd.PersistentFlags().BoolP("decode", "d", false, "Boolean")
}

func convert(cmd *cobra.Command, args []string) {
	decodeMode, err := cmd.Flags().GetBool("decode")
	if err != nil {
		logrus.Panicf("Error when retrieving decode flag, %s", err.Error())
		return
	}

	value := args[0]

	if decodeMode {
		txt, err := base64.StdEncoding.DecodeString(value)
		if err != nil {
			logrus.Fatalf("Cannot Decode String %s, %s", value, err.Error())
			return
		}
		fmt.Printf("(Decode) Result: %s\n", txt)
		fmt.Printf("(Decode) Result[bytes]: %v\n", txt)
		fmt.Printf("(Decode) Result[length]: %d\n", len(txt))
		return
	}

	bytVal := []byte(value)
	txt := base64.StdEncoding.EncodeToString(bytVal)
	fmt.Printf("(Encode) Result: %s\n", txt)
	return
}
