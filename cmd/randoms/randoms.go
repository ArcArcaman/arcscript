package randoms

import (
	"github.com/ArcArcaman/arcscript/cmd/randoms/eddsa"
	"github.com/ArcArcaman/arcscript/cmd/randoms/strings"
	"github.com/spf13/cobra"
)

var (
	RandomStringCmd = &cobra.Command{
		Use:   "string <length>",
		Short: "Generate Random Strings",
		Args:  cobra.ExactArgs(1),
		Run:   strings.Generate,
	}

	RandomEdDSACmd = &cobra.Command{
		Use:   "eddsa",
		Short: "Generate Random EdDSA Key Pairs",
		Args:  cobra.NoArgs,
		Run:   eddsa.Generate,
	}
)

func init() {
	RandomStringCmd.PersistentFlags().StringP("charset", "c", "none", "Character set to use, possible values: [none, urlsafe, alpha, numeric, symbol, alphanumeric, symnumeric, alphasymbolic]")
	RandomEdDSACmd.PersistentFlags().StringP("out", "o", "./eddsa_key", "Output Private Key File location. Public Key File will be named with a \".pub\" as postfix")
}
